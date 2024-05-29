package app

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* Common errors for gRPC request responses. */

var (
	ErrUserNotFound  = status.Errorf(codes.NotFound, "user not found")                // User not found error.
	ErrInvalidUserID = status.Errorf(codes.InvalidArgument, "invalid user ID")        // Invalid user ID error.
	ErrEmptyUserIDs  = status.Errorf(codes.InvalidArgument, "empty user IDs")         // Empty user IDs error.
	ErrInvalidSearch = status.Errorf(codes.InvalidArgument, "invalid search request") // Invalid search error.
)
