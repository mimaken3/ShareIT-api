package infrastructure

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type articleTopicInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewArticleTopicDB(db *gorm.DB) repository.ArticleTopicRepository {
	return &articleTopicInfraStruct{db: db}
}

// 記事に紐づく記事トピックを追加
func (articleTopicRepo *articleTopicInfraStruct) CreateArticleTopic(article model.Article, lastArticleTopicId uint) {
	articleTopic := model.ArticleTopic{}
	articleID := article.ArticleID
	topicsStr := article.ArticleTopics

	// トピック名の配列
	topicsArr := strings.Split(topicsStr, ",")

	// トピックIDの配列
	topicIDsArr := make([]uint, len(topicsArr))

	for i, topicName := range topicsArr {
		var topicID uint
		articleTopicRepo.db.Table("topics").Where("topic_name = ?", topicName).Select("topic_id").Row().Scan(&topicID)
		topicIDsArr[i] = topicID
	}

	// 記事トピックID
	insertArticleTopicId := lastArticleTopicId

	for _, topicID := range topicIDsArr {
		insertArticleTopicId = insertArticleTopicId + 1
		if topicID != 0 {
			// INSERT INTO article_topics VALUES(:lastArticleTopicId + 1, :articleID, :topicID);
			articleTopic.ArticleTopicID = insertArticleTopicId
			articleTopic.ArticleID = articleID
			articleTopic.TopicID = uint(topicID)
			articleTopicRepo.db.Create(&articleTopic)
		}
	}
}

// 最後の記事トピックIDを取得
func (articleTopicRepo *articleTopicInfraStruct) FindLastArticleTopicId() (lastArticleTopicId uint, err error) {
	articleTopic := model.ArticleTopic{}
	// SELECT article_topic_id FROM article_topics ORDER BY article_topic_id DESC LIMIT 1;
	articleTopicRepo.db.Select("article_topic_id").Last(&articleTopic)
	lastArticleTopicId = articleTopic.ArticleTopicID
	return
}

// 記事に紐づく記事トピックを削除
func (articleTopicRepo *articleTopicInfraStruct) DeleteArticleTopic(willBeDeletedArticle model.Article) {
	uintArticleID := willBeDeletedArticle.ArticleID

	// 物理削除
	articleTopicRepo.db.Where("article_id = ?", uintArticleID).Delete(&model.ArticleTopic{})
}

// トピックに紐づく記事トピックを削除
func (articleTopicRepo *articleTopicInfraStruct) DeleteArticleTopicByTopicID(topicID uint) (err error) {
	// 記事トピックが1つしかない場合、それを「その他(1)」に更新
	var articleTopicIDs []uint
	rows, err := articleTopicRepo.db.Raw("select t2.article_topic_id from(select article_id, c from ("+
		"select article_id, count(*) as c from article_topics group by article_id) as t where t.c = 1) as t1 inner join "+
		"(select * from article_topics where topic_id = ?) as t2 on t1.article_id = t2.article_id", topicID).Rows()
	defer rows.Close()
	for rows.Next() {
		var articleTopicID uint
		err = rows.Scan(&articleTopicID)
		if err == nil {
			articleTopicIDs = append(articleTopicIDs, articleTopicID)
		}
	}

	if len(articleTopicIDs) != 0 {
		// 「その他」に更新
		// UPDATE article_topics SET topic_id = 0 WHERE article_topic_id IN (?);
		articleTopicRepo.db.Table("article_topics").Where("article_topic_id IN (?)", articleTopicIDs).Update("topic_id", 1)
	}

	// 物理削除
	articleTopicRepo.db.Where("topic_id = ?", topicID).Delete(&model.ArticleTopic{})

	return
}
