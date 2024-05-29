package validators

import (
	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/app"
)

// Validate user ID request.
func ValidateUserIdRequest(req *proto.UserIdRequest) error {
	if req.GetId() <= 0 {
		return app.ErrInvalidUserID
	}
	return nil
}

// Validate user IDs request.
func ValidateUserIdsRequest(req *proto.UserIdsRequest) error {
	if len(req.GetIds()) == 0 {
		return app.ErrEmptyUserIDs
	}
	return nil
}

// Validate search request.
func ValidateSearchRequest(req *proto.SearchRequest) error {
	if req.GetCity() == "" && req.GetPhone() == 0 {
		return app.ErrInvalidSearch
	}
	return nil
}
