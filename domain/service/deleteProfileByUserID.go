package service

// 削除
func (p *profileServiceStruct) DeleteProfileByUserID(userID uint) (err error) {
	err = p.profileRepo.DeleteProfileByUserID(userID)
	return
}
