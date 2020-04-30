package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// CommentRepository is interface for infrastructure
type CommentRepository interface {
	// 最後のコメントIDを取得
	FindLastCommentID() (lastCommentID uint, err error)

	// コメント作成
	CreateComment(createComment model.Comment, lastCommentID uint) (createdComment model.Comment, err error)

	// 記事のコメント一覧取得
	// FindAllComment(articleID uint)(comment []model.Comment, err error)

	// コメントを編集
	// UpdateComment(updateComment model.Comment) (updatedComment model.Comment, err error)

	// コメントを削除
	// DeleteComment(commentID uint)(err error)
}
