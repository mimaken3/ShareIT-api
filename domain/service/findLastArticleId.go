package service

import "log"

// 最後の記事IDを取得
func (a *articleServiceStruct) FindLastArticleId() (lastArticleId uint, err error) {
	lastArticleId, err = a.articleRepo.FindLastArticleId()
	if err != nil {
		log.Println(err)
	}
	return lastArticleId, err
}
