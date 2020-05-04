package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// コメントを編集
func (c *commentServiceStruct) UpdateComment(updateComment model.Comment) (updatedComment model.ResponseComment, err error) {
	updatedComment, err = c.commentRepo.UpdateComment(updateComment)

	// 署名付きURLを取得
	updatedComment.IconName, err = GetPreSignedURL(updatedComment.IconName)

	return
}
