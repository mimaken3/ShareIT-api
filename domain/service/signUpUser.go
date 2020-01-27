package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ユーザを登録
func (u *userServiceStruct) SignUpUser(user model.User) (signUpedUser model.User, err error) {
	// 最後のユーザIDを取得
	lastUserId, err := u.userRepo.FindLastUserId()

	signUpedUser, err = u.userRepo.SignUpUser(user, lastUserId)

	if err != nil {
		log.Println(err)
	}
	// セキュリティのためパスワードを返さない
	signUpedUser.Password = ""

	return
}
