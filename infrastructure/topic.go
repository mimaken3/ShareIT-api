package infrastructure

import (
	"errors"
	"strings"

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
		return 0, nil
	}

	lastTopicID = topic.TopicID
	return
}

// トピック名の重複チェック
func (topicRepo *topicInfraStruct) CheckTopicName(topicName string) (isDuplicated bool, message string, err error) {
	topic := model.Topic{}
	var sqlStr string

	// スペースが含まれているかチェック
	if strings.Contains(topicName, " ") {
		// スペース込み && 大文字小文字区別しない
		topicNameArray := strings.Fields(topicName)
		var sqlSlice = make([]byte, 0)
		sqlSlice = append(sqlSlice, "regexp '^"...)
		for i, value := range topicNameArray {
			if i == len(topicNameArray)-1 {
				sqlSlice = append(sqlSlice, value...)
			} else {
				sqlSlice = append(sqlSlice, value...)
				sqlSlice = append(sqlSlice, "( |　)*"...)
			}
		}
		sqlSlice = append(sqlSlice, "$'"...)
		sqlStr = string(sqlSlice)

	} else {
		// スペースなし && 大文字小文字区別しない
		sqlStr = "collate utf8_unicode_ci like '" + topicName + "'"
	}

	if result := topicRepo.db.Raw("select * from topics where is_deleted = 0 and convert(topic_name using utf8) " + sqlStr).Scan(&topic); result.Error != nil {
		// レコードがない場合
		isDuplicated = false
		message = ""
		return
	}

	// 重複しているレコードがあった場合
	isDuplicated = true
	message = topic.TopicName

	return
}

// トピックを登録
func (topicRepo *topicInfraStruct) CreateTopic(createTopic model.Topic, lastTopicID uint) (createdTopic model.Topic, err error) {
	// 現在の日付とデフォの削除日を取得
	currentDate, defaultDeletedDate := getDate()

	// DBに保存するトピックを準備
	createdTopic.TopicID = lastTopicID + 1
	createdTopic.TopicName = createTopic.TopicName
	createdTopic.ProposedUserID = createTopic.ProposedUserID
	createdTopic.CreatedDate = currentDate
	createdTopic.UpdatedDate = currentDate
	createdTopic.DeletedDate = defaultDeletedDate

	// 作成
	topicRepo.db.Create(&createdTopic)
	return
}

// 全トピックを取得
func (topicRepo *topicInfraStruct) FindAllTopics() (topics []model.Topic, err error) {

	if result := topicRepo.db.Where("is_deleted = ?", 0).Find(&topics); result.Error != nil {
		return nil, result.Error
	}

	if len(topics) == 0 {
		// レコードがない場合
		return nil, errors.New("record not found")
	}

	return
}

// トピック名を更新
func (topicRepo *topicInfraStruct) UpdateTopicNameByTopicID(topic model.Topic) (updatedTopic model.Topic, err error) {
	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	var _updatedTopic model.Topic

	// 更新
	topicRepo.db.Model(&_updatedTopic).
		Where("topic_id = ? AND is_deleted = ?", topic.TopicID, 0).
		Updates(map[string]interface{}{
			"topic_name":   topic.TopicName,
			"updated_date": currentDate,
		})

	// 更新したトピックを取得
	topicRepo.db.Where("topic_id = ? and is_deleted = 0 and proposed_user_id = ?", topic.TopicID, topic.ProposedUserID).Find(&updatedTopic)

	return
}

// トピックを削除
func (topicRepo *topicInfraStruct) DeleteTopicByTopicID(uintTopicID uint) (err error) {
	deleteTopic := model.Topic{}

	// SELECT * FROM topic WHERE topic_id = :uinttopicID AND is_deleted = 0;
	if result := topicRepo.db.Find(&deleteTopic, "topic_id = ? AND is_deleted = ?", uintTopicID, 0); result.Error != nil {
		// レコードがない場合
		err = result.Error
		return
	}

	// 現在の日付とデフォの削除日を取得
	currentDate, _ := getDate()

	// 削除状態に更新
	topicRepo.db.Model(&deleteTopic).
		Where("topic_id = ? AND is_deleted = ?", uintTopicID, 0).
		Updates(map[string]interface{}{
			"deleted_date": currentDate,
			"is_deleted":   int8(1),
		})

	return nil
}

// ユーザが作成したトピックを取得
func (topicRepo *topicInfraStruct) FindCreatedTopicsByUserID(userID uint) (topics []model.Topic, err error) {
	topics = []model.Topic{}

	if userID == 1 {
		topicRepo.db.Where("is_deleted = 0").Order("created_date desc").Find(&topics)
	} else {
		topicRepo.db.Where("proposed_user_id = ? and is_deleted = 0", userID).Order("created_date desc").Find(&topics)
	}
	return
}
