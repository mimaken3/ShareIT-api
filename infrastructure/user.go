package infrastructure

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type userInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewUserDB(db *gorm.DB) repository.UserRepository {
	return &userInfraStruct{db: db}
}

// 全ユーザを取得
func (userRepo *userInfraStruct) FindAllUsers() (users []model.User, err error) {
	userRepo.db.Find(&users)
	return
}

// ユーザを取得
func (userRepo *userInfraStruct) FindUserByUserId(userId int) (user model.User, err error) {
	userRepo.db.Find(&user, "user_id=?", userId)
	return
}

// ユーザを登録
func (userRepo *userInfraStruct) SignUpUser(user model.User, lastUserId uint) (model.User, error) {
	// TODO: パスワードハッシュ化、もしくはDB通信エラーで使用予定
	var err error

	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	nowTime := time.Now().Format(dateFormat)
	customisedNowTime, _ := time.Parse(dateFormat, nowTime)

	const defaultDeletedDateStr = "9999-12-31 23:59:59"
	defaultDeletedDate, _ := time.Parse(dateFormat, defaultDeletedDateStr)

	user.UserID = lastUserId + 1
	user.CreatedDate = customisedNowTime
	user.UpdatedDate = customisedNowTime
	user.DeletedDate = defaultDeletedDate

	userRepo.db.Create(&user)

	// セキュリティのためパスワードを返さない
	user.Password = ""

	return user, err
}

// 最後のユーザIDを取得
func (userRepo *userInfraStruct) FindLastUserId() (lastUserId uint, err error) {
	user := model.User{}
	// SELECT user_id FROM users ORDER BY user_id DESC LIMIT 1; と同義
	userRepo.db.Select("user_id").Last(&user)
	lastUserId = user.UserID
	return
}
