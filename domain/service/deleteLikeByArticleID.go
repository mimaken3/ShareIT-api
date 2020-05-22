package service

// 記事のいいねを削除
func (l *likeServiceStruct) DeleteLikeByArticleID(articleID uint) (err error) {
	err = l.likeRepo.DeleteLikeByArticleID(articleID)
	return
}
