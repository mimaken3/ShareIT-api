package service

// いいねON/OFF
func (l *likeServiceStruct) ToggleLikeByArticle(userID uint, articleID uint, isLiked bool) (err error) {
	if isLiked {
		// 最後のいいねIDを取得
		lastLikeID, _ := l.likeRepo.GetLastLikeID()

		// いいねを追加
		_ = l.likeRepo.AddLike(userID, articleID, lastLikeID)

	} else {
		// いいねを外す
		_ = l.likeRepo.DeleteLike(userID, articleID)
	}

	return
}
