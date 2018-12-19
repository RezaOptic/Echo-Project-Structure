package controllers

import (
	"github.com/Echo-Project-Structure/logic"
	"github.com/Echo-Project-Structure/models"
	"github.com/labstack/echo"
	"net/http"
)

func Ping(c echo.Context) error {
	b := new(models.Ping)
	if err := c.Bind(b); err != nil {
		return err
	}
	_, err := logic.Ping(*b)
	if err != nil{
		return err
	}
	return c.NoContent(http.StatusCreated)
}