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

	// 全記事を取得
	e.GET("/articles", handler.FindAllArticles())

	// 特定のユーザの前記事を取得
	e.GET("/user/:user_id/articles", handler.FindArticlesByUserId())

	// 特定のトピックを含む記事を取得
	e.GET("/articles/:topic_id", handler.FindArticlesByTopicId())

	// 指定したトピックを含む記事のIDを取得
	e.GET("/articleIds/:topic_id", handler.FindArticleIdsByTopicId())

	handler.DI()
	log.Println("Server running...")
}
