package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT-api/application/server/handler"
)

func InitRouting(e *echo.Echo) {
	// テストレスポンスを返す
	e.GET("/test", handler.TestResponse())

	// ユーザ登録時のチェック
	e.POST("/signUp/check", handler.CheckUserInfo())

	// 全ユーザを取得
	e.GET("/users", handler.FindAllUsers())

	// ユーザを取得
	e.GET("/user/:user_id", handler.FindUserByUserId())

	// ユーザを登録
	e.POST("/user/signUp", handler.SignUpUser())

	// 最後のユーザIDを取得
	e.GET("/user/lastUserId", handler.FindLastUserId())

	// 記事を投稿
	e.POST("/user/:user_id/createArticle", handler.CreateArticle())

	// 全記事を取得(トピック: 文字列区切り)
	e.GET("/articles", handler.FindAllArticles())

	// 記事を取得(トピック: 数値区切り)
	e.GET("/article/:article_id", handler.FindArticleByArticleId())

	// 記事を取得(トピック: 文字列区切り)
	// e.GET("/article/:article_id", handler.FindArticleByArticleId())

	// 記事を更新
	e.PUT("/article/:article_id", handler.UpdateArticleByArticleId())

	// 記事を削除
	e.DELETE("/article/:article_id", handler.DeleteArticleByArticleId())

	// 特定のユーザの全記事を取得(トピック: 文字列区切り) :未実装
	e.GET("/user/:user_id/articles", handler.FindArticlesByUserId())

	// 特定のトピックを含む記事を取得
	e.GET("/articles/topic/:topic_id", handler.FindArticlesByTopicId())

	// 指定したトピックを含む記事トピックを取得
	e.GET("/articleIds/:topic_id", handler.FindArticleIdsByTopicId())

	// 最後の記事IDを取得
	e.GET("/article/lastArticleId", handler.FindLastArticleId())

	// トピックを作成
	e.POST("/topic/create", handler.CreateTopic())

	// 全トピックを取得
	e.GET("/topics", handler.FindAllTopics())

	// トピックを削除
	e.DELETE("/topic/:topic_id", handler.DeleteTopicByTopicID())

	handler.DI()
	log.Println("Server running...")
}
