package service

// ユーザのコメントを全削除
func (c *commentServiceStruct) DeleteCommentByUserID(userID uint) (err error) {
	err = c.commentRepo.DeleteCommentByUserID(userID)
	return
}
