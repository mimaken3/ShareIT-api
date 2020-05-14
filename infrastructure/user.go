package infrastructure

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewUserDB(db *gorm.DB) repository.UserRepository {
	return &userInfraStruct{db: db}
}

// ユーザを登録するのときのみ使用
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

// 全ユーザを取得(ページング)
func (userRepo *userInfraStruct) FindAllUsers(refPg int) (users []model.User, allPagingNum int, err error) {
	offset := (refPg - 1) * 10
	rows, err := userRepo.db.Raw(`
select 
  ddd.*,
  p.content as profile,
	i.icon_name
from 
  (
select 
  u.user_id, 
  u.user_name,
  u.email,
  u.password,
  group_concat(
    ut.topic_name  
    order by 
      ut.user_interested_topics_id
      separator '/'
  ) as interested_topics,
  u.created_date,
  u.updated_date,
  u.deleted_date
from 
  users as u
    , 
  (
    select 
      uit.user_interested_topics_id, 
      uit.user_id, 
      t.topic_name 
    from 
      user_interested_topics as uit 
      left join topics as t on (t.topic_id = uit.topic_id)
  ) as ut   
where 
  u.user_id = ut.user_id 
  and u.is_deleted = 0
group by 
  u.user_id
) as ddd
left join profiles as p on (ddd.user_id = p.user_id) 
left join icons as i on (ddd.user_id = i.user_id)
order by ddd.user_id
limit 10 offset ?
;
	`, offset).Rows()

	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		err = userRepo.db.ScanRows(rows, &user)
		if err == nil {
			users = append(users, user)
		}
	}

	// レコードがない場合
	if len(users) == 0 {
		return nil, 1, errors.New("record not found")
	}

	var count int
	userRepo.db.Table("users").Where("is_deleted = 0").Count(&count)
	allPagingNum = (count / 11) + 1

	return
}

// 全ユーザを取得(セレクトボックス)
func (userRepo *userInfraStruct) FindAllUsersForSelectBox() (users []model.User, err error) {
	if result := userRepo.db.Select("user_id, user_name").Where("is_deleted = 0").Find(&users); result.Error != nil {
		// レコードがない場合
		err = result.Error
		return
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

// ユーザを削除
func (userRepo *userInfraStruct) DeleteUser(userID uint) (err error) {
	deleteUser := model.User{}
	if result := userRepo.db.Find(&deleteUser, "user_id = ? AND is_deleted = ?", userID, 0); result.Error != nil {
		// レコードがない場合
		err = result.Error
		return
	}

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	// 削除状態に更新
	userRepo.db.Model(&deleteUser).
		Where("user_id= ? AND is_deleted = ?", userID, 0).
		Updates(map[string]interface{}{
			"deleted_date": currentDate,
			"is_deleted":   int8(1),
		})

	return nil
}

// ログイン
func (userRepo *userInfraStruct) Login(user model.User) (message string, resultUser model.User, err error) {
	result := userRepo.db.Raw(`
select 
  u.user_id, 
	i.icon_name,
  u.user_name, 
  u.email, 
  u.password, 
  p.content as profile, 
  dd.interested_topics, 
  u.created_date, 
  u.updated_date, 
  u.deleted_date 
from 
  users as u 
  left join profiles as p on (u.user_id = p.user_id) 
	left join icons as i on (u.user_id = i.user_id)
  inner join (
    select 
      td.user_id, 
      group_concat(
        td.topic_name 
        order by 
          td.user_interested_topics_id separator "/"
      ) as interested_topics 
    from 
      (
        select 
          uit.user_interested_topics_id, 
          uit.user_id, 
          t.topic_id, 
          t.topic_name 
        from 
          user_interested_topics as uit 
          inner join topics as t on (uit.topic_id = t.topic_id)
      ) as td 
    group by 
      td.user_id
  ) as dd on (dd.user_id = u.user_id) 
where 
  u.user_name = ? 
	AND u.is_deleted = 0
	;
	`, user.UserName).Scan(&resultUser)

	if result.Error != nil {
		// レコードがない場合
		return "failed", model.User{}, result.Error
	}

	err = bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(user.Password))

	if err != nil {
		// パスワードが一致しなかった場合
		user.Password = ""
		return "fail", user, err
	}

	// パスワードが一致した場合
	resultUser.Password = ""
	return "success", resultUser, nil
}

// ユーザを取得
func (userRepo *userInfraStruct) FindUserByUserId(userId int) (user model.User, err error) {
	result := userRepo.db.Raw(`
select 
  u.user_id, 
	i.icon_name,
  u.user_name, 
  u.email, 
  u.password, 
  p.content as profile, 
  dd.interested_topics, 
  u.created_date, 
  u.updated_date, 
  u.deleted_date 
from 
  users as u 
  left join profiles as p on (u.user_id = p.user_id) 
	left join icons as i on (u.user_id = i.user_id)
  inner join (
    select 
      td.user_id, 
      group_concat(
        td.topic_name 
        order by 
          td.user_interested_topics_id separator "/"
      ) as interested_topics 
    from 
      (
        select 
          uit.user_interested_topics_id, 
          uit.user_id, 
          t.topic_id, 
          t.topic_name 
        from 
          user_interested_topics as uit 
          inner join topics as t on (uit.topic_id = t.topic_id)
      ) as td 
    group by 
      td.user_id
  ) as dd on (dd.user_id = u.user_id) 
where 
  u.user_id = ?
  AND u.is_deleted = 0
;
`, userId).Scan(&user)

	if result.Error != nil {
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

	// 現在の日付とデフォの削除日を取得
	currentDate, defaultDeletedDate := getDate()

	signUpUser.UserID = lastUserId + 1
	signUpUser.UserName = user.UserName
	signUpUser.Email = user.Email
	signUpUser.Password = user.Password
	signUpUser.CreatedDate = currentDate
	signUpUser.UpdatedDate = currentDate
	signUpUser.DeletedDate = defaultDeletedDate

	userRepo.db.Create(&signUpUser)

	user.UserID = lastUserId + 1
	user.CreatedDate = currentDate
	user.UpdatedDate = currentDate
	user.DeletedDate = defaultDeletedDate

	// セキュリティのためパスワードは返さない
	user.Password = ""

	return user, err
}

// 興味トピックが更新されているか確認
func (userRepo *userInfraStruct) CheckUpdateInterestedTopic(willBeUpdatedUser model.User) (isUpdatedInterestedTopic bool, err error) {
	user := model.User{}

	result := userRepo.db.Raw(`
select 
  u.user_id, 
  u.user_name,
  u.email,
  u.password,
  group_concat(
    ut.topic_name  
    order by 
      ut.user_interested_topics_id
      separator '/'
  ) as interested_topics,
  u.created_date,
  u.updated_date,
  u.deleted_date
from 
  users as u, 
  (
    select 
      uit.user_interested_topics_id, 
      uit.user_id, 
      t.topic_name 
    from 
      user_interested_topics as uit 
      left join topics as t on (t.topic_id = uit.topic_id)
  ) as ut 
where 
  u.user_id = ut.user_id 
	and u.user_id = ?
  and u.is_deleted = 0
group by 
  u.user_id;
`, willBeUpdatedUser.UserID).Scan(&user)

	if result.Error != nil {
		// レコードがない場合
		err = result.Error
	}

	if willBeUpdatedUser.InterestedTopics == user.InterestedTopics {
		// 興味トピックが更新されていない場合
		return false, nil
	}
	// 興味トピックが更新されていた場合
	return true, nil
}

// パスワードをハッシュ化
func (userRepo *userInfraStruct) PasswordToHash(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPassword = string(hash)
	return
}

// 最後のユーザIDを取得
func (userRepo *userInfraStruct) FindLastUserId() (lastUserId uint, err error) {
	user := model.User{}
	// SELECT user_id FROM users ORDER BY user_id DESC LIMIT 1; と同義
	if result := userRepo.db.Select("user_id").Last(&user); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}
	lastUserId = user.UserID

	return
}

// ユーザのinterested_topicsにあるトピックを削除
func (userRepo *userInfraStruct) DeleteTopicFromInterestedTopics(deleteTopicID uint) (err error) {
	return
}

// 更新日を更新
func (userRepo *userInfraStruct) UpdateUser(userID uint) (err error) {
	user := model.User{}

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	userRepo.db.Model(&user).Where("user_id = ?", userID).Update("updated_date", currentDate)

	return
}
