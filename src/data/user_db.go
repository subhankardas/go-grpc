package data

import (
	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/app"
	"github.com/subhankardas/go-grpc/src/models"
)

// Simulate the database layer.
type userDatabase struct {
	users map[int32]*proto.User
}

// Mock database for user specific operations.
type DB interface {
	GetUser(int32) (*proto.User, error)          // Gets user by ID.
	GetUsers([]int32) []*proto.User              // Gets users by IDs.
	SearchUsers(models.UserSearch) []*proto.User // Gets users by search criteria.
}

// Create a new user database instance.
func NewUserDatabase() DB {
	return &userDatabase{
		users: db,
	}
}

// Mock implementation for getting user by ID.
func (db *userDatabase) GetUser(id int32) (*proto.User, error) {
	user, found := db.users[id]
	if !found {
		return nil, app.ErrUserNotFound
	}
	return user, nil
}

// Mock implementation for getting users by IDs.
func (db *userDatabase) GetUsers(ids []int32) []*proto.User {
	var users []*proto.User
	for _, id := range ids {
		if user, found := db.users[id]; found {
			users = append(users, user)
		}
	}
	return users
}

// Mock implementation for searching users by search criteria.
// Search criteria is used to filter users by city, phone number, and marital status.
// Returns all users if no search criteria is provided.
func (db *userDatabase) SearchUsers(criteria models.UserSearch) []*proto.User {
	var users []*proto.User
	for _, user := range db.users {
		match := true
		if criteria.City != "" && user.City != criteria.City {
			match = false
		}
		if criteria.Phone != 0 && user.Phone != criteria.Phone {
			match = false
		}
		if user.Married != criteria.Married {
			match = false
		}
		if match {
			users = append(users, user)
		}
	}
	return users
}
