package service

// 登録
func (p *profileServiceStruct) CreateProfileByUserID(content string, userID uint) (err error) {
	// 最後のIDを取得
	lastID, err := p.profileRepo.FindLastProfileID()

	if err != nil {
		return err
	}

	err = p.profileRepo.CreateProfileByUserID(lastID, content, userID)

	return
}
