package main

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	"github.com/mimaken3/ShareIT-api/application/server"
	"google.golang.org/appengine"
)

func main() {
	e := echo.New()

	// CORS
	e.Use(middleware.CORS())

	// 認証チェック
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))))

	server.InitRouting(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		e.Logger.Printf("Defaulting to port %s", port)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
	appengine.Main()
}
