package handler

import (
	"net/http"

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
// func FindAllComment() echo.HandlerFunc{
// 	return func(c echo.Context) error {
//
// 	}
// }

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
