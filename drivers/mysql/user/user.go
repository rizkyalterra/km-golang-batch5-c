package user

import (
	"clean_architecture/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func FromEntities(data *entities.User) User {
	return User{
		Model:    gorm.Model{
			ID:        uint(data.Id),
		},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func (u User) ToEntities() *entities.User {
	return &entities.User{
		Id:       int(u.ID),
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	}
}
