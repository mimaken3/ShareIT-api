package infrastructure

import (
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
// getLastID() (lastID int, err error)

// 追加
// CreateUserTopic(topicArr []int) (err error)

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
