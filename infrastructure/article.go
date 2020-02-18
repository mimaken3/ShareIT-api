package infrastructure

import (
	"time"

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

// 記事を取得
func (articleRepo *articleInfraStruct) FindArticleByArticleId(articleId uint) (article model.Article, err error) {
	// SELECT * FROM article WHERE article_id = :articleId;
	articleRepo.db.Find(&article, "article_id = ?", articleId)
	return
}

// 記事を投稿
func (articleRepo *articleInfraStruct) CreateArticle(createArticle model.Article, lastArticleId uint) (createdArticle model.Article, err error) {
	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	nowTime := time.Now().Format(dateFormat)
	customisedNowTime, _ := time.Parse(dateFormat, nowTime)

	const defaultDeletedDateStr = "9999-12-31 23:59:59"
	defaultDeletedDate, _ := time.Parse(dateFormat, defaultDeletedDateStr)

	// DBに保存する記事のモデルを作成
	createdArticle.ArticleID = lastArticleId + 1
	createdArticle.ArticleTitle = createArticle.ArticleTitle
	createdArticle.CreatedUserID = createArticle.CreatedUserID
	createdArticle.ArticleContent = createArticle.ArticleContent
	createdArticle.ArticleTopics = createArticle.ArticleTopics
	createdArticle.CreatedDate = customisedNowTime
	createdArticle.UpdatedDate = customisedNowTime
	createdArticle.DeletedDate = defaultDeletedDate

	articleRepo.db.Create(&createdArticle)

	return
}

// 特定のユーザの全記事を取得
func (articleRepo *articleInfraStruct) FindArticlesByUserId(userID uint) (articles []model.Article, err error) {
	articleRepo.db.Where("created_user_id = ?", userID).Find(&articles)
	return
}

// 特定のトピックを含む記事を取得
func (articleRepo *articleInfraStruct) FindArticlesByTopicId(articleIds []model.ArticleTopic) (articles []model.Article, err error) {
	for i := 0; i < len(articleIds); i++ {
		// TODO: 要修正 毎回articleを作ってる
		var article = model.Article{}
		articleRepo.db.Where("article_id = ?", articleIds[i].ArticleID).Find(&article)
		articles = append(articles, article)
	}
	return
}

// 指定したトピックを含む記事のIDを取得
func (articleRepo *articleInfraStruct) FindArticleIdsByTopicId(topicID uint) (articleIds []model.ArticleTopic, err error) {
	// SELECT * FROM article_topic WHERE topic_id = :topicID;
	articleRepo.db.Where("topic_id = ?", topicID).Find(&articleIds)
	return articleIds, err
}

// 最後の記事IDを取得
func (articleRepo *articleInfraStruct) FindLastArticleId() (lastArticleId uint, err error) {
	article := model.Article{}
	// SELECT article_id FROM articles ORDER BY article_id DESC LIMIT 1;
	articleRepo.db.Select("article_id").Last(&article)
	lastArticleId = article.ArticleID
	return
}
