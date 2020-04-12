package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// TopicRepository is interface for infrastructure
type TopicRepository interface {
	// 最後のトピックIDを取得
	FindLastTopicID() (lastTopicID uint, err error)

	// トピック名の重複チェック
	CheckTopicName(topicName string) (isDuplicated bool, message string, err error)

	// トピックを登録
	CreateTopic(createTopic model.Topic, lastTopicID uint) (createdTopic model.Topic, err error)

	// 全トピックを取得
	FindAllTopics() (topics []model.Topic, err error)

	// トピックを削除
	DeleteTopicByTopicID(uintTopicID uint) (err error)
}
