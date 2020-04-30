package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事のコメント一覧取得
func (c *commentServiceStruct) FindAllComments(articleID uint) (comments []model.Comment, err error) {
	comments, err = c.commentRepo.FindAllComments(articleID)
	return
}
