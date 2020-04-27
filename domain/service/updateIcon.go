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

	updatedIconName, err = i.iconRepo.UpdateIcon(userID, iconName)

	return
}
