package handler

import (
	"github.com/mimaken3/ShareIT-api/domain/service"
	"github.com/mimaken3/ShareIT-api/infrastructure"
)

var userService service.UserServiceInterface
var articleService service.ArticleServiceInterface
var articleTopicService service.ArticleTopicServiceInterface

func DI() {
	// ユーザ
	// DBと直接やり取りをするrepositoryにDBの情報を外部から注入
	userRepo := infrastructure.NewUserDB(infrastructure.DB)
	userService = service.NewUserService(userRepo)

	// 記事
	// DBと直接やり取りをするrepositoryにDBの情報を外部から注入
	articleRepo := infrastructure.NewArticleDB(infrastructure.DB)
	articleService = service.NewArticleService(articleRepo)

	// 記事とトピック
	// DBと直接やり取りをするrepositoryにDBの情報を外部から注入
	articleTopicRepo := infrastructure.NewArticleTopicDB(infrastructure.DB)
	articleTopicService = service.NewArticleTopicService(articleTopicRepo)
}
