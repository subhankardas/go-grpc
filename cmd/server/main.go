package main

import (
	"log"
	"net"

	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/config"
	"github.com/subhankardas/go-grpc/src/data"
	"github.com/subhankardas/go-grpc/src/services"
	"google.golang.org/grpc"
)

type dependencies struct {
	cfg     *config.Config
	db      data.DB
	userSvc proto.UserServiceServer
}

func main() {
	// Create dependencies
	deps := createDependencies()

	// Start listening for gRPC connections
	listener, err := net.Listen("tcp", deps.cfg.Port)
	if err != nil {
		log.Fatal("Error while listening: ", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register our services
	proto.RegisterUserServiceServer(grpcServer, deps.userSvc)

	// Start server for gRPC requests
	log.Println("Starting gRPC server on", deps.cfg.Port)
	if grpcServer.Serve(listener); err != nil {
		log.Fatal("Error while serving: ", err)
	}
}

// Create dependencies required by our server.
func createDependencies() *dependencies {
	db := data.NewUserDatabase()
	return &dependencies{
		cfg:     config.NewConfig(),
		db:      db,
		userSvc: services.NewUserServer(db),
	}
}
