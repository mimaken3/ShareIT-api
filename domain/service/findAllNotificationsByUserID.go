package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// ユーザの通知一覧を取得
func (n *notificationServiceStruct) FindAllNotificationsByUserID(userID uint) (notifications []model.Notification, err error) {
	notifications, err = n.notificationRepo.FindAllNotificationsByUserID(userID)
	return
}
