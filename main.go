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
	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost", "http://localhost:3000", "https://localhost", "https://localhost:3000"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
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
