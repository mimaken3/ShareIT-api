package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mimaken3/ShareIT-api/application/server/handler"
)

func InitRouting(e *echo.Echo) {
	// テストレスポンスを返す
	e.GET("/test", handler.TestResponse())

	// ユーザ登録時のチェック
	e.POST("/signUp/check", handler.CheckUserInfo())

	// ユーザを登録
	e.POST("/signUp", handler.SignUpUser())

	// ログイン
	e.POST("/login", handler.Login())

	// 全トピックを取得
	e.GET("/topics", handler.FindAllTopics())

	// =========
	// || API ||
	// =========
	apiG := e.Group("/api")
	apiG.Use(middleware.JWTWithConfig(handler.Config))

	//TODO: 指定したトピックを含む記事トピックを取得
	apiG.GET("/articleIds/:topic_id", handler.FindArticleIdsByTopicId())

	// ============
	// || ユーザ ||
	// ============
	userG := apiG.Group("/users")
	// 全ユーザを取得(ページング)
	userG.GET("", handler.FindAllUsers())

	// 最後のユーザIDを取得
	userG.GET("/lastUserId", handler.FindLastUserId())

	// ユーザを取得
	userG.GET("/:user_id", handler.FindUserByUserId())

	// ユーザを更新
	userG.PUT("/:user_id", handler.UpdateUserByUserId())

	// ユーザを削除
	userG.DELETE("/:user_id", handler.DeleteUser())

	// 特定のユーザの全記事を取得(トピック: 文字列区切り)
	userG.GET("/:user_id/articles", handler.FindArticlesByUserId())

	// 記事を投稿
	userG.POST("/:user_id/createArticle", handler.CreateArticle())

	// ==========
	// || 記事 ||
	// ==========
	articleG := apiG.Group("/articles")

	// 全記事を取得(ページング)(トピック: 文字列区切り)
	articleG.GET("", handler.FindAllArticles())

	// 記事を取得(トピック: 数値区切り)
	// e.GET("/article/:article_id", handler.FindArticleByArticleId())

	// 記事を取得(トピック: 文字列区切り)
	articleG.GET("/:article_id", handler.FindArticleByArticleId())

	// 記事を更新
	articleG.PUT("/:article_id", handler.UpdateArticleByArticleId())

	// 記事を削除
	articleG.DELETE("/:article_id", handler.DeleteArticleByArticleId())

	// 特定のトピックを含む記事を取得
	articleG.GET("/topic/:topic_id", handler.FindArticlesByTopicId())

	// 最後の記事IDを取得
	articleG.GET("/lastArticleId", handler.FindLastArticleId())

	// =============
	// || トピック||
	// =============
	topicG := apiG.Group("/topics")

	// トピックを作成
	topicG.POST("/create", handler.CreateTopic())

	// トピックを削除
	topicG.DELETE("/:topic_id", handler.DeleteTopicByTopicID())

	handler.DI()
	log.Println("Server running...")
}
