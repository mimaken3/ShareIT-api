package repository

import "github.com/mimaken3/ShareIT-api/domain/model"

// UserRepository is interface for infrastructure
type UserRepository interface {
	// 全ユーザを取得
	FindAllUsers() (users []model.User, err error)

	// ユーザを取得
	FindUserByUserId(userId int) (user model.User, err error)
}
