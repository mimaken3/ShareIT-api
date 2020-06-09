package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type topicServiceStruct struct {
	topicRepo repository.TopicRepository
}

// Application層はこのInterfaceに依存
type TopicServiceInterface interface {
	// 最後のトピックIDを取得
	// FindLastTopicID() (lastTopicID uint, err error)

	// トピック名の重複チェック
	CheckTopicName(topicName string) (isDuplicated bool, message string, err error)

	// トピックを登録
	CreateTopic(createTopic model.Topic) (createdTopic model.Topic, err error)

	// 全トピックを取得
	FindAllTopics() (topics []model.Topic, err error)

	// トピック名を更新
	UpdateTopicNameByTopicID(topic model.Topic) (updatedTopic model.Topic, err error)

	// トピックを削除
	DeleteTopicByTopicID(uintTopicID uint) (err error)

	// ユーザが作成したトピックを取得
	FindCreatedTopicsByUserID(userID uint) (topics []model.Topic, err error)
}

// DIのための関数
func NewTopicService(a repository.TopicRepository) TopicServiceInterface {
	return &topicServiceStruct{topicRepo: a}
}
