package main

import (
	"fmt"
	"net/http"
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
	if appengine.IsAppEngine() {
		// GAE
		realRootURL := os.Getenv("REAL_ROOT_URL")
		myPublicIP := os.Getenv("MY_PUBLIC_IP")
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{realRootURL, myPublicIP, "http://shareit.fun", "https://shareit.fun"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	} else {
		// Local
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8088"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	}

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
