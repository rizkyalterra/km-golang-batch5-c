package controllers

import (
	"beritaalta/configs"
	"beritaalta/models/news"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNewsController(c echo.Context) error {
	var news []news.News
	result := configs.DB.Find(&news)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, news)
}