package usecases

import (
	"clean_architecture/entities"
	"clean_architecture/utils/bcrypt"
	"clean_architecture/utils/errors"
	"clean_architecture/utils/jwt"
)

type UserUsecase struct {
	repository entities.RepositoryInterface
}

func NewUserUsecase(repoInterface entities.RepositoryInterface) *UserUsecase {
	return &UserUsecase{
		repository: repoInterface,
	}
}

func (s *UserUsecase) Register(data *entities.User) (*entities.User, error) {
	if data.Email == "" {
		return nil, errors.ErrUserEmailEmpty
	} else if data.Name == "" {
		return nil, errors.ErrUserNameEmpty
	} else if data.Password == "" {
		return nil, errors.ErrUserPasswordEmpty
	}

	hashPass, err := bcrypt.Hash(data.Password)
	if err != nil {
		return nil, errors.ErrBcryptPassword
	}

	// check didatabase ada nggk email ?

	data.Password = hashPass
	err = s.repository.Register(data)
	if err != nil {
		return nil, errors.ErrRegisterUserDatabase
	}

	data.Token = jwt.GenerateTokenJWT(data.Id, data.Name)
	
	return data, nil
}