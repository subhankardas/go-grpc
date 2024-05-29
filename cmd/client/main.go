package main

import (
	"context"
	"log"

	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/config"
	"github.com/subhankardas/go-grpc/src/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/* THIS IS A SAMPLE USAGE FOR THE GRPC USER CLIENT. */
func main() {
	// Create config
	cfg := config.NewConfig()

	// Create gRPC client connection
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(cfg.Port, opts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create user client using gRPC client connection
	client := pkg.NewUserClient(conn)

	// Make gRPC request and print response
	req := &proto.UserIdRequest{Id: 2}
	res, err := client.GetUser(context.Background(), req)
	log.Printf("Response: %v err: %v", res, err)
}
