package infrastructure

import (
	"errors"

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
func (notificationRepo *notificationInfraStruct) FindAllNotificationsByUserID(userID uint) (resultNotifications []model.ResultNotification, err error) {
	rows, err :=
		notificationRepo.db.Raw(`
select 
  sub_n.notification_id, 
  sub_n.user_id, 
  sub_n.user_name as source_user_name, 
  sub_n.source_user_id, 
  sub_n.source_user_icon_name, 
  sub_n.is_read, 
  d.destination_type_id, 
  d.destination_type_name_id, 
  d.behavior_type_id, 
  d.behavior_type_name_id, 
  sub_n.created_date 
from 
  (
    select 
      n.*, 
      u.user_name 
    from 
      (
        select 
          * 
        from 
          notifications 
        where 
          notifications.user_id = ? 
          and notifications.source_user_id != ? 
      ) as n 
      left join users as u on n.source_user_id = u.user_id
  ) as sub_n 
  left join destinations as d on sub_n.destination_id = d.destination_id 
order by 
  sub_n.created_date desc
;
		`, userID, userID).Rows()

	defer rows.Close()
	for rows.Next() {
		resultNotification := model.ResultNotification{}
		err = notificationRepo.db.ScanRows(rows, &resultNotification)
		if err == nil {
			resultNotifications = append(resultNotifications, resultNotification)
		}
	}

	// レコードがない場合
	if len(resultNotifications) == 0 {
		return nil, errors.New("record not found")
	}

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
// typeIDはlike_id or comment_id
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
  and destination_id in (
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

		// コメントされた記事を作成したユーザIDを取得
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
      comments 
    where 
      comment_id = ? 
  )
;
`, typeID).Scan(&article)

		if result.Error != nil {
			// レコードがない場合
			// TODO: Do something
			// err = result.Error
		}
		createdArticleUserID := article.CreatedUserID

		// コメントに通知IDを付与するための宣言
		var comment model.Comment
		comment.CommentID = typeID

		// いいねに通知IDを付与
		notificationRepo.db.Model(&comment).Update("notification_id", notificationID)

		notification.UserID = createdArticleUserID
		notification.SourceUserID = sourceUserID

		// 目的地を保存
		destination.DestinationTypeID = 1 // 1: ユーザ詳細画面
		destination.DestinationTypeNameID = articleID
		destination.BehaviorTypeID = 2 // 2: コメントをする
		destination.BehaviorTypeNameID = typeID
		notificationRepo.db.Create(&destination)
	}

	// 保存
	notificationRepo.db.Create(&notification)

	return
}

// 通知の未読を既読にする
func (notificationRepo *notificationInfraStruct) ReadNotificationByNotificationID(notificationID uint) (resultNotification model.ResultNotification, err error) {
	var notification model.Notification
	notification.NotificationID = notificationID
	notificationRepo.db.Model(&notification).Update("is_read", int8(1))

	notificationRepo.db.Raw(`
select 
  n.notification_id, 
  n.user_id, 
  u.user_name as source_user_name, 
  n.source_user_id, 
  n.source_user_icon_name, 
  n.is_read, 
  d.destination_type_id, 
  d.destination_type_name_id, 
  d.behavior_type_id, 
  d.behavior_type_name_id,
  n.created_date
from 
  notifications as n 
  left join users as u on n.source_user_id = u.user_id 
  left join destinations as d on n.destination_id = d.destination_id 
where 
  notification_id = ? 
;
	`, notificationID).Scan(&resultNotification)

	return
}
