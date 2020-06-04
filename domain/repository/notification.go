package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// NotificationRepository is interface for infrastructure
type NotificationRepository interface {
	// ユーザの通知一覧を取得
	FindAllNotificationsByUserID(userID uint) (notifications []model.Notification, err error)

	// 最後の通知IDを取得
	FindLastNotificationID() (lastNotificationID uint, err error)

	// 最後の通知元情報IDを取得
	FindLastDestinationID() (lastDestinationID uint, err error)

	// 通知を追加
	CreateNotification(sourceUserID uint, notificationType uint, typeID uint, articleID uint, lastNotificationID uint, lastDestinationID uint) (notificationID uint, err error)
}
