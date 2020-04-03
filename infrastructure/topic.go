package infrastructure

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mimaken3/ShareIT-api/domain/model"
	"github.com/mimaken3/ShareIT-api/domain/repository"
)

type topicInfraStruct struct {
	db *gorm.DB
}

// DIのための関数
func NewTopicDB(db *gorm.DB) repository.TopicRepository {
	return &topicInfraStruct{db: db}
}

// 最後のトピックIDを取得
func (topicRepo *topicInfraStruct) FindLastTopicID() (lastTopicID uint, err error) {
	topic := model.Topic{}

	// SELECT topic_id FROM topics WHERE is_deleted = 0 ORDER BY topic_id DESC LIMIT 1;
	if result := topicRepo.db.Select("topic_id").Where("is_deleted = ?", 0).Last(&topic); result.Error != nil {
		// レコードがない場合
		err = result.Error
		return
	}

	lastTopicID = topic.TopicID
	return
}

// トピックを登録
func (topicRepo *topicInfraStruct) CreateTopic(createTopic model.Topic, lastTopicID uint) (createdTopic model.Topic, err error) {
	// 現在の日付を取得
	const dateFormat = "2006-01-02 15:04:05"
	nowTime := time.Now().Format(dateFormat)
	customisedNowTime, _ := time.Parse(dateFormat, nowTime)

	const defaultDeletedDateStr = "9999-12-31 23:59:59"
	defaultDeletedDate, _ := time.Parse(dateFormat, defaultDeletedDateStr)

	// DBに保存するトピックを準備
	createdTopic.TopicID = lastTopicID + 1
	createdTopic.TopicName = createTopic.TopicName
	createdTopic.ProposedUserID = createTopic.ProposedUserID
	createdTopic.CreatedDate = customisedNowTime
	createdTopic.UpdatedDate = customisedNowTime
	createdTopic.DeletedDate = defaultDeletedDate

	// 作成
	topicRepo.db.Create(&createdTopic)
	return
}
