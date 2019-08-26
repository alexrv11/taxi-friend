package main

import (
	"alex/taxi-server/app"
	"github.com/labstack/echo"
)

func main(){
	e := echo.New()

	app.Start(e)
}

