package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type commentServiceStruct struct {
	commentRepo repository.CommentRepository
}

// Application層はこのInterfaceに依存
type CommentServiceInterface interface {
	// コメント作成
	CreateComment(createComment model.Comment) (createdComment model.Comment, err error)

	// 記事のコメント一覧取得
	FindAllComments(articleID uint) (comments []model.Comment, err error)

	// コメントを編集
	// UpdateComment(updateComment model.Comment) (updatedComment model.Comment, err error)

	// コメントを削除
	// DeleteComment(commentID uint)(err error)
}

// DIのための関数
func NewCommentService(c repository.CommentRepository) CommentServiceInterface {
	return &commentServiceStruct{commentRepo: c}
}
