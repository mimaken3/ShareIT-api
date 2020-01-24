package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ユーザを取得
func (u *userServiceStruct) FindUserByUserIdService(userId int) (model.User, error) {
	user, err := u.userRepo.FindUserByUserId(userId)
	if err != nil {
		log.Println(err)
	}
	return user, err
}
