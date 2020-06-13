package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

type ResponseNotification struct {
	IsEmpty       bool                       `json:"is_empty"`
	Notifications []model.ResultNotification `json:"notifications"`
}

// ユーザの通知一覧を取得
func FindAllNotificationsByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		intUserID, _ := strconv.Atoi(c.Param("user_id"))
		userID := uint(intUserID)

		resultNotifications, _ := notificationService.FindAllNotificationsByUserID(userID)

		for i := 0; i < len(resultNotifications); i++ {
			// ユーザIDから署名付きURLを取得
			resultNotifications[i].SourceUserIconName, _ = iconService.GetPreSignedURLByUserID(resultNotifications[i].SourceUserID)
		}

		var responseNotification ResponseNotification
		if len(resultNotifications) == 0 {
			responseNotification.IsEmpty = true
		} else {
			responseNotification.IsEmpty = false
		}
		responseNotification.Notifications = resultNotifications

		return c.JSON(http.StatusOK, responseNotification)
	}
}

// 通知の未読を既読に
func ReadNotificationByNotificationID() echo.HandlerFunc {
	return func(c echo.Context) error {
		responseNotification := model.ResultNotification{}
		if err := c.Bind(&responseNotification); err != nil {
			return err
		}

		// 通知IDを取得
		notification, _ := notificationService.ReadNotificationByNotificationID(responseNotification.NotificationID)

		// ユーザIDから署名付きURLを取得
		notification.SourceUserIconName, _ = iconService.GetPreSignedURLByUserID(notification.SourceUserID)
		return c.JSON(http.StatusOK, notification)
	}
}
