package infrastructure

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type commentInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewCommentDB(db *gorm.DB) repository.CommentRepository {
	return &commentInfraStruct{db: db}
}

// 最後のコメントIDを取得
func (commentRepo *commentInfraStruct) FindLastCommentID() (lastCommentID uint, err error) {
	comment := model.Comment{}
	if result := commentRepo.db.Select("comment_id").Last(&comment); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}

	lastCommentID = comment.CommentID
	return
}

// コメント作成
func (commentRepo *commentInfraStruct) CreateComment(createComment model.Comment, lastCommentID uint) (createdComment model.Comment, err error) {
	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	nowTime := time.Now().Format(dateFormat)
	customisedNowTime, _ := time.Parse(dateFormat, nowTime)

	const defaultDeletedDateStr = "9999-12-31 23:59:59"
	defaultDeletedDate, _ := time.Parse(dateFormat, defaultDeletedDateStr)

	// DBに保存する記事のモデルを作成
	createdComment.CommentID = lastCommentID + 1
	createdComment.ArticleID = createComment.ArticleID
	createdComment.UserID = createComment.UserID
	createdComment.Content = createComment.Content
	createdComment.CreatedDate = customisedNowTime
	createdComment.UpdatedDate = customisedNowTime
	createdComment.DeletedDate = defaultDeletedDate

	commentRepo.db.Create(&createdComment)

	return
}

// 記事のコメント一覧取得
func (commentRepo *commentInfraStruct) FindAllComments(articleID uint) (comments []model.Comment, err error) {
	// SELECT * FROM comments WHERE is_deleted = 0 AND article_id = ?;
	commentRepo.db.Where("is_deleted = 0 AND article_id = ?", articleID).Find(&comments)
	if len(comments) == 0 {
		return []model.Comment{}, errors.New("record not found")
	}

	return
}

// コメントを編集
func (commentRepo *commentInfraStruct) UpdateComment(updateComment model.Comment) (updatedComment model.Comment, err error) {
	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	updateTime := time.Now().Format(dateFormat)
	customisedUpdateTime, _ := time.Parse(dateFormat, updateTime)

	commentID := updateComment.CommentID

	// 更新するフィールドを設定
	updateContent := updateComment.Content

	// 更新
	commentRepo.db.Model(&updatedComment).
		Where("comment_id = ?", commentID).
		Updates(map[string]interface{}{
			"content":      updateContent,
			"updated_date": customisedUpdateTime,
		})

	updatedComment.CommentID = updateComment.CommentID
	updatedComment.ArticleID = updateComment.ArticleID
	updatedComment.UserID = updateComment.UserID

	return
}

// コメントを削除
// func (commentRepo *commentInfraStruct) DeleteComment(commentID uint) (err error) {
// 	return
// }
