package messaging

import "github.com/PatrickHuang888/go-seata/protocol/pb"

func NewTmRegRequest(appId string, txGroup string) *pb.RegisterTMRequestProto{
	return &pb.RegisterTMRequestProto{AbstractIdentifyRequest:&pb.AbstractIdentifyRequestProto{
		AbstractMessage:&pb.AbstractMessageProto{MessageType: pb.MessageTypeProto_TYPE_REG_CLT},
		Version: "1.4.2", ApplicationId: appId, TransactionServiceGroup: txGroup}}
}
