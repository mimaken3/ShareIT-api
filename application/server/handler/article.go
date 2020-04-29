package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

type ArticlesResult struct {
	IsEmpty      bool            `json:"is_empty"`
	RefPg        int             `json:"ref_pg"`
	AllPagingNum int             `json:"all_paging_num"`
	Articles     []model.Article `json:"articles"`
}

// テストレスポンスを返す
func TestResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "成功！！")
	}
}

// 全記事を取得(ページング)
func FindAllArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		if refPg == 0 {
			refPg = 1
		}

		var articlesResult ArticlesResult
		articles, allPagingNum, err := articleService.FindAllArticlesService(refPg)

		if err != nil {
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = articles

			return c.JSON(http.StatusOK, articlesResult)
		}

		articlesResult.IsEmpty = false
		articlesResult.RefPg = refPg
		articlesResult.AllPagingNum = allPagingNum
		articlesResult.Articles = articles

		return c.JSON(http.StatusOK, articlesResult)
	}
}

// 記事を取得
func FindArticleByArticleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 記事IDを取得
		articleId, _ := strconv.Atoi(c.Param("article_id"))
		article, err := articleService.FindArticleByArticleId(uint(articleId))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, article)
	}
}

// 記事を更新
func UpdateArticleByArticleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		willBeUpdatedArticle := model.Article{}

		if err := c.Bind(&willBeUpdatedArticle); err != nil {
			return err
		}

		// 記事IDを取得
		articleID, _ := strconv.Atoi(c.Param("article_id"))

		// パラメータのIDと受け取ったモデルのIDが違う場合、エラーを返す
		if uint(articleID) != willBeUpdatedArticle.ArticleID {
			return c.String(http.StatusBadRequest, "param article_id and send article id are different")
		}

		// 記事トピックの末尾に/があったらそれを削除
		articleTopics := willBeUpdatedArticle.ArticleTopics
		if strings.LastIndex(articleTopics, "/") == len(articleTopics)-1 {
			willBeUpdatedArticle.ArticleTopics = strings.TrimSuffix(articleTopics, "/")
		}

		// 記事トピックが更新されているか確認
		isUpdatedArticleTopic := articleService.CheckUpdateArticleTopic(willBeUpdatedArticle)

		if isUpdatedArticleTopic {
			// 記事トピックを更新
			articleTopicService.UpdateArticleTopic(willBeUpdatedArticle)
		}

		// 記事を更新
		updatedArticle, err := articleService.UpdateArticleByArticleId(willBeUpdatedArticle)

		if err != nil {
			//TODO: Badステータスを返す
			return err
		}

		return c.JSON(http.StatusOK, updatedArticle)
	}
}

// 記事を削除
func DeleteArticleByArticleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleId, _ := strconv.Atoi(c.Param("article_id"))

		// intをuintに変換
		var uintArticleId uint = uint(articleId)

		err := articleService.DeleteArticleByArticleId(uintArticleId)

		// TODO: 記事トピックは削除しない？
		// if err == nil {
		// 	// 記事トピックを削除
		// 	willBeDeletedArticle := model.Article{ArticleID: uintArticleId}
		// 	articleTopicService.DeleteArticleTopic(willBeDeletedArticle)
		// 	// TODO: 記事トピック削除が失敗したらエラーを返す
		//
		// 	// 削除に成功したら
		// 	return c.String(http.StatusOK, "Successfully deleted article")
		// } else {
		// 	// 削除に失敗したら
		// 	return c.String(http.StatusBadRequest, err.Error())
		// }

		if err == nil {
			// 削除に成功したら
			return c.String(http.StatusOK, "Successfully deleted article")
		} else {
			// 	// 削除に失敗したら
			return c.String(http.StatusBadRequest, err.Error())
		}
	}
}

// 記事を投稿
func CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		createArticle := model.Article{}
		if err := c.Bind(&createArticle); err != nil {
			return err
		}
		// 記事を追加
		createdArticle, _ := articleService.CreateArticle(createArticle)

		// 記事トピックを追加
		articleTopicService.CreateArticleTopic(createdArticle)

		return c.JSON(http.StatusOK, createdArticle)
	}
}

// 特定のユーザの全記事を取得
func FindArticlesByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		userID, _ := strconv.Atoi(c.Param("user_id"))
		// intをuintに変換
		var uintUserID uint = uint(userID)

		if refPg == 0 {
			refPg = 1
		}

		var articlesResult ArticlesResult

		articles, allPagingNum, err := articleService.FindArticlesByUserIdService(uintUserID, refPg)
		if err != nil {
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = articles

			return c.JSON(http.StatusOK, articlesResult)
		}

		articlesResult.IsEmpty = false
		articlesResult.RefPg = refPg
		articlesResult.AllPagingNum = allPagingNum
		articlesResult.Articles = articles

		return c.JSON(http.StatusOK, articlesResult)
	}
}

// 特定のトピックを含む記事を取得
func FindArticlesByTopicId() echo.HandlerFunc {
	return func(c echo.Context) error {
		topicID, _ := strconv.Atoi(c.Param("topic_id"))
		var uintTopicID uint = uint(topicID)

		// 指定したトピックを含む記事のIDを取得
		var articleIds []model.ArticleTopic
		articleIds, _ = articleService.FindArticleIdsByTopicIdService(uintTopicID)

		articles, err := articleService.FindArticlesByTopicIdService(articleIds)
		if err != nil {
			// ない場合
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, articles)
	}
}

// 指定したトピックを含む記事トピックを取得
func FindArticleIdsByTopicId() echo.HandlerFunc {
	return func(c echo.Context) error {
		topicID, _ := strconv.Atoi(c.Param("topic_id"))
		// intをuintに変換
		var uintTopicID uint = uint(topicID)

		var articleIds []model.ArticleTopic
		articleIds, err := articleService.FindArticleIdsByTopicIdService(uintTopicID)
		if err != nil {
			// ない場合
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, articleIds)
	}
}

// 最後の記事IDを取得
func FindLastArticleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		lastArticleId, err := articleService.FindLastArticleId()
		if err != nil {
			// ない場合
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, lastArticleId)
	}
}
