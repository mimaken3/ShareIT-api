package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type iconInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewIconDB(db *gorm.DB) repository.IconRepository {
	return &iconInfraStruct{db: db}
}

// 最後のアイコンIDを取得
func (iconRepo *iconInfraStruct) FindLastIconID() (lastIconID uint, err error) {
	icon := model.Icon{}
	if result := iconRepo.db.Select("icon_id").Last(&icon); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}
	lastIconID = icon.IconID

	return
}

// アイコンを登録
func (iconRepo *iconInfraStruct) RegisterIcon(userID uint, iconName string, lastIconID uint) (registeredIconName string, err error) {
	icon := model.Icon{}

	icon.IconID = lastIconID + 1
	icon.UserID = userID
	icon.IconName = iconName

	iconRepo.db.Create(&icon)

	return icon.IconName, err
}

// アイコンを更新
func (iconRepo *iconInfraStruct) UpdateIcon(userID uint, formatName string) (updatedIconName string, err error) {
	updateIcon := model.Icon{}

	iconRepo.db.Model(&updateIcon).Where("user_id = ?", userID).Update("icon_name", formatName)

	return updateIcon.IconName, nil
}
