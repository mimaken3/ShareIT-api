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
func (commentRepo *commentInfraStruct) FindAllComments(articleID uint) (comments []model.ResponseComment, err error) {
	rows, err := commentRepo.db.Raw(`
SELECT 
	c.comment_id, 
	c.article_id, 
	c.user_id, 
	u.user_name, 
	i.icon_name,
	c.content, 
	c.created_date, 
	c.updated_date, 
	c.deleted_date 
FROM 
	comments as c 
	left join users as u on (c.user_id = u.user_id) 
	left join icons as i on (c.user_id = i.user_id)
WHERE 
	c.is_deleted = 0 
	AND c.article_id = ? 
;
	`, articleID).Rows()

	defer rows.Close()
	for rows.Next() {
		comment := model.ResponseComment{}
		err = commentRepo.db.ScanRows(rows, &comment)
		if err == nil {
			comments = append(comments, comment)
		}
	}

	// レコードがない場合
	if len(comments) == 0 {
		return []model.ResponseComment{}, errors.New("record not found")
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
func (commentRepo *commentInfraStruct) DeleteComment(commentID uint) (err error) {
	deleteComment := model.Comment{}
	// SELECT * FROM comments WHERE comment_id = :commentID AND is_deleted = 0;
	if result := commentRepo.db.Find(&deleteComment, "comment_id = ? AND is_deleted = ?", commentID, 0); result.Error != nil {
		// レコードがない場合
		err = result.Error
		return
	}

	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	deleteTime := time.Now().Format(dateFormat)
	customisedDeleteTime, _ := time.Parse(dateFormat, deleteTime)

	// 削除状態に更新
	commentRepo.db.Model(&deleteComment).
		Where("comment_id = ? AND is_deleted = ?", commentID, 0).
		Updates(map[string]interface{}{
			"deleted_date": customisedDeleteTime,
			"is_deleted":   int8(1),
		})

	return nil
}
