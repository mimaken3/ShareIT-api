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
		users, err := userService.FindAllUsersService()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, users)
	}
}

// ユーザ登録のチェック
func CheckUserInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		checkUser := model.User{}
		resultUserInfo := model.CheckUserInfo{}
		if err := c.Bind(&checkUser); err != nil {
			return err
		}
		resultUserInfo, _ = userService.CheckUserInfoService(checkUser)

		return c.JSON(http.StatusOK, resultUserInfo)
	}
}

// ユーザを取得
func FindUserByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ユーザIDを取得
		userId, _ := strconv.Atoi(c.Param("user_id"))
		user, err := userService.FindUserByUserIdService(userId)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
}

// ユーザを登録
func SignUpUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}
		c.Bind(&user)

		// ユーザを登録
		signUpedUser, err := userService.SignUpUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, signUpedUser)
		}

		// トピックを登録
		userInterestedTopicService.CreateUserTopic(signUpedUser.InterestedTopics, signUpedUser.UserID)

		return c.JSON(http.StatusOK, signUpedUser)
	}
}

// 最後のユーザIDを取得
func FindLastUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		lastUserId, err := userService.FindLastUserId()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, lastUserId)
	}
}
