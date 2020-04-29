package service

import "fmt"

// アイコンを登録
func (i *iconServiceStruct) RegisterIcon(userID uint, formatName string) (registeredIconName string, err error) {
	var iconName string

	if formatName == "default.png" {
		iconName = formatName
	} else {
		iconName = fmt.Sprint(userID) + "." + formatName
	}

	// 最後のアイコンIDを取得
	lastIconID, err := i.iconRepo.FindLastIconID()

	registeredIconName, err = i.iconRepo.RegisterIcon(userID, iconName, lastIconID)
	return
}
