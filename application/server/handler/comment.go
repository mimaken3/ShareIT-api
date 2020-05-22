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

		var comments []model.ResponseComment
		comments, err := commentService.FindAllComments(articleID)

		if err != nil {
			return c.JSON(http.StatusOK, []model.Comment{})
		}

		return c.JSON(http.StatusOK, comments)
	}
}

// コメントを編集
func UpdateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		updateComment := model.Comment{}
		if err := c.Bind(&updateComment); err != nil {
			return err
		}

		updatedComment, err := commentService.UpdateComment(updateComment)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, updatedComment)
	}
}

// コメントを削除
func DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		// コメントIDを取得
		_commentID, _ := strconv.Atoi(c.Param("comment_id"))
		commentID := uint(_commentID)

		err := commentService.DeleteComment(commentID)

		if err == nil {
			// 削除に成功したら
			return c.String(http.StatusOK, "Successfully deleted comment")
		} else {
			// 削除に失敗したら
			return c.String(http.StatusBadRequest, err.Error())
		}

	}
}
