package v1

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"sync/atomic"
)

/*
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
 */

const (
	SerializerProtoBuf          = 2
	TmRegisterRequestClassName  = "io.seata.protocol.protobuf.RegisterTMRequestProto"
	TmRegisterResponseClassName = "io.seata.protocol.protobuf.RegisterTMResponseProto"

	StartLength = 16

	MSGTYPE_RESQUEST_SYNC = RequestType(0)
)

var (
	requestId = uint32(0)

	MagicCodeBytes = []byte{0xda, 0xda}
	Version        = byte(1)
)

type RequestType byte

func NewTmRegRequest(appId string, txGroup string) *pb.RegisterTMRequestProto {
	return &pb.RegisterTMRequestProto{AbstractIdentifyRequest: &pb.AbstractIdentifyRequestProto{
		AbstractMessage: &pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_CLT},
		Version:         "1.4.2", ApplicationId: appId, TransactionServiceGroup: txGroup}}
}

func EncodeMessage(reqType RequestType, messageType pb.MessageTypeProto, msg proto.Message) (uint32, []byte, error) {
	buffer := new(bytes.Buffer)

	buffer.Write(MagicCodeBytes)
	buffer.WriteByte(Version)

	// full length 4 bytes
	buffer.Write(make([]byte, 4))
	// head length 2 bytes
	buffer.Write(make([]byte, 2))

	// message type
	buffer.WriteByte(byte(reqType))

	// codec protobuf
	buffer.WriteByte(0x2)

	// compressor none
	// todo:
	buffer.WriteByte(0)

	// request id 4 bytes
	buf := make([]byte, 4)
	reqId:= nextRequestId()
	binary.BigEndian.PutUint32(buf, reqId)
	buffer.Write(buf)

	// optional headmap
	// todo:
	var headMapLength uint16

	var className string
	switch messageType {
	case pb.MessageTypeProto_TYPE_REG_CLT:
		// body
		className = TmRegisterRequestClassName
	default:
		return 0, nil, errors.New("message type unknown")
	}
	classNameLength := uint32(len(className))
	binary.BigEndian.PutUint32(buf, classNameLength)
	buffer.Write(buf)
	buffer.Write([]byte(className))

	body, err := proto.Marshal(msg)
	if err != nil {
		return 0, nil, errors.WithStack(err)
	}
	buffer.Write(body)

	bs := buffer.Bytes()

	headLength := StartLength + headMapLength
	binary.BigEndian.PutUint16(bs[7:9], headLength)

	fullLength := uint32(headLength) + 4 + classNameLength + uint32(len(body))
	binary.BigEndian.PutUint32(bs[3:7], fullLength)

	return reqId, bs, nil
}

func Decode(buffer *bytes.Buffer) (msg proto.Message, err error) {
	buf := make([]byte, 4)

	// className
	buffer.Read(buf[:4])
	classNameLength := binary.BigEndian.Uint32(buf)
	className := make([]byte, classNameLength)
	buffer.Read(className)

	switch string(className) {
	case TmRegisterResponseClassName:
		msg = &pb.RegisterTMResponseProto{}
	case TmRegisterRequestClassName:
		msg = &pb.RegisterTMRequestProto{}

	default:
		err = errors.New("response class name unknown")
		return
	}

	if err = proto.Unmarshal(buffer.Bytes(), msg); err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func nextRequestId() uint32 {
	atomic.AddUint32(&requestId, 1)
	return requestId
}

func ReadMessage(conn net.Conn) (reqId uint32, msg proto.Message, err error) {

	buf := make([]byte, StartLength)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return 0, nil, errors.Wrap(err, "read message start err")
	}

	fmt.Print("after read start\n")

	// magic code
	if buf[0] != MagicCodeBytes[0] {
		err = errors.New("magic code error")
		return
	}
	if buf[1] != MagicCodeBytes[1] {
		err = errors.New("magic code error")
		return
	}

	// skip head now
	//buffer.Read(buf[:2])
	//headLength = binary.BigEndian.Uint16(buf[:2])
	// todo:

	// message type
	//buffer.ReadByte()

	serializer := buf[10]
	if serializer != SerializerProtoBuf {
		err = errors.New("serializer only support protobuf")
		return
	}

	// skip compressor
	// todo:

	// request id
	reqId = binary.BigEndian.Uint32(buf[12:16])

	length := binary.BigEndian.Uint32(buf[3:7])

	buf = make([]byte, length-StartLength)
	if _, err = io.ReadFull(conn, buf); err != nil {
		err = errors.WithStack(err)
		return
	}

	buffer := bytes.NewBuffer(buf)
	msg, err = Decode(buffer)
	return
}
