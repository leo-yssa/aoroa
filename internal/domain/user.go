package domain

type User struct {
	ID   uint
	Name string
}

func NewUser(id uint, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}