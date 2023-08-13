package model

type authDomain struct {
	ID       int
	email    string
	password string
	name     string
}

func (ud *authDomain) GetId() int {
	return ud.ID
}

func (ud *authDomain) SetId(id int) {
	ud.ID = id
}

func (ud *authDomain) GetEmail() string {
	return ud.email
}

func (ud *authDomain) GetPassword() string {
	return ud.password
}

func (ud *authDomain) GetName() string {
	return ud.name
}

func (ud *authDomain) SetName(name string) {
	ud.name = name
}
