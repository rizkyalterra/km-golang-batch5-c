package main

import (
	"beritaalta/configs"
	"beritaalta/controllers"

	"github.com/labstack/echo/v4"
)

func main(){
	configs.LoadEnv()
	configs.InitDatabase()
	e := echo.New()
	e.GET("/news", controllers.GetNewsController)
	e.Start(":8000")
}