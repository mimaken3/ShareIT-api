package service

// 削除
func (ui *userInterestedTopicServiceStruct) DeleteUserTopicByTopicID(topicID int) (err error) {
	return ui.userInterestedTopicRepo.DeleteUserTopicByTopicID(topicID)
}
