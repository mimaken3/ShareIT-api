package main

import (
	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT-api/application/server"
)

func main() {
	e := echo.New()

	server.InitRouting(e)

	e.Logger.Fatal(e.Start(":1454"))
}
