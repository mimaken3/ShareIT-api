package service

// コメントを削除
func (c *commentServiceStruct) DeleteComment(commentID uint) (err error) {
	err = c.commentRepo.DeleteComment(commentID)
	return
}
