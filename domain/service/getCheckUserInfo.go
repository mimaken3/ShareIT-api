package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ユーザ登録のチェック
func (u *userServiceStruct) CheckUserInfoService(checkUser model.User) (resultUserInfo model.CheckUserInfo, err error) {
	resultUserInfo, err = u.userRepo.CheckUserInfo(checkUser)
	if err != nil {
		log.Println(err)
	}
	return
}
