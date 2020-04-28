package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 更新日を更新
func (u *userServiceStruct) UpdateUser(userID uint) (updatedUser model.User, err error) {
	err = u.userRepo.UpdateUser(userID)
	if err != nil {
		return model.User{}, err
	}
	updatedUser, err = u.userRepo.FindUserByUserId(int(userID))

	// パスワードを隠す
	updatedUser.Password = ""

	return
}
