package service

// 記事のコメントを全削除
func (c *commentServiceStruct) DeleteCommentByArticleID(articleID uint) (err error) {
	err = c.commentRepo.DeleteCommentByArticleID(articleID)
	return
}
