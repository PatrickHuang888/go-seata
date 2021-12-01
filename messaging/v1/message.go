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

	MsgTypeRequestSync       = MessageType(0)
	MsgTypeResponse          = MessageType(1)
	MsgTypeRequestOneway     = MessageType(2)
	MsgTypeHeartbeatRequest  = MessageType(3)
	MsgTypeHeartbeatResponse = MessageType(4)
)

var (
	requestId = uint32(0)

	MagicCodeBytes = []byte{0xda, 0xda}
	Version        = byte(1)
	SeataVersion   = "1.4.2"
)

type MessageType byte

func (mt MessageType) String() string {
	switch mt {
	case MsgTypeRequestSync:
		return "MsgTypeRequestSync"
	case MsgTypeRequestOneway:
		return "MsgTypeRequestOneway"
	case MsgTypeResponse:
		return "MsgTypeResponse"
	case MsgTypeHeartbeatRequest:
		return "MsgTypeHeartbeatRequest"
	case MsgTypeHeartbeatResponse:
		return "MsgTypeHeartbeatResponse"
	}
	return ""
}

type Message struct {
	Id  uint32
	Tp  MessageType
	CMP byte // compressor
	SER byte // serilizer
	Ver byte
	Msg proto.Message
}

func (m Message) String() string {
	return fmt.Sprintf("Message id %d, tp %s, msg %s", m.Id, m.Tp.String(),  m.Msg)
}

func newPbRmRegRequest(appId string, txGroup string, resourceIds string) *pb.RegisterRMRequestProto {
	return &pb.RegisterRMRequestProto{ResourceIds: resourceIds, AbstractIdentifyRequest: &pb.AbstractIdentifyRequestProto{
		AbstractMessage: &pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_RM},
		Version:         SeataVersion, ApplicationId: appId, TransactionServiceGroup: txGroup}}
}

func NewRmRegRequest(appId string, txGroup string, resourceIds string) *Message {
	msg := &Message{Id: nextRequestId()}
	msg.Tp = MsgTypeRequestOneway
	msg.Ver = Version
	msg.Msg = newPbRmRegRequest(appId, txGroup, resourceIds)
	return msg
}

func newPbTmRegRequest(appId string, txGroup string) *pb.RegisterTMRequestProto {
	return &pb.RegisterTMRequestProto{AbstractIdentifyRequest: &pb.AbstractIdentifyRequestProto{
		AbstractMessage: &pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_CLT},
		Version:         SeataVersion, ApplicationId: appId, TransactionServiceGroup: txGroup}}
}

func NewTmRegRequest(appId string, txGroup string) Message {
	msg := Message{Id: nextRequestId()}
	msg.Tp = MsgTypeRequestSync
	msg.SER = SerializerProtoBuf
	msg.Ver = Version
	msg.Msg = newPbTmRegRequest(appId, txGroup)
	return msg
}

func NewTmRegResponse(id uint32) *Message {
	msg := &Message{}
	msg.Id = id
	msg.Tp = MsgTypeResponse
	msg.SER = SerializerProtoBuf
	msg.Ver = Version
	msg.Msg = &pb.RegisterTMResponseProto{AbstractIdentifyResponse: &pb.AbstractIdentifyResponseProto{
		AbstractResultMessage: &pb.AbstractResultMessageProto{AbstractMessage: &pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_CLT_RESULT}}}}
	return msg
}

func EncodeMessage(msg *Message) ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.Write(MagicCodeBytes)
	buffer.WriteByte(Version)

	// full length 4 bytes
	buffer.Write(make([]byte, 4))
	// head length 2 bytes
	buffer.Write(make([]byte, 2))

	// message type
	buffer.WriteByte(byte(msg.Tp))

	// codec SHOULD be protobuf
	buffer.WriteByte(0x2)

	// compressor none
	// todo:
	buffer.WriteByte(msg.CMP)

	// request id 4 bytes
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(msg.Id))
	buffer.Write(buf)

	// optional headmap
	// todo:
	var headMapLength uint16

	var typeName string
	switch msg.Msg.(type) {
	case *pb.RegisterTMRequestProto:
		typeName = TypeNameTmRegisterRequest
	case *pb.RegisterTMResponseProto:
		typeName = TypeNameTmRegisterResponse
	case *pb.RegisterRMRequestProto:
		typeName = TypeNameRmRegisterRequest
	case *pb.RegisterRMResponseProto:
		typeName= TypeNameRmRegisterResponse
	default:
		return nil, errors.New("message type unknown")
	}

	typeNameLength := uint32(len(typeName))
	binary.BigEndian.PutUint32(buf, typeNameLength)
	buffer.Write(buf)
	buffer.Write([]byte(typeName))

	body, err := proto.Marshal(msg.Msg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buffer.Write(body)

	bs := buffer.Bytes()

	headLength := StartLength + headMapLength
	binary.BigEndian.PutUint16(bs[7:9], headLength)

	fullLength := uint32(headLength) + 4 + typeNameLength + uint32(len(body))
	binary.BigEndian.PutUint32(bs[3:7], fullLength)

	return bs, nil
}

func EncodeProtoMessage(msgType MessageType, msg proto.Message) (uint32, []byte, error) {
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

func DecodePbMessage(buffer *bytes.Buffer) (msg proto.Message, err error) {
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

func ReadMessage(conn net.Conn) (msg *Message, err error) {
	msg = &Message{}

	buf := make([]byte, StartLength)
	if _, err = io.ReadFull(conn, buf); err != nil {
		err = errors.Wrap(err, "read message start err")
		return
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
	msg.Tp = MessageType(buf[9])

	msg.SER = buf[10]
	if msg.SER != SerializerProtoBuf {
		err = errors.New("serializer only support protobuf")
		return
	}

	// skip compressor
	// todo:

	// request id
	msg.Id = binary.BigEndian.Uint32(buf[12:16])

	length := binary.BigEndian.Uint32(buf[3:7])

	buf = make([]byte, length-StartLength)
	if _, err = io.ReadFull(conn, buf); err != nil {
		err = errors.WithStack(err)
		return
	}

	buffer := bytes.NewBuffer(buf)
	msg.Msg, err = DecodePbMessage(buffer)
	return
}
