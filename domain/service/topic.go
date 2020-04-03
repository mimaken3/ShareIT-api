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

	// トピックを登録
	CreateTopic(createTopic model.Topic) (createdTopic model.Topic, err error)
}

// DIのための関数
func NewTopicService(a repository.TopicRepository) TopicServiceInterface {
	return &topicServiceStruct{topicRepo: a}
}
