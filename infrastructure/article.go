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
	IsPrivate      int8      `json:"is_private"`
	IsDeleted      int8      `json:"-"`
}

func (CreateArticle) TableName() string {
	return "articles"
}

// 全記事を取得(ページング)
func (articleRepo *articleInfraStruct) FindAllArticles(refPg int, userID uint) (articles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10
	rows, err :=
		articleRepo.db.Raw(
			`
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
      sub_a2.* 
    from 
      articles as sub_a2
      inner join (
        select 
          case 
          when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
          when is_private = 0 and is_deleted = 0 then article_id end as article_id 
        from 
          articles 
        having 
          article_id is not null 
        limit 
          10 offset ?
      ) as sub_a on sub_a2.article_id = sub_a.article_id
  ) as a, -- ユーザの公開/非公開を考慮したarticlesの10件
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
group by 
  a.article_id
;
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
	row := articleRepo.db.Raw(`
select 
  count(*) 
from 
  (
    select 
      case 
      when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
      when is_private = 0 and is_deleted = 0 then article_id 
      end as article_id 
    from 
      articles 
    having 
      article_id is not null
  ) as t;
	`, userID).Row()
	row.Scan(&count)
	allPagingNum = (count / 11) + 1

	return
}

// 記事を検索(ページング)
func (articleRepo *articleInfraStruct) SearchAllArticles(refPg int, userID uint, loginUserID uint, topicIDs []uint) (searchedArticles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10

	if userID == 0 && topicIDs[0] != 0 {
		// 「全ユーザ」かつ「特定トピック」の場合
		rows, err :=
			articleRepo.db.Raw(`
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
  a.deleted_date, 
  a.is_private 
from 
  (
    select 
      * 
    from 
      (
        select 
          a.* 
        from 
          articles as a 
          inner join (
            select 
              case 
                when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
                when is_private = 0 and is_deleted = 0 then article_id 
              end as article_id 
            from 
              articles 
            having 
              article_id is not null
          ) as sub_a on a.article_id = sub_a.article_id
      ) as articles 
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
group by 
  a.article_id 
limit 
  10 offset ? 
;	
		`, loginUserID, topicIDs, offset).Rows()
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
  (
    select 
      a.* 
    from 
      articles as a 
      inner join (
        select 
          case 
          	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
          	when is_private = 0 and is_deleted = 0 then article_id 
          end as article_id 
        from 
          articles 
        having 
          article_id is not null
      ) as sub_a on a.article_id = sub_a.article_id
  ) as articles 
where 
  is_deleted = 0 
  and article_id in (
    select 
      article_id 
    from 
      article_topics 
    where 
      topic_id in (?)
  )
;
	`, loginUserID, topicIDs).Row()
		row.Scan(&count)
		allPagingNum = (count / 11) + 1

		return searchedArticles, allPagingNum, nil
	}
	// 「特定のユーザ」かつ「特定のトピック」の場合
	rows, err :=
		articleRepo.db.Raw(`
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
  a.deleted_date,
  a.is_private
from 
  (
    select 
      * 
    from 
      (
        select 
          a.* 
        from 
          articles as a 
          inner join (
            select 
              case 
								when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
								when is_private = 0 and is_deleted = 0 then article_id 
							end as article_id 
            from 
              articles 
            having 
              article_id is not null
          ) as sub_a on a.article_id = sub_a.article_id
      ) as articles 
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
limit 
  10 offset ? 
;	
`, loginUserID, topicIDs, userID, offset).Rows()
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
  (
    select 
      a.* 
    from 
      articles as a 
      inner join (
        select 
          case 
          	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
          	when is_private = 0  and is_deleted = 0 then article_id 
          end as article_id 
        from 
          articles 
        having 
          article_id is not null
      ) as sub_a on a.article_id = sub_a.article_id
  ) as articles 
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
  )
;
	`, loginUserID, userID, topicIDs).Row()
	row.Scan(&count)
	allPagingNum = (count / 11) + 1

	return searchedArticles, allPagingNum, nil
}

// 記事を取得
func (articleRepo *articleInfraStruct) FindArticleByArticleId(loginUserID uint, articleId uint) (article model.Article, err error) {
	result := articleRepo.db.Raw(`
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
  a.deleted_date, 
  a.is_private 
from 
  (
    select 
      a.* 
    from 
      articles as a 
      inner join (
        select 
          case 
          	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
          	when is_private = 0 and is_deleted = 0 then article_id 
          end as article_id 
        from 
          articles 
        having 
          article_id is not null
      ) as sub_a on a.article_id = sub_a.article_id
  ) as a, 
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
  a.article_id
;   
	`, loginUserID, articleId).Scan(&article)

	if result.Error != nil {
		// レコードがない場合
		err = result.Error
	}
	return
}

// 記事を投稿
func (articleRepo *articleInfraStruct) CreateArticle(createArticle model.Article, lastArticleId uint) (createdArticle model.Article, err error) {
	// 現在の日付とデフォの削除日を取得
	currentDate, defaultDeletedDate := getDate()

	ar := CreateArticle{}

	// DBに保存する記事のモデルを作成
	ar.ArticleID = lastArticleId + 1
	ar.ArticleTitle = createArticle.ArticleTitle
	ar.ArticleContent = createArticle.ArticleContent
	ar.CreatedUserID = createArticle.CreatedUserID
	ar.CreatedDate = currentDate
	ar.UpdatedDate = currentDate
	ar.DeletedDate = defaultDeletedDate
	ar.IsPrivate = createArticle.IsPrivate

	articleRepo.db.Create(&ar)

	// DBに保存した記事を返す
	createdArticle.ArticleID = lastArticleId + 1
	createdArticle.ArticleTitle = createArticle.ArticleTitle
	createdArticle.ArticleContent = createArticle.ArticleContent
	createdArticle.ArticleTopics = createArticle.ArticleTopics
	createdArticle.CreatedUserID = createArticle.CreatedUserID
	createdArticle.CreatedDate = currentDate
	createdArticle.UpdatedDate = currentDate
	createdArticle.DeletedDate = defaultDeletedDate
	createdArticle.IsPrivate = createArticle.IsPrivate

	return
}

// 記事を更新
func (articleRepo *articleInfraStruct) UpdateArticleByArticleId(willBeUpdatedArticle model.Article) (updatedArticle model.Article, err error) {
	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	updateId := willBeUpdatedArticle.ArticleID

	// 更新するフィールドを設定
	updateTitle := willBeUpdatedArticle.ArticleTitle
	updateContent := willBeUpdatedArticle.ArticleContent
	updateIsPrivate := willBeUpdatedArticle.IsPrivate

	// 更新
	articleRepo.db.Model(&updatedArticle).
		Where("article_id = ?", updateId).
		Updates(map[string]interface{}{
			"article_title":   updateTitle,
			"article_content": updateContent,
			"updated_date":    currentDate,
			"is_private":      updateIsPrivate,
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
	updatedArticle.IsPrivate = willBeUpdatedArticle.IsPrivate

	return
}

// 特定のユーザの全記事を取得
func (articleRepo *articleInfraStruct) FindArticlesByUserId(userID uint, loginUserID uint, refPg int) (articles []model.Article, allPagingNum int, err error) {
	offset := (refPg - 1) * 10
	rows, err :=
		articleRepo.db.Raw(`
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
  a.deleted_date, 
  a.is_private 
from 
  (
    select 
      * 
    from 
      (
        select 
          a.* 
        from 
          articles as a 
          inner join (
            select 
              case 
              	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
              	when is_private = 0 and is_deleted = 0 then article_id
              end as article_id 
            from 
              articles 
            having 
              article_id is not null
          ) as sub_a on a.article_id = sub_a.article_id
      ) as articles 
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
limit 
  10 offset ? 
;
`, loginUserID, userID, offset).Rows()

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
	row := articleRepo.db.Raw(`
select 
  count(*) 
from 
  (
    select 
      a.* 
    from 
      articles as a 
      inner join (
        select 
          case 
          	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
          	when is_private = 0  and is_deleted = 0 then article_id 
          end as article_id 
        from 
          articles 
        having 
          article_id is not null
      ) as sub_a on a.article_id = sub_a.article_id
  ) as articles 
where 
  is_deleted = 0 
  and created_user_id = ? 
  and article_id in (
    select 
      article_id 
    from 
      article_topics 
  )
;
	`, loginUserID, userID).Row()
	row.Scan(&count)
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
      a.deleted_date, 
      a.is_private 
    from 
      (
        select 
          a.* 
        from 
          articles as a 
          inner join (
            select 
              case 
              	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id
                when is_private = 0 and is_deleted = 0 then article_id 
              end as article_id 
            from 
              articles 
            having 
              article_id is not null
          ) as sub_a on a.article_id = sub_a.article_id
      ) as a, 
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
  10 offset ? 
;
`, loginUserID, articlesIDArr, offset).Scan(&articles)

	// レコードがない場合
	if len(articles) == 0 {
		return nil, 1, errors.New("record not found")
	}

	var count int
	row := articleRepo.db.Raw(`
select 
  count(*) 
from 
  (
    select 
      a.* 
    from 
      articles as a 
      inner join (
        select 
          case 
          	when is_private = 1 and created_user_id = ? and is_deleted = 0 then article_id 
          	when is_private = 0 and is_deleted = 0 then article_id 
          end as article_id 
        from 
          articles 
        having 
          article_id is not null
      ) as sub_a on a.article_id = sub_a.article_id
  ) as articles 
where 
  article_id in(?)
;
	`, loginUserID, articlesIDArr).Row()
	row.Scan(&count)
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

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	// 削除状態に更新
	articleRepo.db.Model(&deleteArticle).
		Where("article_id = ? AND is_deleted = ?", articleId, 0).
		Updates(map[string]interface{}{
			"deleted_date": currentDate,
			"is_deleted":   int8(1),
		})

	return nil
}

// ユーザの記事を全削除
func (articleRepo *articleInfraStruct) DeleteArticleByUserID(userID uint) (err error) {
	deleteArticle := []model.Article{}

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	// 削除状態に更新
	articleRepo.db.Model(&deleteArticle).
		Where("created_user_id = ? AND is_deleted = ?", userID, 0).
		Updates(map[string]interface{}{
			"deleted_date": currentDate,
			"is_deleted":   int8(1),
		})

	return nil
}
