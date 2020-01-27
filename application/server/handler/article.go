package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// 全記事を取得
func FindAllArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		articles, _ := articleService.FindAllArticlesService()
		return c.JSON(http.StatusOK, articles)
	}
}
