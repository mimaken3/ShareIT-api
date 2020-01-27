package service

import "log"

// 最後のユーザIDを取得
func (u *userServiceStruct) FindLastUserId() (lastUserId uint, err error) {
	lastUserId, err = u.userRepo.FindLastUserId()
	if err != nil {
		log.Println(err)
	}
	return
}
