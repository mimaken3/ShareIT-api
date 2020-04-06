package repository

// UserInterestedTopicRepository is interface for infrastructure
type UserInterestedTopicRepository interface {
	// 最後のIDを取得
	// getLastID() (lastID int, err error)

	// 追加
	// CreateUserTopic(topicArr []int) (err error)

	// 更新
	// UpdateUserTopic(topicArr []int) (err error)

	// 削除
	// DeleteUserTopic(topicArr []int) (err error)

	// 削除(トピックが削除されたら)
	DeleteUserTopicByTopicID(topicID int) (err error)

	// 削除(ユーザが削除されたら)
	// DeleteUserTopicByUserID(userID int) (err error)

	// ユーザ毎に取得
	// getTopicIDByUserID(userID int) (topicIDArr []int, err error)

	// トピック毎に取得
	// getTopicIDByUserID(topicID int) (userIDArr []int, err error)
}
