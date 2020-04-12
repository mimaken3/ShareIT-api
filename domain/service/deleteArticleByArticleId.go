package service

// 記事を削除
func (a *articleServiceStruct) DeleteArticleByArticleId(articleId uint) (err error) {
	return a.articleRepo.DeleteArticleByArticleId(articleId)
}
