package main

import (
	"alex/taxi-server/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main(){
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("taxifriend.com")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	app.Start(e)
}

