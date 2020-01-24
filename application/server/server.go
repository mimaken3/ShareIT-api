package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT/application/server/handler"
)

func InitRouting(e *echo.Echo) {
	// 全ユーザを取得
	e.GET("/users", handler.FindAllUsers())

	// ユーザを取得
	e.GET("/user/:user_id", handler.FindUserByUserId())

	handler.DI()
	log.Println("Server running...")
}
