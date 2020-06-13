package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事をいいねした全ユーザ取得
func (u *userServiceStruct) FindAllLikedUsersByArticleID(articleID uint) (users []model.User, err error) {
	users, err = u.userRepo.FindAllLikedUsersByArticleID(articleID)
	if err != nil {
		return users, err
	}

	for i := 0; i < len(users); i++ {
		users[i].Email = ""

		// セキュリティのためパスワードを返さない
		users[i].Password = ""
	}

	return
}
