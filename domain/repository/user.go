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

	// 最後のユーザIDを取得
	FindLastUserId() (lastUserId uint, err error)
}
