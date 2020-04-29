package service

// 更新
func (p *profileServiceStruct) UpdateProfileByUserID(content string, userID uint) (err error) {

	err = p.profileRepo.UpdateProfileByUserID(content, userID)

	return
}
