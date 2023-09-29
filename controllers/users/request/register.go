package request

import "clean_architecture/entities"

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d Register) ToEntities() *entities.User {
	return &entities.User{
		Email:    d.Email,
		Password: d.Password,
		Name:     d.Name,
	}
}