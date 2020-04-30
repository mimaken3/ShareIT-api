package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// コメント作成
func (c *commentServiceStruct) CreateComment(createComment model.Comment) (createdComment model.ResponseComment, err error) {
	// 最後のコメントIDを取得
	lastCommentID, err := c.commentRepo.FindLastCommentID()

	createdComment, err = c.commentRepo.CreateComment(createComment, lastCommentID)

	// 署名付きURLを取得
	createdComment.IconName, err = GetPreSignedURL(createdComment.IconName)

	return
}
