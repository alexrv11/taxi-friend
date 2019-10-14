package main

import (
	"alex/taxi-server/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main(){
	e := echo.New()

	e.AutoTLSManager.Cache = autocert.DirCache("~/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	app.Start(e)
}

