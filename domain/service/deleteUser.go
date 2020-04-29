package service

// ユーザを削除
func (u *userServiceStruct) DeleteUser(userID uint) (err error) {
	err = u.userRepo.DeleteUser(userID)
	return
}
