package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ユーザの通知一覧を取得
func FindAllNotificationsByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		a := c.Param("user_id")
		// intUserID, _ := strconv.Atoi(c.Param("user_id"))
		// userID := uint(intUserID)

		return c.String(http.StatusOK, a)
	}
}
