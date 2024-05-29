package mocks

import (
	"github.com/subhankardas/go-grpc/proto"
	"github.com/subhankardas/go-grpc/src/app"
	"github.com/subhankardas/go-grpc/src/data"
	"github.com/subhankardas/go-grpc/src/models"
)

/* THIS IS A PLACEHOLDER MOCK DATABASE, USE MOCKGEN FOR BETTER IMPLEMENTATION */
type mockUserDB struct {
	users map[int32]*proto.User
}

func NewMockUserDatabase() data.DB {
	return &mockUserDB{
		users: map[int32]*proto.User{
			1: {
				Id:      1,
				Fname:   "Subhankar",
				City:    "Mumbai",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
			2: {
				Id:      2,
				Fname:   "Kamal",
				City:    "Delhi",
				Phone:   1234567890,
				Height:  5.9,
				Married: true,
			},
			3: {
				Id:      3,
				Fname:   "Divya",
				City:    "Mumbai",
				Phone:   1234567890,
				Height:  4.3,
				Married: false,
			},
		},
	}
}

func (db *mockUserDB) GetUser(id int32) (*proto.User, error) {
	// Some mock error
	if id > 3 {
		return nil, app.ErrUserNotFound
	}
	// Some mock success
	return db.users[id], nil
}

func (db *mockUserDB) GetUsers(ids []int32) []*proto.User {
	// Some mock error
	if len(ids) == 0 {
		return []*proto.User{}
	}
	// Some mock success
	return []*proto.User{db.users[ids[0]], db.users[ids[1]]}
}

func (db *mockUserDB) SearchUsers(criteria models.UserSearch) []*proto.User {
	// Some mock error
	if criteria.City == "" && criteria.Phone == 0 && criteria.Married == false {
		return []*proto.User{}
	}
	// Some mock success
	return []*proto.User{db.users[1], db.users[2]}
}
