package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// UserRepository is interface for infrastructure
type UserRepository interface {
	// 全ユーザを取得
	FindAllUsers() (users []model.User, err error)

	// ユーザ登録のチェック
	CheckUserInfo(checkUser model.User) (resultUserInfo model.CheckUserInfo, err error)

	// ユーザを取得
	FindUserByUserId(userId int) (user model.User, err error)

	// ユーザを登録
	SignUpUser(user model.User, lastUserId uint) (signedUpUser model.User, err error)

	// ログイン
	Login(user model.User) (message string, resultUser model.User, err error)

	// 興味トピックが更新されているか確認
	CheckUpdateInterestedTopic(willBeUpdatedUser model.User) (isUpdatedInterestedTopic bool, err error)

	// パスワードをハッシュ化
	PasswordToHash(password string) (hashedPassword string, err error)

	// パスワードが一致するかのチェック
	// VerifyPassword(user model.User) (loginUser model.User, err error)

	// 最後のユーザIDを取得
	FindLastUserId() (lastUserId uint, err error)

	// ユーザのinterested_topicsにあるトピックを削除
	DeleteTopicFromInterestedTopics(deleteTopicID uint) (err error)
}
