package main

import (
	"alex/taxi-server/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	app.Start(e)
}

