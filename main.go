package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mimaken3/ShareIT-api/application/server"
	"google.golang.org/appengine"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	server.InitRouting(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Printf("Defaulting to port %s", port)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	appengine.Main()
}
