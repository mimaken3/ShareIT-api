package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// ログインチェック
func (u *userServiceStruct) Login(user model.User) (message string, resultUser model.User, err error) {
	message, resultUser, err = u.userRepo.Login(user)
	return
}
