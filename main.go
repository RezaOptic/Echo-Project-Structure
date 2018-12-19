package main

import (
	"github.com/Echo-Project-Structure/config"
	"github.com/Echo-Project-Structure/controllers"
	"github.com/Echo-Project-Structure/utils/auth"

	"github.com/labstack/echo"
)

func main() {
	config.Init()
	auth.Init(&config.ClientSecret)

	e := echo.New()
	config.RegisterMiddlewares(e)
	controllers.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
