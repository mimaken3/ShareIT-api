package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// いいねON/OFF
func ToggleLikeByArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		like := model.Like{}
		if err := c.Bind(&like); err != nil {
			return err
		}

		// いいねしたユーザIDを取得
		userID := like.UserID

		// いいねした記事IDを取得
		_articleID, _ := strconv.Atoi(c.Param("article_id"))
		articleID := uint(_articleID)
		if !(articleID == like.ArticleID) {
			return c.String(http.StatusBadRequest, "URL、もしくはBodyの中身が違います")
		}

		// ページング番号を取得
		isLiked, _ := strconv.ParseBool(c.QueryParam("is_liked"))

		// いいねをトグルした後の記事を取得
		_ = likeService.ToggleLikeByArticle(userID, articleID, isLiked)

		article, err := articleService.FindArticleByArticleId(articleID)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// いいね情報を付与した記事を取得
		var sliceArticle []model.Article
		sliceArticle = append(sliceArticle, article)

		updatedArticles, err := likeService.GetLikeInfoByArtiles(userID, sliceArticle)

		return c.JSON(http.StatusOK, updatedArticles[0])
	}
}
