package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mimaken3/ShareIT-api/application/server"
	"google.golang.org/appengine"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	// セッションを設定
	// store := session.NewCookieStore([]byte("secret-key"))

	// セッション保持時間(10分)
	// store.MaxAge(60 * 10)

	server.InitRouting(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		e.Logger.Printf("Defaulting to port %s", port)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	appengine.Main()
}
