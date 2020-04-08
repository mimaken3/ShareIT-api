package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ユーザを登録
func (u *userServiceStruct) SignUpUser(user model.User) (signUpedUser model.User, err error) {
	// 最後のユーザIDを取得
	lastUserId, err := u.userRepo.FindLastUserId()

	// パスワードをハッシュ化
	hashedPassword, err := u.userRepo.PasswordToHash(user.Password)

	// ハッシュ化失敗時
	if err != nil {
		return model.User{}, err
	}

	// ハッシュ化したパスワードをセット
	user.Password = hashedPassword

	// ユーザ登録
	signUpedUser, err = u.userRepo.SignUpUser(user, lastUserId)

	// ユーザ登録失敗時
	if err != nil {
		return model.User{}, err
	}
	// セキュリティのためパスワードを返さない
	signUpedUser.Password = ""

	return
}
