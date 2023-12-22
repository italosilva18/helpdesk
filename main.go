package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.Start(":8080")
	//e.Logger.Fatal(e.Start(":1323"))
}
