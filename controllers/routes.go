package controllers

import (
	"github.com/labstack/echo"
)

//RegisterRoutes register routes
func RegisterRoutes(e *echo.Echo) {
	e.POST("/ping", Ping)

}
