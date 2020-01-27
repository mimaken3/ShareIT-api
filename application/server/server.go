package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT-api/application/server/handler"
)

func InitRouting(e *echo.Echo) {
	// 全ユーザを取得
	e.GET("/users", handler.FindAllUsers())

	// ユーザを取得
	e.GET("/user/:user_id", handler.FindUserByUserId())

	// ユーザを登録
	e.POST("/user/signUp", handler.SignUpUser())

	// 最後のユーザIDを取得
	e.GET("/user/lastUserId", handler.FindLastUserId())

	handler.DI()
	log.Println("Server running...")
}
