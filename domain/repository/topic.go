package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// TopicRepository is interface for infrastructure
type TopicRepository interface {
	// 最後のトピックIDを取得
	FindLastTopicID() (lastTopicID uint, err error)

	// トピックを登録
	CreateTopic(createTopic model.Topic, lastTopicID uint) (createdTopic model.Topic, err error)
}