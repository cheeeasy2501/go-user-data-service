package model

type User struct {
	Id        uint64
	Email     string
	Password  string
	FirstName string
	LastName  string
	Active    bool
}
