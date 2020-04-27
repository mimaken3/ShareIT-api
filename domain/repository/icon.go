package repository

// IconRepository is interface for infrastructure
type IconRepository interface {
	// 最後のアイコンIDを取得
	FindLastIconID() (lastIconID uint, err error)

	// アイコンを登録
	RegisterIcon(userID uint, formatName string, lastIconID uint) (registeredIconName string, err error)

	// アイコンを更新
	// UpdateIcon(userID uint, formatName string) (updatedIconName string, err error)
}
