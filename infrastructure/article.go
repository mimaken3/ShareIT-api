package infrastructure

import (
	"errors"
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

// 登録or更新するのときのみ使用
type CreateArticle struct {
	ArticleID      uint      `gorm:"primary_key" json:"article_id"`
	ArticleTitle   string    `gorm:"size:255" json:"article_title"`
	ArticleContent string    `gorm:"size:1000" json:"article_content"`
	CreatedUserID  uint      `json:"created_user_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
	DeletedDate    time.Time `json:"deleted_date"`
	IsDeleted      int8      `json:"-"`
}

func (CreateArticle) TableName() string {
	return "articles"
}

// 全記事を取得(ページング)
func (articleRepo *articleInfraStruct) FindAllArticles(refPg int) (articles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10
	rows, err :=
		articleRepo.db.Raw(
			`
select * from(
select 
  a.article_id, 
  a.article_title, 
  a.article_content, 
  group_concat(
    att.topic_name 
    order by 
      att.article_topic_id
			 separator '/'
  ) as article_topics, 
  a.created_user_id, 
  a.created_date, 
  a.updated_date, 
  a.deleted_date 
from 
  articles as a, 
  (
  select 
  at.article_topic_id,
  at.article_id,
  t.topic_name
from 
  article_topics as at 
  left join topics as t on at.topic_id = t.topic_id
  ) as att 
where 
  a.article_id = att.article_id 
  and is_deleted = 0 
group by 
  a.article_id
  ) as tt
  limit 10 offset ? 
  ;
`, offset).Rows()

	defer rows.Close()
	for rows.Next() {
		article := model.Article{}
		err = articleRepo.db.ScanRows(rows, &article)
		if err == nil {
			articles = append(articles, article)
		}
	}

	// レコードがない場合
	if len(articles) == 0 {
		return nil, 1, errors.New("record not found")
	}

	var count int
	articleRepo.db.Table("articles").Where("is_deleted = 0").Count(&count)
	allPagingNum = (count / 11) + 1

	return
}

// 記事を検索(ページング)
func (articleRepo *articleInfraStruct) SearchAllArticles(refPg int, userID uint, topicIDs []uint) (searchedArticles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10

	if userID == 0 && topicIDs[0] != 0 {
		// 「全ユーザ」かつ「特定トピック」の場合
		rows, err :=
			articleRepo.db.Raw(`
select 
  * 
from 
  (
    select 
      a.article_id, 
      a.article_title, 
      a.article_content, 
      group_concat(
        att.topic_name 
        order by 
          att.article_topic_id separator '/'
      ) as article_topics, 
      a.created_user_id, 
      a.created_date, 
      a.updated_date, 
      a.deleted_date 
    from 
      (
        select 
          * 
        from 
          articles 
        where 
          article_id in (
            select 
              article_id 
            from 
              article_topics 
            where 
              topic_id in (?)
          )
      ) as a, 
      (
        select 
          at.article_topic_id, 
          at.article_id, 
          t.topic_name 
        from 
          article_topics as at 
          left join topics as t on at.topic_id = t.topic_id
      ) as att 
    where 
      a.article_id = att.article_id 
      and is_deleted = 0 
    group by 
      a.article_id
  ) as tt 
limit 
  10 offset ?
;	
		`, topicIDs, offset).Rows()
		defer rows.Close()
		for rows.Next() {
			article := model.Article{}
			err = articleRepo.db.ScanRows(rows, &article)
			if err == nil {
				searchedArticles = append(searchedArticles, article)
			}
		}

		// レコードがない場合
		if len(searchedArticles) == 0 {
			return nil, 1, errors.New("record not found")
		}

		var count int
		row := articleRepo.db.Raw(`
select 
  count(*) 
from 
  articles 
where 
  is_deleted = 0 
  and article_id in (
    select 
      article_id 
    from 
      article_topics 
    where 
      topic_id in (?)
  );
	`, topicIDs).Row()
		row.Scan(&count)
		allPagingNum = (count / 11) + 1

		return searchedArticles, allPagingNum, nil
	}
	// 「特定のユーザ」かつ「特定のトピック」の場合
	rows, err :=
		articleRepo.db.Raw(`
select 
  * 
from 
  (
    select 
      a.article_id, 
      a.article_title, 
      a.article_content, 
      group_concat(
        att.topic_name 
        order by 
          att.article_topic_id separator '/'
      ) as article_topics, 
      a.created_user_id, 
      a.created_date, 
      a.updated_date, 
      a.deleted_date 
    from 
      (
        select 
          * 
        from 
          articles 
        where 
          article_id in (
            select 
              article_id 
            from 
              article_topics 
            where 
              topic_id in (?)
							and created_user_id = ?
          )
      ) as a, 
      (
        select 
          at.article_topic_id, 
          at.article_id, 
          t.topic_name 
        from 
          article_topics as at 
          left join topics as t on at.topic_id = t.topic_id
      ) as att 
    where 
      a.article_id = att.article_id 
      and is_deleted = 0 
    group by 
      a.article_id
  ) as tt 
limit 
  10 offset ?
;	
		`, topicIDs, userID, offset).Rows()
	defer rows.Close()
	for rows.Next() {
		article := model.Article{}
		err = articleRepo.db.ScanRows(rows, &article)
		if err == nil {
			searchedArticles = append(searchedArticles, article)
		}
	}

	// レコードがない場合
	if len(searchedArticles) == 0 {
		return nil, 1, errors.New("record not found")
	}

	var count int
	row := articleRepo.db.Raw(`
select 
  count(*) 
from 
  articles 
where 
  is_deleted = 0 
	and created_user_id = ?
  and article_id in (
    select 
      article_id 
    from 
      article_topics 
    where 
      topic_id in (?)
  );
	`, userID, topicIDs).Row()
	row.Scan(&count)
	allPagingNum = (count / 11) + 1

	return searchedArticles, allPagingNum, nil
}

// 記事を取得
func (articleRepo *articleInfraStruct) FindArticleByArticleId(articleId uint) (article model.Article, err error) {
	result := articleRepo.db.Raw(`
select 
  a.article_id, 
  a.article_title, 
  a.article_content, 
  group_concat(
    att.topic_name 
    order by 
      att.article_topic_id
			 separator '/'
  ) as article_topics, 
  a.created_user_id, 
  a.created_date, 
  a.updated_date, 
  a.deleted_date 
from 
  articles as a, 
  (
    select 
      at.article_topic_id, 
      at.article_id, 
      at.topic_id, 
      t.topic_name 
    from 
      article_topics as at 
      left join topics as t on at.topic_id = t.topic_id
  ) as att 
where 
  a.article_id = att.article_id 
  and a.article_id = ? 
  and is_deleted = 0 
group by 
  a.article_id;
	`, articleId).Scan(&article)

	if result.Error != nil {
		// レコードがない場合
		err = result.Error
	}
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

	ar := CreateArticle{}

	// DBに保存する記事のモデルを作成
	ar.ArticleID = lastArticleId + 1
	ar.ArticleTitle = createArticle.ArticleTitle
	ar.ArticleContent = createArticle.ArticleContent
	ar.CreatedUserID = createArticle.CreatedUserID
	ar.CreatedDate = customisedNowTime
	ar.UpdatedDate = customisedNowTime
	ar.DeletedDate = defaultDeletedDate

	articleRepo.db.Create(&ar)

	// DBに保存した記事を返す
	createdArticle.ArticleID = lastArticleId + 1
	createdArticle.ArticleTitle = createArticle.ArticleTitle
	createdArticle.ArticleContent = createArticle.ArticleContent
	createdArticle.ArticleTopics = createArticle.ArticleTopics
	createdArticle.CreatedUserID = createArticle.CreatedUserID
	createdArticle.CreatedDate = customisedNowTime
	createdArticle.UpdatedDate = customisedNowTime
	createdArticle.DeletedDate = defaultDeletedDate

	return
}

// 記事を更新
func (articleRepo *articleInfraStruct) UpdateArticleByArticleId(willBeUpdatedArticle model.Article) (updatedArticle model.Article, err error) {
	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	updateTime := time.Now().Format(dateFormat)
	customisedUpdateTime, _ := time.Parse(dateFormat, updateTime)

	updateId := willBeUpdatedArticle.ArticleID

	// 更新するフィールドを設定
	updateTitle := willBeUpdatedArticle.ArticleTitle
	updateContent := willBeUpdatedArticle.ArticleContent

	// 更新
	articleRepo.db.Model(&updatedArticle).
		Where("article_id = ?", updateId).
		Updates(map[string]interface{}{
			"article_title":   updateTitle,
			"article_content": updateContent,
			"updated_date":    customisedUpdateTime,
		})

	// 興味トピックを文字列で,区切りで取得
	var articleTopicsStr string
	err = articleRepo.db.Raw(`
select 
  group_concat(
    att.topic_name 
    order by 
      att.article_topic_id
			 separator '/'
  ) as article_topics
from 
  articles as a, 
  (
    select 
      at.article_topic_id, 
      at.article_id, 
      at.topic_id, 
      t.topic_name 
    from 
      article_topics as at 
      left join topics as t on at.topic_id = t.topic_id
  ) as att 
where 
  a.article_id = att.article_id 
  and a.article_id = ? 
  and is_deleted = 0 
group by 
  a.article_id;
			`, willBeUpdatedArticle.ArticleID).Row().Scan(&articleTopicsStr)

	if err != nil {
		return model.Article{}, err
	}

	updatedArticle.ArticleTopics = articleTopicsStr

	// updateで値の入ってないフィールドに値を格納
	updatedArticle.ArticleID = willBeUpdatedArticle.ArticleID
	updatedArticle.CreatedUserID = willBeUpdatedArticle.CreatedUserID
	updatedArticle.CreatedDate = willBeUpdatedArticle.CreatedDate
	updatedArticle.DeletedDate = willBeUpdatedArticle.DeletedDate

	return
}

// 特定のユーザの全記事を取得
func (articleRepo *articleInfraStruct) FindArticlesByUserId(userID uint, refPg int) (articles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10
	rows, err :=
		articleRepo.db.Raw(`
select 
  * 
from 
  (
    select 
      a.article_id, 
      a.article_title, 
      a.article_content, 
      group_concat(
        att.topic_name 
        order by 
          att.article_topic_id separator '/'
      ) as article_topics, 
      a.created_user_id, 
      a.created_date, 
      a.updated_date, 
      a.deleted_date 
    from 
      (
        select 
          * 
        from 
          articles 
        where 
          created_user_id = ?
      ) as a, 
      (
        select 
          at.article_topic_id, 
          at.article_id, 
          t.topic_name 
        from 
          article_topics as at 
          left join topics as t on at.topic_id = t.topic_id
      ) as att 
    where 
      a.article_id = att.article_id 
      and is_deleted = 0 
    group by 
      a.article_id
  ) as tt 
limit 
  10 offset ?;
`, userID, offset).Rows()

	defer rows.Close()
	for rows.Next() {
		article := model.Article{}
		err = articleRepo.db.ScanRows(rows, &article)
		if err == nil {
			articles = append(articles, article)
		}
	}

	// レコードがない場合
	if len(articles) == 0 {
		return nil, 1, errors.New("record not found")
	}

	var count int
	articleRepo.db.Table("articles").Where("is_deleted = 0 AND created_user_id = ?", userID).Count(&count)
	allPagingNum = (count / 11) + 1

	return
}

// 特定のトピックを含む全記事を取得
func (articleRepo *articleInfraStruct) FindArticlesByTopicId(articleIds []model.ArticleTopic, loginUserID uint, refPg int) (articles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10

	var articlesIDArr []uint

	// 構造体の配列からuintの配列に変換
	for _, v := range articleIds {
		articlesIDArr = append(articlesIDArr, v.ArticleID)
	}

	articleRepo.db.Raw(`
select 
	* 
from 
	(
		select 
			a.article_id, 
			a.article_title, 
			a.article_content, 
			group_concat(
				att.topic_name 
				order by 
					att.article_topic_id separator '/'
			) as article_topics, 
			a.created_user_id, 
			a.created_date, 
			a.updated_date, 
			a.deleted_date 
		from 
			articles as a, 
			(
				select 
					at.article_topic_id, 
					at.article_id, 
					at.topic_id, 
					t.topic_name 
				from 
					article_topics as at 
					left join topics as t on at.topic_id = t.topic_id
			) as att 
		where 
			a.article_id = att.article_id 
			and a.article_id in (?) 
			and is_deleted = 0 
		group by 
			a.article_id
	) as ddd 
limit 
	10 offset ?;
`, articlesIDArr, offset).Scan(&articles)

	// レコードがない場合
	if len(articles) == 0 {
		return nil, 1, errors.New("record not found")
	}

	var count int
	// article_topicsテーブルに削除フラグがないため(正確なレコード数が取得出来てない)
	// 以下の記述でレコードの数を取得
	articleRepo.db.Table("articles").Where("is_deleted = 0 AND article_id IN (?)", articlesIDArr).Count(&count)
	allPagingNum = (count / 11) + 1

	return
}

// 指定したトピックを含む記事トピックを取得
func (articleRepo *articleInfraStruct) FindArticleIdsByTopicId(topicID uint) (articleIds []model.ArticleTopic, err error) {
	// SELECT * FROM article_topics WHERE topic_id = :topicID AND is_deleted = 0;
	articleRepo.db.Where("topic_id = ?", topicID).Find(&articleIds)

	// レコードがない場合
	if len(articleIds) == 0 {
		return nil, errors.New("record not found")
	}

	return
}

// 最後の記事IDを取得
func (articleRepo *articleInfraStruct) FindLastArticleId() (lastArticleId uint, err error) {
	article := model.Article{}
	// SELECT article_id FROM articles WHERE is_deleted = 0 ORDER BY article_id DESC LIMIT 1;
	if result := articleRepo.db.Select("article_id").Last(&article); result.Error != nil {
		// レコードがない場合
		return 0, nil
	}

	lastArticleId = article.ArticleID
	return
}

// 記事のトピックが更新されているか確認
func (articleRepo *articleInfraStruct) CheckUpdateArticleTopic(willBeUpdatedArticle model.Article) bool {
	article := model.Article{}
	updateArticleId := willBeUpdatedArticle.ArticleID
	topicsStr := willBeUpdatedArticle.ArticleTopics

	articleRepo.db.Raw(`
select 
  a.article_id, 
  a.article_title, 
  a.article_content, 
  group_concat(
    att.topic_name 
    order by 
      att.article_topic_id
			 separator '/'
  ) as article_topics, 
  a.created_user_id, 
  a.created_date, 
  a.updated_date, 
  a.deleted_date 
from 
  articles as a, 
  (
    select 
      at.article_topic_id, 
      at.article_id, 
      at.topic_id, 
      t.topic_name 
    from 
      article_topics as at 
      left join topics as t on at.topic_id = t.topic_id
  ) as att 
where 
  a.article_id = att.article_id 
  and a.article_id = ?
  and is_deleted = 0 
group by 
  a.article_id;
	`, updateArticleId).Scan(&article)

	if article.ArticleTopics == topicsStr {
		// 記事トピックが更新されていない場合
		return false
	}
	// 記事トピックが更新されていた場合
	return true
}

// 記事を削除
func (articleRepo *articleInfraStruct) DeleteArticleByArticleId(articleId uint) (err error) {
	deleteArticle := model.Article{}
	// SELECT * FROM article WHERE article_id = :articleId AND is_deleted = 0;
	if result := articleRepo.db.Find(&deleteArticle, "article_id = ? AND is_deleted = ?", articleId, 0); result.Error != nil {
		// レコードがない場合
		err = result.Error
		return
	}

	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	deleteTime := time.Now().Format(dateFormat)
	customisedDeleteTime, _ := time.Parse(dateFormat, deleteTime)

	// 削除状態に更新
	articleRepo.db.Model(&deleteArticle).
		Where("article_id = ? AND is_deleted = ?", articleId, 0).
		Updates(map[string]interface{}{
			"deleted_date": customisedDeleteTime,
			"is_deleted":   int8(1),
		})

	return nil
}
