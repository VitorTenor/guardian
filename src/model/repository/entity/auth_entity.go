package entity

type AuthEntity struct {
	ID       int
	Email    string
	Password []byte
	Name     string
}
