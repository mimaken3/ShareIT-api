package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ユーザを登録
func (u *userServiceStruct) SignUpUser(user model.User) (signUpedUser model.User, err error) {
	signUpedUser, err = u.userRepo.SignUpUser(user)
	if err != nil {
		log.Println(err)
	}
	// セキュリティのためパスワードを返さない
	signUpedUser.Password = ""

	return
}
