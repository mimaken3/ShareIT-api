package service

import (
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type userInterestedTopicServiceStruct struct {
	userInterestedTopicRepo repository.UserInterestedTopicRepository
}

// Application層はこのInterfaceに依存
type UserInterestedTopicServiceInterface interface {
	// 追加
	CreateUserTopic(topicStr string, userID uint) (err error)

	// 更新
	// UpdateUserTopic(topicArr []int) (err error)

	// 削除
	// DeleteUserTopic(topicArr []int) (err error)

	// 削除(トピックが削除されたら)
	DeleteUserTopicByTopicID(topicID int) (err error)

	// 削除(ユーザが削除されたら)
	// DeleteUserTopicByUserID(userID int) (err error)

	// ユーザ毎に取得
	// getTopicIDByUserID(userID int) (topicIDArr []int, err error)

	// トピック毎に取得
	// getTopicIDByUserID(topicID int) (userIDArr []int, err error)
}

// DIのための関数
func NewUserInterestedTopicService(ui repository.UserInterestedTopicRepository) UserInterestedTopicServiceInterface {
	return &userInterestedTopicServiceStruct{userInterestedTopicRepo: ui}
}
