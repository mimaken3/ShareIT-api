package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

type ArticlesResult struct {
	IsSearched   bool            `json:"is_searched"`
	SearchUser   uint            `json:"search_user"`
	SearchTopics string          `json:"search_topics"`
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
		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		userID := uint(intUserID)

		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		if refPg == 0 {
			refPg = 1
		}

		var articlesResult ArticlesResult
		articles, allPagingNum, err := articleService.FindAllArticlesService(refPg, userID)

		if err != nil {
			// １つもなかった場合
			articlesResult.IsSearched = false
			articlesResult.SearchUser = 0
			articlesResult.SearchTopics = "0"
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = articles

			return c.JSON(http.StatusOK, articlesResult)
		}

		// 各記事にいいね情報を付与
		updatedArticles, err := likeService.GetLikeInfoByArtiles(userID, articles)

		articlesResult.IsSearched = false
		articlesResult.SearchUser = 0
		articlesResult.SearchTopics = "0"
		articlesResult.IsEmpty = false
		articlesResult.RefPg = refPg
		articlesResult.AllPagingNum = allPagingNum
		articlesResult.Articles = updatedArticles

		return c.JSON(http.StatusOK, articlesResult)
	}
}

// 記事を検索
func SearchAllArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ユーザIDを取得
		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		userID := uint(intUserID)

		// ログイン中のユーザIDを取得
		_loginUserID, _ := strconv.Atoi(c.QueryParam("login_user_id"))
		loginUserID := uint(_loginUserID)

		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		// トピックを取得
		topicIDStr := c.QueryParam("topic_id")

		var articlesResult ArticlesResult
		searchedArticles, allPagingNum, err := articleService.SearchAllArticles(refPg, userID, loginUserID, topicIDStr)

		if err != nil {
			// １つもなかった場合
			articlesResult.IsSearched = true
			articlesResult.SearchUser = userID
			articlesResult.SearchTopics = topicIDStr
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = searchedArticles

			return c.JSON(http.StatusOK, articlesResult)
		}

		// 各記事にいいね情報を付与
		articles, err := likeService.GetLikeInfoByArtiles(userID, searchedArticles)

		articlesResult.IsSearched = true
		articlesResult.SearchUser = userID
		articlesResult.SearchTopics = topicIDStr
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

		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		loginUserID := uint(intUserID)

		article, err := articleService.FindArticleByArticleId(loginUserID, uint(articleId))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var sliceArticle []model.Article
		sliceArticle = append(sliceArticle, article)

		updatedArticles, err := likeService.GetLikeInfoByArtiles(loginUserID, sliceArticle)

		return c.JSON(http.StatusOK, updatedArticles[0])
	}
}

// 記事を更新
func UpdateArticleByArticleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		willBeUpdatedArticle := model.Article{}

		if err := c.Bind(&willBeUpdatedArticle); err != nil {
			return err
		}

		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		userID := uint(intUserID)

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

		var sliceArticle []model.Article
		sliceArticle = append(sliceArticle, updatedArticle)

		updatedArticles, err := likeService.GetLikeInfoByArtiles(userID, sliceArticle)
		if err != nil {
			//TODO: Badステータスを返す
			return err
		}

		return c.JSON(http.StatusOK, updatedArticles[0])
	}
}

// 記事を削除
func DeleteArticleByArticleId() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleId, _ := strconv.Atoi(c.Param("article_id"))

		// intをuintに変換
		var uintArticleId uint = uint(articleId)

		// 記事を削除
		err := articleService.DeleteArticleByArticleId(uintArticleId)

		// 記事のコメントを全削除
		err = commentService.DeleteCommentByArticleID(uintArticleId)

		// 記事のいいねを削除
		err = likeService.DeleteLikeByArticleID(uintArticleId)

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

		// ログイン中のユーザIDを取得
		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		loginUserID := uint(intUserID)

		paramUserID, _ := strconv.Atoi(c.Param("user_id"))
		// intをuintに変換
		var paramUintUserID uint = uint(paramUserID)

		if !(loginUserID == paramUintUserID && loginUserID == createArticle.CreatedUserID) {
			return c.String(http.StatusBadRequest, "URLが間違っています")
		}

		// 記事を追加
		createdArticle, _ := articleService.CreateArticle(createArticle)

		// 記事トピックを追加
		articleTopicService.CreateArticleTopic(createdArticle)

		var sliceArticle []model.Article
		sliceArticle = append(sliceArticle, createdArticle)

		// 各記事にいいね情報を付与
		updatedArticles, _ := likeService.GetLikeInfoByArtiles(loginUserID, sliceArticle)

		return c.JSON(http.StatusOK, updatedArticles[0])
	}
}

// 特定のユーザの全記事を取得
func FindArticlesByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		// ログイン中のユーザIDを取得
		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		loginUserID := uint(intUserID)

		userID, _ := strconv.Atoi(c.Param("user_id"))
		// intをuintに変換
		var uintUserID uint = uint(userID)

		if refPg == 0 {
			refPg = 1
		}

		var articlesResult ArticlesResult

		articles, allPagingNum, err := articleService.FindArticlesByUserIdService(uintUserID, loginUserID, refPg)
		if err != nil {
			articlesResult.IsSearched = false
			articlesResult.SearchUser = 0
			articlesResult.SearchTopics = "0"
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = articles

			return c.JSON(http.StatusOK, articlesResult)
		}

		// 各記事にいいね情報を付与
		updatedArticles, err := likeService.GetLikeInfoByArtiles(loginUserID, articles)

		articlesResult.IsSearched = false
		articlesResult.SearchUser = 0
		articlesResult.SearchTopics = "0"
		articlesResult.IsEmpty = false
		articlesResult.RefPg = refPg
		articlesResult.AllPagingNum = allPagingNum
		articlesResult.Articles = updatedArticles

		return c.JSON(http.StatusOK, articlesResult)
	}
}

// 特定のユーザのいいねした記事を取得(ページング)
func FindAllLikedArticlesByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		if refPg == 0 {
			refPg = 1
		}

		// ログイン中のユーザIDを取得
		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		loginUserID := uint(intUserID)

		// 取得する記事のユーザIDを取得
		_userID, _ := strconv.Atoi(c.Param("user_id"))
		userID := uint(_userID)

		articles, allPagingNum, err := articleService.FindAllLikedArticlesByUserID(userID, loginUserID, refPg)
		var articlesResult ArticlesResult

		if err != nil {
			articlesResult.IsSearched = false
			articlesResult.SearchUser = loginUserID
			articlesResult.SearchTopics = "0"
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = articles

			return c.JSON(http.StatusOK, articlesResult)
		}

		// 各記事にいいね情報を付与
		updatedArticles, err := likeService.GetLikeInfoByArtiles(loginUserID, articles)

		articlesResult.IsSearched = false
		articlesResult.SearchUser = loginUserID
		articlesResult.SearchTopics = "0"
		articlesResult.IsEmpty = false
		articlesResult.RefPg = refPg
		articlesResult.AllPagingNum = allPagingNum
		articlesResult.Articles = updatedArticles

		return c.JSON(http.StatusOK, articlesResult)
	}
}

// 特定のトピックを含む記事を取得
func FindArticlesByTopicId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ページング番号を取得
		refPg, _ := strconv.Atoi(c.QueryParam("ref_pg"))

		if refPg == 0 {
			refPg = 1
		}

		// ログイン中のユーザIDを取得
		intUserID, _ := strconv.Atoi(c.QueryParam("user_id"))
		loginUserID := uint(intUserID)

		_topicID := c.Param("topic_id")
		topicID, _ := strconv.Atoi(_topicID)
		var uintTopicID uint = uint(topicID)

		// 指定したトピックを含む記事のIDを取得
		var articleIds []model.ArticleTopic
		articleIds, _ = articleService.FindArticleIdsByTopicIdService(uintTopicID)

		var articlesResult ArticlesResult

		articles, allPagingNum, err := articleService.FindArticlesByTopicIdService(articleIds, loginUserID, refPg)
		if err != nil {
			articlesResult.IsSearched = false
			articlesResult.SearchUser = 0
			articlesResult.SearchTopics = _topicID
			articlesResult.IsEmpty = true
			articlesResult.AllPagingNum = allPagingNum
			articlesResult.Articles = articles

			return c.JSON(http.StatusOK, articlesResult)
		}

		// 各記事にいいね情報を付与
		updatedArticles, err := likeService.GetLikeInfoByArtiles(loginUserID, articles)

		articlesResult.IsSearched = false
		articlesResult.SearchUser = 0
		articlesResult.SearchTopics = _topicID
		articlesResult.IsEmpty = false
		articlesResult.RefPg = refPg
		articlesResult.AllPagingNum = allPagingNum
		articlesResult.Articles = updatedArticles

		return c.JSON(http.StatusOK, articlesResult)
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
