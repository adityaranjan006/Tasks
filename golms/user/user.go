package user

type User struct {
	Name   string
	Number string
}

func NewUser(name, number string) *User {
	return &User{
		Name:   name,
		Number: number,
	}
}
