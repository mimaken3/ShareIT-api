package service

// ユーザのinterested_topicsにあるトピックを削除
func (u *userServiceStruct) DeleteTopicFromInterestedTopics(deleteTopicID uint) (err error) {
	err = u.userRepo.DeleteTopicFromInterestedTopics(deleteTopicID)

	return
}
