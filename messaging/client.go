package messaging

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"sync/atomic"

	"github.com/pkg/errors"

	"github.com/PatrickHuang888/go-seata/protocol/pb"
)

/**
 * <pre>
 * 0     1     2     3     4     5     6     7     8     9    10     11    12    13    14    15    16
 * +-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+
 * |   magic   |Proto|     Full length       |    Head   | Msg |Seria|Compr|     RequestId         |
 * |   code    |colVer|    (head+body)      |   Length  |Type |lizer|ess  |                       |
 * +-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
 * |                                                                                               |
 * |                                   Head Map [Optional]                                         |
 * +-----------+-----------+-----------+-----------+-----------+-----------+-----------+-----------+
 * |                                                                                               |
 * |                                         body                                                  |
 * |                                                                                               |
 * |                                        ... ...                                                |
 * +-----------------------------------------------------------------------------------------------+
 * </pre>
 * <p>
 * <li>Full Length: include all data </li>
 * <li>Head Length: include head data from magic code to head map. </li>
 * <li>Body Length: Full Length - Head Length</li>
 * </p>
 * https://github.com/seata/seata/issues/893
 **/

const (
	MSGTYPE_RESQUEST_SYNC = byte(0)
	FixedHeadLength= 16
	StartLength= 7
	SerializerProtoBuf=2
)

var (
	MagicCodeBytes = []byte{0xda, 0xda}
	Version        = byte(1)

	requestId = uint32(0)
)

func nextRequestId() uint32 {
	atomic.AddUint32(&requestId, 1)
	return requestId
}

type Client struct {
	conn net.Conn
}

func NewClient(svrAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", svrAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{conn: conn}, nil
}

func (c *Client) SendTmRegMsg(msg *pb.RegisterTMRequestProto) (*pb.RegisterTMResponseProto, error) {

	buffer := new(bytes.Buffer)

	buffer.Write(MagicCodeBytes)
	buffer.WriteByte(Version)

	// full length 4 bytes
	buffer.Write(make([]byte, 4))
	// head length 2 bytes
	buffer.Write(make([]byte, 2))

	// message type
	buffer.WriteByte(MSGTYPE_RESQUEST_SYNC)
	// codec protobuf
	buffer.WriteByte(0x2)
	// compressor none
	buffer.WriteByte(0)
	// request id 4 bytes
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, nextRequestId())
	buffer.Write(buf)

	// optional headmap
	var headMapLength uint16

	// body
	className:= "io.seata.protocol.protobuf.RegisterTMRequestProto"
	classNameLength:= uint32(len(className))
	binary.BigEndian.PutUint32(buf, classNameLength)
	buffer.Write(buf)
	buffer.Write([]byte(className))

	body, err := proto.Marshal(msg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buffer.Write(body)

	bs := buffer.Bytes()

	headLength := 16 + headMapLength
	binary.BigEndian.PutUint16(bs[7:9], headLength)

	var fullLength uint32
	fullLength = uint32(headLength) + 4+ classNameLength+ uint32(len(body))
	binary.BigEndian.PutUint32(bs[3:7], fullLength)

	n := 0
	for n < len(bs) {
		nt, err := c.conn.Write(bs)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		n += nt
	}
	fmt.Printf("rm registry msg written\n")

	// =======read response ==========
	buf = make([]byte, 7)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return nil, errors.Wrap(err, "read message start err")
	}

	// magic code
	b, err := buffer.ReadByte()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if b != MagicCodeBytes[0] {
		return nil, errors.New("magic code error")
	}
	b, err = buffer.ReadByte()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if b != MagicCodeBytes[1] {
		return nil, errors.New("magic code error")
	}
	fullLength = binary.BigEndian.Uint32(buf[3:])

	buf = make([]byte, fullLength-StartLength)
	buffer= bytes.NewBuffer(buf)
	buffer.Reset()
	if _, err :=io.CopyN(buffer, c.conn, int64(len(buf))); err != nil {
		return nil, errors.WithStack(err)
	}
	buffer.Read(buf[:2])
	headLength = binary.BigEndian.Uint16(buf[:2])
	// skip headmap now

	// message type
	buffer.ReadByte()

	serializer, _ := buffer.ReadByte()
	if serializer!=SerializerProtoBuf {
		return nil, errors.New("serializer only support protobuf")
	}

	// compressor
	buffer.ReadByte()

	// request id
	buffer.Read(buf[:4])
	requestId= binary.BigEndian.Uint32(buf[:4])

	// className
	buffer.Read(buf[:4])
	classNameLength= binary.BigEndian.Uint32(buf)
	buffer.Read(make([]byte, classNameLength))

	rsp := &pb.RegisterTMResponseProto{}
	if err = proto.Unmarshal(buffer.Bytes(), rsp); err != nil {
		return nil, errors.WithStack(err)
	}
	return rsp, nil
}

