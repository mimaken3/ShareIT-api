package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 通知の未読を既読にする
func (n *notificationServiceStruct) ReadNotificationByNotificationID(notificationID uint) (resultNotification model.ResultNotification, err error) {
	resultNotification, err = n.notificationRepo.ReadNotificationByNotificationID(notificationID)
	return
}
