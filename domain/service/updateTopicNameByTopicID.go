package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// トピック名を更新
func (t *topicServiceStruct) UpdateTopicNameByTopicID(topic model.Topic) (updatedTopic model.Topic, err error) {
	updatedTopic, err = t.topicRepo.UpdateTopicNameByTopicID(topic)
	return
}
