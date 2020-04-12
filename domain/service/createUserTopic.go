package service

import "log"

// 追加
func (ui *userInterestedTopicServiceStruct) CreateUserTopic(topicStr string, userID uint) (err error) {
	// 最後のIDを取得
	lastID, err := ui.userInterestedTopicRepo.GetLastID()

	err = ui.userInterestedTopicRepo.CreateUserTopic(topicStr, lastID, userID)

	if err != nil {
		log.Println(err)
	}
	return
}
