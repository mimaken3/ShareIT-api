package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// トピックを登録
func (a *topicServiceStruct) CreateTopic(createTopic model.Topic) (createdTopic model.Topic, err error) {
	// 最後のトピックIDを取得
	lastTopicID, err := a.topicRepo.FindLastTopicID()

	createdTopic, err = a.topicRepo.CreateTopic(createTopic, lastTopicID)

	if err != nil {
		log.Println(err)
	}
	return
}
