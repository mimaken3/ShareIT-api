package infrastructure

import (
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
