package service

import "github.com/mimaken3/ShareIT-api/domain/model"

func (ui *userInterestedTopicServiceStruct) UpdateUserInterestedTopic(willBeUpdatedUser model.User) (err error) {
	// ユーザIDに紐づくトピックを削除
	err = ui.userInterestedTopicRepo.DeleteUserInterestedTopic(willBeUpdatedUser)
	if err != nil {
		return err
	}

	// 最後のIDを取得
	lastUserInterestedTopicID, err := ui.userInterestedTopicRepo.GetLastID()
	if err != nil {
		return err
	}

	// 追加
	err = ui.userInterestedTopicRepo.CreateUserTopic(willBeUpdatedUser.InterestedTopics, lastUserInterestedTopicID, willBeUpdatedUser.UserID)

	return
}
