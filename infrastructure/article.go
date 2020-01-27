package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type articleInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewArticleDB(db *gorm.DB) repository.ArticleRepository {
	return &articleInfraStruct{db: db}
}

// 全記事を取得
func (articleRepo *articleInfraStruct) FindAllArticles() (articles []model.Article, err error) {
	articleRepo.db.Find(&articles)
	return
}
