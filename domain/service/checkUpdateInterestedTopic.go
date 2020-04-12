package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// 記事のトピックが更新されているか確認
func (u *userServiceStruct) CheckUpdateInterestedTopic(willBeUpdatedUser model.User) (isUpdatedInterestedTopic bool, err error) {
	isUpdatedInterestedTopic, err = u.userRepo.CheckUpdateInterestedTopic(willBeUpdatedUser)

	return
}
