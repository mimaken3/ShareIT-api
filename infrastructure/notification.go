package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type notificationInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewNotificationDB(db *gorm.DB) repository.NotificationRepository {
	return &notificationInfraStruct{db: db}
}

// ユーザの通知一覧を取得
func (notificationRepo *notificationInfraStruct) FindAllNotificationsByUserID(userID uint) (notifications []model.Notification, err error) {
	return
}

// 最後の通知IDを取得
func (notificationRepo *notificationInfraStruct) FindLastNotificationID() (lastNotificationID uint, err error) {
	notification := model.Notification{}

	// SELECT notification_id FROM notifications ORDER BY notification_id DESC LIMIT 1;
	if result := notificationRepo.db.Select("notification_id").Last(&notification); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}

	lastNotificationID = notification.NotificationID
	return
}

// 通知を追加
func (notificationRepo *notificationInfraStruct) CreateNotification(sourceUserID uint, notificationType uint, typeID uint, articleID uint, lastNotificationID uint) (notificationID uint, err error) {
	var notification model.Notification
	var article model.Article

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()
	notificationID = lastNotificationID + 1

	notification.NotificationID = notificationID
	notification.SourceUserID = sourceUserID
	notification.NotificationType = notificationType

	if notificationType == 1 {
		// いいねの通知を作成する場合

		// いいねされた記事を作成したユーザIDを取得
		var like model.Like
		result := notificationRepo.db.Raw(`
select 
  created_user_id 
from 
  articles 
where 
  article_id = (
    select 
      article_id 
    from 
      likes 
    where 
      like_id = ? 
  )
;
`, typeID).Scan(&article)

		if result.Error != nil {
			// レコードがない場合
			// TODO: Do something
			// err = result.Error
		}

		// 過去にいいねしていたら、作成日と既読を更新
		var registeredNotification model.Notification
		result2 := notificationRepo.db.Raw(`
select 
  * 
from 
  notifications 
where 
  user_id = ? 
  and source_user_id = ? 
  and article_id = ? 
  and like_id = ? 
;
`, article.CreatedUserID, sourceUserID, articleID, typeID).Scan(&registeredNotification)

		// いいねに通知IDを付与
		var _like model.Like
		_like.LikeID = typeID

		if result2.Error == nil {
			// 更新
			notificationRepo.db.Model(&registeredNotification).
				Updates(map[string]interface{}{"is_read": 0, "created_date": currentDate})

			// いいねに通知IDを付与
			notificationRepo.db.Model(&_like).Update("notification_id", registeredNotification.NotificationID)

			return
		}

		// いいねに通知IDを付与
		notificationRepo.db.Model(&_like).Update("notification_id", notificationID)

		notification.UserID = article.CreatedUserID
		notification.SourceUserID = sourceUserID
		notification.ArticleID = articleID
		notification.LikeID = typeID

		// いいねと通知IDを紐付け
		notificationRepo.db.Model(&like).Where("like_id = ?", typeID).Update("notification_id", notificationID)
	} else if notificationType == 2 {
		// コメントの通知を作成する場合
		notification.ArticleID = articleID
		notification.CommentID = typeID
	}
	notification.CreatedDate = currentDate

	// 保存
	notificationRepo.db.Create(&notification)

	return
}
