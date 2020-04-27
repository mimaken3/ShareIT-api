package service

import (
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type iconServiceStruct struct {
	iconRepo repository.IconRepository
}

// Application層はこのInterfaceに依存
type IconServiceInterface interface {
	// アイコンを登録
	RegisterIcon(userID uint, formatName string) (registeredIconName string, err error)

	// アイコンを更新
	UpdateIcon(userID uint, formatName string) (updatedIconName string, err error)
}

// DIのための関数
func NewIconService(i repository.IconRepository) IconServiceInterface {
	return &iconServiceStruct{iconRepo: i}
}
