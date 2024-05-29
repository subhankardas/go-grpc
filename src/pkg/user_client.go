package pkg

import (
	"github.com/subhankardas/go-grpc/proto"
	"google.golang.org/grpc"
)

// NewUserClient creates a new user client using gRPC client connection.
// The returned object will be used by third party packages to make gRPC 
// requests to our user service.
func NewUserClient(conn *grpc.ClientConn) proto.UserServiceClient {
	return proto.NewUserServiceClient(conn)
}
