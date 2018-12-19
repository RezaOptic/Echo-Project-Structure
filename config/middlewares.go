package config

import (
	"net/http"

	"github.com/Echo-Project-Structure/utils"
	"github.com/Echo-Project-Structure/utils/auth"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//RegisterMiddlewares Registers Middlewares
func RegisterMiddlewares(e *echo.Echo) {

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = customHTTPErrorHandler

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var jsonBody interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		jsonBody = he.Message
	}
	if be, ok := err.(utils.HTTPError); ok {
		c.JSON(be.StatusCode, be)
		return
	}
	if be, ok := err.(auth.HTTPError); ok {
		c.JSON(be.StatusCode, be)
		return
	}
	if err := c.JSON(code, jsonBody); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
