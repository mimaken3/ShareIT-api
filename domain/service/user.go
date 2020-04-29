package service

import (
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type userServiceStruct struct {
	userRepo repository.UserRepository
}

// Application層はこのInterfaceに依存
type UserServiceInterface interface {
	// 全ユーザを取得(ページング)
	FindAllUsersService(refPg int) (users []model.User, allPagingNum int, err error)

	// ユーザ登録のチェック
	CheckUserInfoService(checkUser model.User) (resultUserInfo model.CheckUserInfo, err error)

	// ユーザを取得
	FindUserByUserIdService(userId int) (user model.User, err error)

	// ユーザを登録
	SignUpUser(user model.User) (signedUpUser model.User, err error)

	// ユーザを削除
	DeleteUser(userID uint) (err error)

	// ログインチェック
	Login(user model.User) (message string, resultUser model.User, err error)

	// 興味トピックが更新されているか確認
	CheckUpdateInterestedTopic(willBeUpdatedUser model.User) (isUpdatedInterestedTopic bool, err error)

	// 最後のユーザIDを取得
	FindLastUserId() (lastUserId uint, err error)

	// ユーザのinterested_topicsにあるトピックを削除
	DeleteTopicFromInterestedTopics(deleteTopicID uint) (err error)

	// 更新日を更新
	UpdateUser(userID uint) (updatedUser model.User, err error)
}

// DIのための関数
func NewUserService(u repository.UserRepository) UserServiceInterface {
	return &userServiceStruct{userRepo: u}
}
