package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 全トピックを取得
func (t *topicServiceStruct) FindAllTopics() (topics []model.Topic, err error) {
	topics, err = t.topicRepo.FindAllTopics()
	if err != nil {
		log.Println(err)
	}
	return
}
