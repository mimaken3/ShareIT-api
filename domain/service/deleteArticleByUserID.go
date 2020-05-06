package service

// ユーザの記事を全削除
func (a *articleServiceStruct) DeleteArticleByUserID(userID uint) (err error) {
	err = a.articleRepo.DeleteArticleByUserID(userID)
	return
}
