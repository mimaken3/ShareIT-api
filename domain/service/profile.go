package service

import (
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type profileServiceStruct struct {
	profileRepo repository.ProfileRepository
}

// Application層はこのInterfaceに依存
type ProfileServiceInterface interface {
	// 登録
	CreateProfileByUserID(content string, userID uint) (err error)

	// 更新
	UpdateProfileByUserID(content string, userID uint) (err error)

	// 削除
	DeleteProfileByUserID(userID uint) (err error)
}

// DIのための関数
func NewProfileService(p repository.ProfileRepository) ProfileServiceInterface {
	return &profileServiceStruct{profileRepo: p}
}
