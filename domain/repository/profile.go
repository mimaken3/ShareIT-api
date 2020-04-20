package repository

// ProfileRepository is interface for infrastructure
type ProfileRepository interface {
	// 最後のIDを取得
	FindLastProfileID() (lastID uint, err error)

	// 登録
	CreateProfileByUserID(lastID uint, content string, userID uint) (err error)
}
