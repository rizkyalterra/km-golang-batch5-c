package main

import (
	"clean_architecture/configs"
	"clean_architecture/controllers/users"
	"clean_architecture/drivers/mysql"
	userRepo "clean_architecture/drivers/mysql/user"
	"clean_architecture/routes"
	"clean_architecture/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.InitDB(configs.InitConfigDB())
	e := echo.New()

	
	userRepoInterface := userRepo.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepoInterface)
	userController := users.NewUserController(userUsecase)

	var listController = routes.ListController {
		Usercontroller: userController,
	}

	listController.InitRoute(e)
	e.Start(":8081")
}