package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// コメントを編集
func (c *commentServiceStruct) UpdateComment(updateComment model.Comment) (updatedComment model.Comment, err error) {
	updatedComment, err = c.commentRepo.UpdateComment(updateComment)
	return
}
