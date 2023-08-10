package model

type authDomain struct {
	id       string
	email    string
	password string
	name     string
}

func (ud *authDomain) GetId() string {
	return ud.id
}

func (ud *authDomain) SetId(id string) {
	ud.id = id
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
