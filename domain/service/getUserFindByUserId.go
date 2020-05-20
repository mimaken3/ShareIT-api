package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ユーザを取得
func (u *userServiceStruct) FindUserByUserIdService(userId int) (model.User, error) {
	user, err := u.userRepo.FindUserByUserId(userId)

	if err != nil {
		return user, err
	}

	// 署名付きURLを取得
	user.IconName, err = GetPreSignedURL(user.IconName)

	user.Email = ""

	// セキュリティのためパスワードを返さない
	user.Password = ""

	return user, err
}
