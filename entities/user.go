package entities

type User struct {
	Id       int
	Email    string
	Password string
	Token    string
	Name     string
}

type UsecaseInterface interface {
	Register(data *User) (*User, error)
}

type RepositoryInterface interface {
	Register(data *User) error
}
