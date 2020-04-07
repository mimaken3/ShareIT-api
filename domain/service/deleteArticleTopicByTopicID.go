package service

// トピックに紐づく記事トピックを削除
func (a *articleTopicServiceStruct) DeleteArticleTopicByTopicID(topicID uint) (err error) {
	err = a.articleTopicRepo.DeleteArticleTopicByTopicID(topicID)
	return
}
