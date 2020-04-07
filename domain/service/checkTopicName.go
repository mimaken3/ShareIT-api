package service

// トピック名の重複チェック
func (t *topicServiceStruct) CheckTopicName(topicName string) (isDuplicated bool, message string, err error) {
	isDuplicated, message, err = t.topicRepo.CheckTopicName(topicName)

	return
}
