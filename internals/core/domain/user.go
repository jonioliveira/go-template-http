package domain

type User struct {
	id       int
	email    string
	password string
}

func NewUser(id int, email string, Password string) *User {
	return &User{
		id:       id,
		email:    email,
		password: Password,
	}
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetId() int {
	return u.id
}
