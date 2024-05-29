package services

import (
	"context"
	"log"
	"net"
	"reflect"
	"testing"

	"github.com/subhankardas/go-grpc/mocks"
	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/test/bufconn"
)

const bufferSize = 1024 * 1024

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(bufferSize)
	grpcServer := grpc.NewServer()

	// Inject mock database
	db := mocks.NewMockUserDatabase()

	// Register our services
	proto.RegisterUserServiceServer(grpcServer, NewUserServer(db))

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Error while listening:", err)
		}
	}()
}

func TestGetUser(t *testing.T) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	resolver.SetDefaultScheme("passthrough")
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), opts)
	if err != nil {
		log.Printf("error connecting to server: %v", err)
		panic(err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	tests := []struct {
		name string
		req  *proto.UserIdRequest
		res  *proto.UserResponse
		err  error
	}{
		{
			name: "valid user id",
			req:  &proto.UserIdRequest{Id: 1},
			res: &proto.UserResponse{User: &proto.User{
				Id: 1, Fname: "Subhankar", City: "Mumbai", Phone: 1234567890, Height: 5.8, Married: true,
			}},
			err: nil,
		},
		{
			name: "invalid user id",
			req:  &proto.UserIdRequest{Id: 4},
			res:  nil,
			err:  app.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := client.GetUser(context.Background(), tt.req)

			if err != nil && tt.err == nil {
				t.Errorf("error getting user: %v", err)
			}
			if err == nil && tt.err != nil {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if err == nil && !reflect.DeepEqual(res.User, tt.res.User) {
				t.Errorf("expected response: %v, got: %v", tt.res.User, res.User)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	resolver.SetDefaultScheme("passthrough")
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), opts)
	if err != nil {
		log.Printf("error connecting to server: %v", err)
		panic(err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	tests := []struct {
		name string
		req  *proto.UserIdsRequest
		res  *proto.UsersResponse
		err  error
	}{
		{
			name: "valid user ids",
			req:  &proto.UserIdsRequest{Ids: []int32{1, 2, 3}},
			res: &proto.UsersResponse{Users: []*proto.User{
				{Id: 1, Fname: "Subhankar", City: "Mumbai", Phone: 1234567890, Height: 5.8, Married: true},
				{Id: 2, Fname: "Kamal", City: "Delhi", Phone: 1234567890, Height: 5.9, Married: true},
			}},
			err: nil,
		},
		{
			name: "empty user ids",
			req:  &proto.UserIdsRequest{Ids: []int32{}},
			res:  nil,
			err:  app.ErrEmptyUserIDs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := client.GetUsers(context.Background(), tt.req)
			if err != nil && tt.err == nil {
				t.Errorf("error getting users: %v", err)
			}
			if err == nil && tt.err != nil {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if err == nil && !reflect.DeepEqual(res.Users, tt.res.Users) {
				t.Errorf("expected response: %v, got: %v", tt.res.Users, res.Users)
			}
		})
	}
}
