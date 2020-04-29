package service

import "fmt"

// アイコンを更新
func (i *iconServiceStruct) UpdateIcon(userID uint, formatName string) (updatedIconName string, err error) {
	var iconName string

	if formatName == "default.png" {
		iconName = formatName
	} else {
		iconName = fmt.Sprint(userID) + "." + formatName
	}

	uIconName, err := i.iconRepo.UpdateIcon(userID, iconName)

	// 署名付きURLを取得
	updatedIconName, err = GetPreSignedURL(uIconName)

	return
}
