package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

type APIResult struct {
	Token   string     `json:"token"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
	User    model.User `json:"user"`
}

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

		// プロフィールを登録
		err = profileService.CreateProfileByUserID(user.Profile, signUpedUser.UserID)

		return c.JSON(http.StatusOK, signUpedUser)
	}
}

// ログイン
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		if err := c.Bind(&user); err != nil {
			return err
		}

		message, resultUser, _ := userService.Login(user)
		// TODO: err を使用

		var api APIResult
		if message == "success" {
			// 成功

			// Set claims
			claims := &jwtCustomClaims{
				resultUser.UserID,
				resultUser.UserName,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				},
			}

			// tokenを作成
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString(signingKey)

			if err != nil {
				return err
			}

			api.Token = t
			api.Code = 200
			api.Message = message
			api.User = resultUser

			return c.JSON(http.StatusOK, api)
		}
		// 失敗
		api.Token = ""
		api.Code = 500
		api.Message = message
		api.User = resultUser

		return c.JSON(http.StatusBadRequest, api)
	}
}

// ユーザを更新
func UpdateUserByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		willBeUpdatedUser := model.User{}

		if err := c.Bind(&willBeUpdatedUser); err != nil {
			return err
		}

		// ユーザIDを取得
		userID, _ := strconv.Atoi(c.Param("user_id"))

		// パラメータのIDと受け取ったモデルのIDが違う場合、エラーを返す
		if uint(userID) != willBeUpdatedUser.UserID {
			return c.String(http.StatusBadRequest, "param userID and send user id are different")
		}

		// 興味トピックの末尾に/があったらそれを削除
		interestedTopics := willBeUpdatedUser.InterestedTopics
		if strings.LastIndex(interestedTopics, "/") == len(interestedTopics)-1 {
			willBeUpdatedUser.InterestedTopics = strings.TrimSuffix(interestedTopics, "/")
		}

		// 興味トピックが更新されているか確認
		isUpdatedInterestedTopic, err := userService.CheckUpdateInterestedTopic(willBeUpdatedUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if isUpdatedInterestedTopic {
			// 興味トピックを更新
			err = userInterestedTopicService.UpdateUserInterestedTopic(willBeUpdatedUser)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
		}

		// プロフィールを更新
		profileService.UpdateProfileByUserID(willBeUpdatedUser.Profile, willBeUpdatedUser.UserID)

		// 更新日を更新
		updatedUser, err := userService.UpdateUser(willBeUpdatedUser.UserID)

		return c.JSON(http.StatusOK, updatedUser)

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
