package data

import "github.com/subhankardas/go-grpc/proto"

/* MOCK USER DATABASE */
var db = map[int32]*proto.User{
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
	4: {
		Id:      4,
		Fname:   "Pooja",
		City:    "Mumbai",
		Phone:   1234567890,
		Height:  5,
		Married: false,
	},
	5: {
		Id:      5,
		Fname:   "Deepak",
		City:    "Agra",
		Phone:   1234567890,
		Height:  4.6,
		Married: true,
	},
}
