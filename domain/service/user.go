package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type userServiceStruct struct {
	userRepo repository.UserRepository
}

// Application層はこのInterfaceに依存
type UserServiceInterface interface {
	// 全ユーザを取得
	FindAllUsersService() (users []model.User, err error)

	// ユーザ登録のチェック
	CheckUserInfoService(checkUser model.User) (resultUserInfo model.CheckUserInfo, err error)

	// ユーザを取得
	FindUserByUserIdService(userId int) (user model.User, err error)

	// ユーザを登録
	SignUpUser(user model.User) (signedUpUser model.User, err error)

	// 最後のユーザIDを取得
	FindLastUserId() (lastUserId uint, err error)

	// ユーザのinterested_topicsにあるトピックを削除
	DeleteTopicFromInterestedTopics(deleteTopicID uint) (err error)
}

// DIのための関数
func NewUserService(u repository.UserRepository) UserServiceInterface {
	return &userServiceStruct{userRepo: u}
}
