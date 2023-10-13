package controllers

import (
	"Prioritas2/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestGetBooks(t *testing.T){
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetBooksController(c)

	assert.NoError(t, err) 
	assert.Equal(t, http.StatusOK, rec.Code)
}

// create create book controller
func TestCreateBook(t *testing.T){
    config.InitDB()
    e := echo.New()
    req := httptest.NewRequest(http.MethodPost, "/books", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    err := CreateBookController(c)

    assert.NoError(t, err) 
    assert.Equal(t, http.StatusOK, rec.Code)
}