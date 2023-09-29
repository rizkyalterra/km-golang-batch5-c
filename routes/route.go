package routes

import (
	usercontroller "clean_architecture/controllers/users"

	"github.com/labstack/echo/v4"
)

type ListController struct {
	Usercontroller *usercontroller.UserController
}

func (c *ListController) InitRoute(e *echo.Echo) {
	e.POST("/register", c.Usercontroller.Register)
	
	// eAuth := e.Group("")
	// e.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
}