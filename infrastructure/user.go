package infrastructure

import (
	"errors"
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

type SignUpUser struct {
	UserID      uint      `gorm:"primary_key" json:"user_id"`
	UserName    string    `gorm:"size:255" json:"user_name"`
	Email       string    `gorm:"size:255" json:"email"`
	Password    string    `gorm:"size:255" json:"password"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DeletedDate time.Time `json:"deleted_date"`
	IsDeleted   int8      `json:"-"`
}

func (SignUpUser) TableName() string {
	return "users"
}

// 全ユーザを取得
func (userRepo *userInfraStruct) FindAllUsers() (users []model.User, err error) {
	userRepo.db.Find(&users, "is_deleted = ?", 0)

	// レコードがない場合
	if len(users) == 0 {
		return nil, errors.New("record not found")
	}

	return
}

// ユーザ登録のチェック
func (userRepo *userInfraStruct) CheckUserInfo(checkUser model.User) (resultUserInfo model.CheckUserInfo, err error) {

	// ユーザ名の重複チェック
	if userRepo.db.Where("user_name = ? AND is_deleted = ?", checkUser.UserName, 0).First(&model.User{}).RecordNotFound() {
		resultUserInfo.ResultUserNameNum = 0
		resultUserInfo.ResultUserNameText = "このユーザ名は登録出来ます！"
	} else {
		resultUserInfo.ResultUserNameNum = 1
		resultUserInfo.ResultUserNameText = "このユーザ名は既に使われています..."
	}

	// メアドの重複チェック
	if userRepo.db.Where("email = ? AND is_deleted = ?", checkUser.Email, 0).First(&model.User{}).RecordNotFound() {
		resultUserInfo.ResultEmailNum = 0
		resultUserInfo.ResultEmailText = "このメールアドレスは登録出来ます！"
	} else {
		resultUserInfo.ResultEmailNum = 1
		resultUserInfo.ResultEmailText = "このメールアドレスは既に使われています..."
	}

	resultUserInfo.UserName = checkUser.UserName
	resultUserInfo.Email = checkUser.Email

	return
}

// ユーザを取得
func (userRepo *userInfraStruct) FindUserByUserId(userId int) (user model.User, err error) {
	if result := userRepo.db.Find(&user, "user_id = ? AND is_deleted = ?", userId, 0); result.Error != nil {
		// レコードがない場合
		err = result.Error
	}
	return
}

// ユーザを登録
func (userRepo *userInfraStruct) SignUpUser(user model.User, lastUserId uint) (model.User, error) {
	// TODO: パスワードハッシュ化、もしくはDB通信エラーで使用予定
	var err error

	signUpUser := SignUpUser{}

	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	nowTime := time.Now().Format(dateFormat)
	customisedNowTime, _ := time.Parse(dateFormat, nowTime)

	const defaultDeletedDateStr = "9999-12-31 23:59:59"
	defaultDeletedDate, _ := time.Parse(dateFormat, defaultDeletedDateStr)

	signUpUser.UserID = lastUserId + 1
	signUpUser.UserName = user.UserName
	signUpUser.Email = user.Email
	signUpUser.Password = user.Password
	signUpUser.CreatedDate = customisedNowTime
	signUpUser.UpdatedDate = customisedNowTime
	signUpUser.DeletedDate = defaultDeletedDate

	userRepo.db.Create(&signUpUser)

	user.UserID = lastUserId + 1
	user.CreatedDate = customisedNowTime
	user.UpdatedDate = customisedNowTime
	user.DeletedDate = defaultDeletedDate

	// セキュリティのためパスワードは返さない
	user.Password = ""

	return user, err
}

// 最後のユーザIDを取得
func (userRepo *userInfraStruct) FindLastUserId() (lastUserId uint, err error) {
	user := model.User{}
	// SELECT user_id FROM users ORDER BY user_id DESC LIMIT 1; と同義
	if result := userRepo.db.Select("user_id").Where("is_deleted = ?", 0).Last(&user); result.Error != nil {
		// レコードがない場合
		err = result.Error
	}
	lastUserId = user.UserID

	return
}

// ユーザのinterested_topicsにあるトピックを削除
func (userRepo *userInfraStruct) DeleteTopicFromInterestedTopics(deleteTopicID uint) (err error) {
	// var topicIDArrFromAT []int
	// var createdUserIDArr []int
	//
	// rows, err := userRepo.db.Raw("SELECT article_id FROM article_topics WHERE topic_id = ?", deleteTopicID).Rows()
	//
	// defer rows.Close()
	// for rows.Next() {
	// 	article_id := 0
	// 	err = rows.Scan(&article_id)
	// 	if err == nil {
	// 		topicIDArrFromAT = append(topicIDArrFromAT, article_id)
	// 	}
	// }
	//
	// // レコードがない場合
	// if len(topicIDArrFromAT) == 0 {
	// 	return err
	// }
	//
	// userIDRows, err := userRepo.db.Raw("SELECT created_user_id FROM articles WHERE article_id IN (?)", topicIDArrFromAT).Rows()
	// defer rows.Close()
	// for userIDRows.Next() {
	// 	userID := 0
	// 	err = userIDRows.Scan(&userID)
	// 	if err == nil {
	// 		createdUserIDArr = append(createdUserIDArr, userID)
	// 	}
	// }
	//
	// fmt.Println(createdUserIDArr)

	return
}
