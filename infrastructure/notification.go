package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type notificationInfraStruct struct {
	db *gorm.DB
}

type checkDuplicatedNotification struct {
	NotificationID        uint `gorm:"primary_key" json:"notification_id"`
	UserID                uint `json:"user_id"`
	SourceUserID          uint `json:"source_user_id"`
	DestinationTypeID     uint `json:"destination_type_id"`
	DestinationTypeNameID uint `json:"destination_type_name_id"`
	BehaviorTypeID        uint `json:"behavior_type_id"`
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

// 最後の通知元情報IDを取得
func (notificationRepo *notificationInfraStruct) FindLastDestinationID() (lastDestinationID uint, err error) {
	destination := model.Destination{}

	if result := notificationRepo.db.Select("destination_id").Last(&destination); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}

	lastDestinationID = destination.DestinationID

	return
}

// 通知を追加
func (notificationRepo *notificationInfraStruct) CreateNotification(sourceUserID uint, notificationType uint, typeID uint, articleID uint, lastNotificationID uint, lastDestinationID uint) (notificationID uint, err error) {
	var notification model.Notification
	var destination model.Destination
	var article model.Article

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()
	notificationID = lastNotificationID + 1
	destinationID := lastDestinationID + 1

	// 通知
	notification.NotificationID = notificationID
	notification.SourceUserID = sourceUserID
	notification.DestinationID = destinationID
	notification.CreatedDate = currentDate

	// 目的地
	destination.DestinationID = destinationID

	if notificationType == 1 {
		// いいねの通知を作成する場合

		// いいねされた記事を作成したユーザIDを取得
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
		createdArticleUserID := article.CreatedUserID

		// 過去にいいねしていたら、作成日と既読を更新のみ行う
		var registeredNotification model.Notification
		result2 := notificationRepo.db.Raw(`
select 
  * 
from 
  notifications 
where 
  user_id = ? 
  and source_user_id = ? 
  and destination_id = (
    select 
      destination_id 
    from 
      destinations 
    where 
      destination_type_id = 1 
      and behavior_type_id = 1 
      and destination_type_name_id = ? 
  )
;
`, createdArticleUserID, sourceUserID, articleID).Scan(&registeredNotification)

		// いいねに通知IDを付与するための宣言
		var like model.Like
		like.LikeID = typeID

		if result2.Error == nil {
			// 更新
			notificationRepo.db.Model(&registeredNotification).
				Updates(map[string]interface{}{"is_read": 0, "created_date": currentDate})

			// いいねに通知IDを付与
			// 前のいいねはDBから消えてるので、新しいいいねに既に存在していた通知IDを付与
			notificationRepo.db.Model(&like).Update("notification_id", registeredNotification.NotificationID)

			return
		}

		// いいねに通知IDを付与
		notificationRepo.db.Model(&like).Update("notification_id", notificationID)

		notification.UserID = createdArticleUserID
		notification.SourceUserID = sourceUserID

		// 目的地を保存
		destination.DestinationTypeID = 1 // 1: ユーザ詳細画面
		destination.DestinationTypeNameID = articleID
		destination.BehaviorTypeID = 1 // 1: いいねをする
		destination.BehaviorTypeNameID = typeID
		notificationRepo.db.Create(&destination)

		// いいねと通知IDを紐付け いる？
		// notificationRepo.db.Model(&like).Where("like_id = ?", typeID).Update("notification_id", notificationID)
	} else if notificationType == 2 {
		// コメントの通知を作成する場合
		// notification.ArticleID = articleID
		// notification.CommentID = typeID
	}

	// 保存
	notificationRepo.db.Create(&notification)

	return
}
