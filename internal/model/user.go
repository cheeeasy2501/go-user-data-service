package model

type User struct {
	Id        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Active    bool   `json:"active"`
}
