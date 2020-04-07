package service

import (
	"log"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 指定したトピックを含む記事トピックを取得
func (a *articleServiceStruct) FindArticleIdsByTopicIdService(topicID uint) (articleIds []model.ArticleTopic, err error) {
	articleIds, err = a.articleRepo.FindArticleIdsByTopicId(topicID)
	if err != nil {
		log.Println(err)
	}
	return
}
