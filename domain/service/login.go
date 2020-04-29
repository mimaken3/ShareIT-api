package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// ログインチェック
func (u *userServiceStruct) Login(user model.User) (message string, resultUser model.User, err error) {
	message, resultUser, err = u.userRepo.Login(user)

	if resultUser.IconName != "" {
		headerIconURL, err := GetHeaderIconURL(resultUser.IconName)
		if err != nil {
			return "Failed to create icon URL for header", model.User{}, err
		}
		resultUser.IconName = headerIconURL
	}
	return
}
