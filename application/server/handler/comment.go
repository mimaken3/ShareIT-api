package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// コメント作成
func CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		createComment := model.Comment{}
		if err := c.Bind(&createComment); err != nil {
			return err
		}

		createdComment, err := commentService.CreateComment(createComment)

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, createdComment)
	}
}

// 記事のコメント一覧取得
func FindAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 記事IDを取得
		_articleID, _ := strconv.Atoi(c.Param("article_id"))
		articleID := uint(_articleID)

		var comments []model.Comment
		comments, err := commentService.FindAllComments(articleID)

		if err != nil {
			// if err == "record not found" {
			// 	return c.JSON(http.StatusOK, err.Error)
			// } else {
			return c.String(http.StatusBadRequest, err.Error())
			// }
		}

		return c.JSON(http.StatusOK, comments)
	}
}

// コメントを編集
// func UpdateComment() echo.HandlerFunc{
// 	return func(c echo.Context) error {
//
// 	}
// }

// コメントを削除
// func DeleteComment() echo.HandlerFunc{
// 	return func(c echo.Context) error {
//
// 	}
// }
