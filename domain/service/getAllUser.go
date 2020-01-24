package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 全ユーザを取得
func (u *userServiceStruct) FindAllUsersService() ([]model.User, error) {
	users, err := u.userRepo.FindAllUsers()
	if err != nil {
		log.Println(err)
	}
	return users, err
}
