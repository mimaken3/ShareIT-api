package service

// トピックを削除
func (a *topicServiceStruct) DeleteTopicByTopicID(uintTopicID uint) (err error) {
	// a.articleTopicRepo.DeleteArticleTopic(willBeDeletedArticle)

	// usersにあるトピックを削除
	// articlesにあるトピックを削除

	// article_topicsにあるトピックを削除
	// willBeDeletedArticle := model.ArticleTopic{}
	// willBeDeletedArticle.TopicID = uintTopicID
	// at.articleTopicRepo.DeleteArticleTopic(willBeDeletedArticle)

	// topicsにあるトピックを削除
	err = a.topicRepo.DeleteTopicByTopicID(uintTopicID)

	return
}
