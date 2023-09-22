package controllers

import (
	"beritaalta/configs"
	"beritaalta/models/base"
	"beritaalta/models/news"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteNewsController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse {
			Status: false,
			Message: "Id news is not number",
			Data: nil,
		})
	}
	result := configs.DB.Delete(&news.News{}, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusNoContent, base.BaseResponse {
		Status: true,
		Message: "Success delete",
		Data: nil,
	})
}

func AddNewsController(c echo.Context) error {
	var newsRequestAdd news.NewsRequestAdd
	c.Bind(&newsRequestAdd)

	//validasi
	if newsRequestAdd.Title == "" {
		return c.JSON(http.StatusBadRequest, base.BaseResponse {
			Status: false,
			Message: "Title still empty",
			Data: nil,
		})
	}

	var newsDB news.News
	newsDB = newsDB.MapFromRequestAdd(newsRequestAdd)
	result := configs.DB.Create(&newsDB)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse {
			Status: false,
			Message: "Error when add news in database",
			Data: nil,
		})
	}

	var newsResponse news.NewsResponse
	newsResponse.MapFromDatabase(newsDB)

	return c.JSON(http.StatusOK, base.BaseResponse {
		Status: false,
		Message: "Success add news",
		Data: newsResponse,
	})
}

func GetNewsController(c echo.Context) error {
	var newsDB []news.News
	result := configs.DB.Find(&newsDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var newsResponse []news.NewsResponse
	newsResponse = news.MapFromDatabaseList(newsDB)

	return c.JSON(http.StatusOK, base.BaseResponse {
		Status: true,
		Message: "Success get News",
		Data: newsResponse,
	})
}