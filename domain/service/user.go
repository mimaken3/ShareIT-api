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

	// ユーザを取得
	FindUserByUserIdService(userId int) (user model.User, err error)
}

// DIのための関数
func NewUserService(u repository.UserRepository) UserServiceInterface {
	return &userServiceStruct{userRepo: u}
}
