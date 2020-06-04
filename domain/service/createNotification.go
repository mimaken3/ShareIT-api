package service

// 通知を追加
func (n *notificationServiceStruct) CreateNotification(sourceUserID uint, notificationType uint, typeID uint, articleID uint) (notificationID uint, err error) {
	// 最後の通知IDを取得
	lastNotificationID, _ := n.notificationRepo.FindLastNotificationID()

	// 最後の通知元情報IDを取得
	lastDestinationID, _ := n.notificationRepo.FindLastDestinationID()

	// 通知を追加
	notificationID, err = n.notificationRepo.CreateNotification(sourceUserID, notificationType, typeID, articleID, lastNotificationID, lastDestinationID)
	return
}
