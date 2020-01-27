package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT-api/domain/model"
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

// ユーザを登録
func SignUpUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}
		c.Bind(&user)
		signUpedUser, err := userService.SignUpUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, signUpedUser)
		}
		return c.JSON(http.StatusOK, signUpedUser)
	}
}

// 最後のユーザIDを取得
func FindLastUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		lastUserId, err := userService.FindLastUserId()
		if err != nil {
			return c.JSON(http.StatusBadRequest, lastUserId)
		}
		return c.JSON(http.StatusOK, lastUserId)
	}
}
