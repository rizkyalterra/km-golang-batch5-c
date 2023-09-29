package users

import (
	"clean_architecture/controllers"
	"clean_architecture/controllers/users/request"
	"clean_architecture/controllers/users/response"
	"clean_architecture/entities"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase entities.UsecaseInterface
}

func NewUserController(userUsecase entities.UsecaseInterface) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (uc *UserController) Register(c echo.Context) error {
	var register request.Register
	c.Bind(&register)
	data, err := uc.userUsecase.Register(register.ToEntities())
	if err != nil {
		return controllers.NewErrorResponse(c, err)
	}
	return controllers.NewSuccessResponse(c, response.FromEntities(data))
}