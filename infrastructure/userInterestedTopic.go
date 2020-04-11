package infrastructure

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type userInterestedTopicInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewUserInterestedTopicDB(db *gorm.DB) repository.UserInterestedTopicRepository {
	return &userInterestedTopicInfraStruct{db: db}
}

// 最後のIDを取得
func (uiRepo *userInterestedTopicInfraStruct) GetLastID() (lastID int, err error) {
	ui := model.UserInterestedTopic{}
	uiRepo.db.Select("user_interested_topics_id").Last(&ui)
	lastID = int(ui.UserInterestedTopicsID)

	return
}

// 追加(複数ok)
func (uiRepo *userInterestedTopicInfraStruct) CreateUserTopic(topicStr string, lastID int, userID uint) (err error) {
	ui := model.UserInterestedTopic{}
	topicsArr := strings.Split(topicStr, "/")

	var topicIDArr []uint
	for _, topicName := range topicsArr {
		var topicID uint

		err := uiRepo.db.Table("topics").Where("topic_name = ?", topicName).Select("topic_id").Row().Scan(&topicID)

		if err == nil {
			topicIDArr = append(topicIDArr, topicID)
		}
	}

	var insertID uint = uint(lastID)

	for _, topicID := range topicIDArr {
		insertID = insertID + 1
		if topicStr != "" {
			ui.UserInterestedTopicsID = insertID
			ui.UserID = userID
			ui.TopicID = uint(topicID)
			uiRepo.db.Create(&ui)
		}
	}
	return
}

// 更新
// UpdateUserTopic(topicArr []int) (err error)

// 削除
// DeleteUserTopic(topicArr []int) (err error)

// 削除(トピックが削除されたら)
func (uiRepo *userInterestedTopicInfraStruct) DeleteUserTopicByTopicID(topicID int) (err error) {

	// 興味トピックが1つしかない場合、それを「その他(1)」に更新
	var userInterestedTopicsIDs []uint

	rows, err := uiRepo.db.Raw("select t2.user_interested_topics_id from(select user_id, c from ("+
		"select user_id, count(*) as c from user_interested_topics group by user_id) as t where t.c = 1"+
		") as t1 inner join (select * from user_interested_topics where topic_id = ?) as t2 on t1.user_id = t2.user_id", topicID).Rows()

	defer rows.Close()
	for rows.Next() {
		var userInterestedTopicID uint
		err = rows.Scan(&userInterestedTopicID)
		if err == nil {
			userInterestedTopicsIDs = append(userInterestedTopicsIDs, userInterestedTopicID)
		}
	}

	if len(userInterestedTopicsIDs) != 0 {
		// 	// 「その他」に更新
		uiRepo.db.Table("user_interested_topics").Where("user_interested_topics_id IN (?)", userInterestedTopicsIDs).Update("topic_id", 1)
	}

	// 物理削除
	uiRepo.db.Where("topic_id = ?", topicID).Delete(&model.UserInterestedTopic{})
	return
}

// 削除(ユーザが削除されたら)
// DeleteUserTopicByUserID(userID int) (err error)

// ユーザ毎に取得
// getTopicIDByUserID(userID int) (topicIDArr []int, err error)

// トピック毎に取得
// getTopicIDByUserID(topicID int) (userIDArr []int, err error)
