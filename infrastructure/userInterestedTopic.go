package infrastructure

import (
	"strconv"
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
	topicsArr := strings.Split(topicStr, ",")
	var insertID uint = uint(lastID)

	for _, topicStr := range topicsArr {
		insertID = insertID + 1
		if topicStr != "" {
			topicID, _ := strconv.Atoi(topicStr)
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
