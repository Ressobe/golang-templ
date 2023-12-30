package main

import (
	"github.com/Ressobe/golang-templ/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/assets", "assets")

	userHandler := handler.UserHandler{}
	app.GET("/", userHandler.HandleUserShow)

	app.Start(":3000")
}
