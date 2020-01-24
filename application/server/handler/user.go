package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// 全ユーザを取得
func FindAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, _ := userService.FindAllUsersService()
		return c.JSON(http.StatusOK, users)
	}
}

// ユーザを取得
func FindUserByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ユーザIDを取得
		userId, _ := strconv.Atoi(c.Param("user_id"))
		user, _ := userService.FindUserByUserIdService(userId)
		return c.JSON(http.StatusOK, user)
	}
}
