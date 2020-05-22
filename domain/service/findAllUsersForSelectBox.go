package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 全ユーザを取得(セレクトボックス)
func (u *userServiceStruct) FindAllUsersForSelectBox(userID uint) (users []model.User, err error) {
	_users, err := u.userRepo.FindAllUsersForSelectBox()

	if err != nil {
		return []model.User{}, err
	}

	users = make([]model.User, 0, len(_users)+1)
	allUser := model.User{UserID: 0, UserName: "全ユーザ"}

	// 0番目に「全ユーザ」を追加
	users = append(users, allUser)

	// 1番目に指定したユーザを追加
	deletedUsers := make([]model.User, 0, len(_users))
	for index, user := range _users {
		if user.UserID == userID {
			users = append(users, user)
			deletedUsers = UnsetUsers(_users, index)
		}
	}

	// 2番目移行に上で追加したもの以外の全てを追加
	for _, user := range deletedUsers {
		users = append(users, user)
	}

	return
}
