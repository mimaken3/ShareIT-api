package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// ユーザが作成したトピックを取得
func (t *topicServiceStruct) FindCreatedTopicsByUserID(userID uint) (topics []model.Topic, err error) {
	topics, err = t.topicRepo.FindCreatedTopicsByUserID(userID)
	return
}
