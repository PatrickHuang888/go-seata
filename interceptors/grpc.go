package interceptors

import (
	"context"
	"github.com/PatrickHuang888/go-seata/api"
	"github.com/PatrickHuang888/go-seata/logging"
	"github.com/PatrickHuang888/go-seata/protocol/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func TCCParticipantInterceptor(res api.TCCResource) grpc.UnaryServerInterceptor {

	if res.BranchType()!=
	rm.RegisterResource(res.BranchType(), res.Id())

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("no metadata")
		}

		vs, found := md[api.SeataXidKey]
		if !found {
			vs, found = md[strings.ToLower(api.SeataXidKey)]
		}

		if !found {
			logging.Info("no xid, not in a tx")
			return handler(ctx, req)
		}

		xid := vs[0]

		_, err := rm.RegisterBranch(pb.BranchTypeProto_TCC, xid, res.Id())
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func TCCLauncher(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

}
