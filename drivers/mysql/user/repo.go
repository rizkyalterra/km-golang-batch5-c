package user

import (
	"clean_architecture/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Register(data *entities.User) (error) {
	userDb := FromEntities(data)
	err := ur.db.Create(&userDb).Error
	if err != nil {
		return err
	}
	data.Id = int(userDb.ID)
	return nil
}