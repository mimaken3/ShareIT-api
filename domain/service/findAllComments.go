package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事のコメント一覧取得
func (c *commentServiceStruct) FindAllComments(articleID uint) (comments []model.ResponseComment, err error) {
	comments, err = c.commentRepo.FindAllComments(articleID)

	for i := 0; i < len(comments); i++ {
		// 署名付きURLを取得
		comments[i].IconName, err = GetPreSignedURL(comments[i].IconName)
	}

	return
}
