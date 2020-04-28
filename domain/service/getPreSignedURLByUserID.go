package service

// ユーザIDから署名付きURLを取得
func (i *iconServiceStruct) GetPreSignedURLByUserID(userID uint) (preSignedURL string, err error) {
	// ユーザIDからアイコン名を取得
	iconName, err := i.iconRepo.FindIconNameByUserID(userID)

	// 署名付きURLを取得
	preSignedURL, err = GetPreSignedURL(iconName)

	return
}
