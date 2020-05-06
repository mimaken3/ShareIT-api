package service

// ユーザが付けたいいねを全削除
func (l *likeServiceStruct) DeleteLikeByUserID(userID uint) (err error) {
	err = l.likeRepo.DeleteLikeByUserID(userID)
	return
}
