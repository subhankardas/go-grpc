package models

// UserSearch is used to filter users by city, phone number, and marital status.
type UserSearch struct {
	City    string
	Phone   int64
	Married bool
}
