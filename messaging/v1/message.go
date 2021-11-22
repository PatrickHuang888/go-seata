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
	SerializerProtoBuf         = 2
	TypeNameTmRegisterRequest  = "io.seata.protocol.protobuf.RegisterTMRequestProto"
	TypeNameTmRegisterResponse = "io.seata.protocol.protobuf.RegisterTMResponseProto"
	TypeNameRmRegisterRequest  = "io.seata.protocol.protobuf.RegisterRMRequestProto"
	TypeNameRmRegisterResponse = "io.seata.protocol.protobuf.RegisterRMResponseProto"

	StartLength = 16

	MSGTYPE_RESQUEST_SYNC     = MessageType(0)
	MSGTYPE_RESQUEST_ONEWAY   = MessageType(2)
	MSGTYPE_HEARTBEAT_REQUEST = MessageType(3)
)

var (
	requestId = uint32(0)

	MagicCodeBytes = []byte{0xda, 0xda}
	Version        = byte(1)
	SeataVersion   = "1.4.2"
)

type MessageType byte

func NewRmRegMessage(appId string, txGroup string, resourceIds string) *pb.RegisterRMRequestProto {
	return &pb.RegisterRMRequestProto{ResourceIds: resourceIds, AbstractIdentifyRequest: &pb.AbstractIdentifyRequestProto{
		AbstractMessage: &pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_RM},
		Version:         SeataVersion, ApplicationId: appId, TransactionServiceGroup: txGroup}}
}

func NewTmRegRequest(appId string, txGroup string) *pb.RegisterTMRequestProto {
	return &pb.RegisterTMRequestProto{AbstractIdentifyRequest: &pb.AbstractIdentifyRequestProto{
		AbstractMessage: &pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_CLT},
		Version:         SeataVersion, ApplicationId: appId, TransactionServiceGroup: txGroup}}
}

func EncodeMessage(msgType MessageType, msg proto.Message) (uint32, []byte, error) {
	buffer := new(bytes.Buffer)

	buffer.Write(MagicCodeBytes)
	buffer.WriteByte(Version)

	// full length 4 bytes
	buffer.Write(make([]byte, 4))
	// head length 2 bytes
	buffer.Write(make([]byte, 2))

	// message type
	buffer.WriteByte(byte(msgType))

	// codec protobuf
	buffer.WriteByte(0x2)

	// compressor none
	// todo:
	buffer.WriteByte(0)

	// request id 4 bytes
	buf := make([]byte, 4)
	reqId := nextRequestId()
	binary.BigEndian.PutUint32(buf, reqId)
	buffer.Write(buf)

	// optional headmap
	// todo:
	var headMapLength uint16

	var typeName string
	switch msg.(type) {
	case *pb.RegisterTMRequestProto:
		typeName = TypeNameTmRegisterRequest
	case *pb.RegisterRMRequestProto:
		typeName = TypeNameRmRegisterRequest
	default:
		return 0, nil, errors.New("message type unknown")
	}
	typeNameLength := uint32(len(typeName))
	binary.BigEndian.PutUint32(buf, typeNameLength)
	buffer.Write(buf)
	buffer.Write([]byte(typeName))

	body, err := proto.Marshal(msg)
	if err != nil {
		return 0, nil, errors.WithStack(err)
	}
	buffer.Write(body)

	bs := buffer.Bytes()

	headLength := StartLength + headMapLength
	binary.BigEndian.PutUint16(bs[7:9], headLength)

	fullLength := uint32(headLength) + 4 + typeNameLength + uint32(len(body))
	binary.BigEndian.PutUint32(bs[3:7], fullLength)

	return reqId, bs, nil
}

func Decode(buffer *bytes.Buffer) (msg proto.Message, err error) {
	buf := make([]byte, 4)

	buffer.Read(buf[:4])
	typeNameLength := binary.BigEndian.Uint32(buf)
	typeName := make([]byte, typeNameLength)
	buffer.Read(typeName)

	switch string(typeName) {
	case TypeNameTmRegisterResponse:
		msg = &pb.RegisterTMResponseProto{}
	case TypeNameTmRegisterRequest:
		msg = &pb.RegisterTMRequestProto{}
	case TypeNameRmRegisterRequest:
		msg = &pb.RegisterRMRequestProto{}
	case TypeNameRmRegisterResponse:
		msg = &pb.RegisterRMResponseProto{}
	default:
		err = errors.New("response type name unknown")
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

	// message type always response
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
