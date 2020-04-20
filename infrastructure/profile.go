package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type profileInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewProfileDB(db *gorm.DB) repository.ProfileRepository {
	return &profileInfraStruct{db: db}
}

// 最後のIDを取得
func (pRepo *profileInfraStruct) FindLastProfileID() (lastID uint, err error) {
	p := model.Profile{}
	if result := pRepo.db.Select("profile_id").Last(&p); result.Error != nil {
		return 0, nil
	}

	lastID = p.ProfileID

	return
}

// 登録
func (pRepo *profileInfraStruct) CreateProfileByUserID(lastID uint, content string, userID uint) (err error) {
	p := model.Profile{}
	p.ProfileID = lastID + 1
	p.Content = content
	p.UserID = userID

	pRepo.db.Create(&p)

	return
}

// 更新
func (pRepo *profileInfraStruct) UpdateProfileByUserID(content string, userID uint) (err error) {
	p := model.Profile{}
	pRepo.db.Model(&p).Where("user_id = ?", userID).Update("content", content)

	return
}

// 削除
// DeleteProfileByUSerID(userID uint)(err error)
