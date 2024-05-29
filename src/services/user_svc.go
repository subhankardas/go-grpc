package services

import (
	"context"
	"log"

	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/data"
	"github.com/subhankardas/go-grpc/src/models"
	"github.com/subhankardas/go-grpc/src/validators"
)

// Implements gRPC server for user service.
type userServer struct {
	proto.UnimplementedUserServiceServer
	db data.DB // Database instance for user specific operations.
}

// Create a new user server instance.
func NewUserServer(db data.DB) proto.UserServiceServer {
	return &userServer{
		db: db,
	}
}

// GetUser implements gRPC call for getting user by ID.
func (s *userServer) GetUser(ctx context.Context, req *proto.UserIdRequest) (*proto.UserResponse, error) {
	// Validate user ID request
	if err := validators.ValidateUserIdRequest(req); err != nil {
		log.Printf("Error validating get user request. err: %v", err)
		return nil, err
	}

	id := req.GetId()

	log.Println("Getting user by ID:", id)
	user, err := s.db.GetUser(id)
	if err != nil {
		log.Printf("Error getting user by ID: %v err: %v", id, err)
		return nil, err
	}
	return &proto.UserResponse{User: user}, nil
}

// GetUsers implements gRPC call for getting users by IDs.
func (s *userServer) GetUsers(ctx context.Context, req *proto.UserIdsRequest) (*proto.UsersResponse, error) {
	// Validate user IDs request
	if err := validators.ValidateUserIdsRequest(req); err != nil {
		log.Printf("Error validating get users request. err: %v", err)
		return nil, err
	}

	ids := req.GetIds()

	log.Println("Getting users by IDs:", ids)
	users := s.db.GetUsers(ids)
	return &proto.UsersResponse{Users: users}, nil
}

// SearchUsers implements gRPC call for searching users by search criteria.
func (s *userServer) SearchUsers(ctx context.Context, req *proto.SearchRequest) (*proto.UsersResponse, error) {
	// Validate search request
	if err := validators.ValidateSearchRequest(req); err != nil {
		log.Printf("Error validating search request. err: %v", err)
		return nil, err
	}

	criteria := s.getCriteria(req)

	log.Println("Searching users by criteria:", criteria)
	users := s.db.SearchUsers(criteria)
	return &proto.UsersResponse{Users: users}, nil
}

// Get criteria from search request.
func (s *userServer) getCriteria(req *proto.SearchRequest) models.UserSearch {
	return models.UserSearch{
		City:    req.GetCity(),
		Phone:   req.GetPhone(),
		Married: req.GetMarried(),
	}
}
