package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 重複したトピックのJSON用
type DuplicatedTopic struct {
	IsDuplicated bool   `json:"is_duplicated"`
	Message      string `json:"message"`
}

// トピック名の重複チェック
// func CheckTopicName() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		topic := model.Topic{}
// 		c.Bind(&topic)
//
// 		isDuplicated, message, err := topicService.CheckTopicName(topic.TopicName)
// 		if err != nil {
// 			return err
// 		}
// 		return c.String(http.StatusOK, "isDuplicated:"+strconv.FormatBool(isDuplicated)+" message:"+message)
// 	}
// }

// トピックを作成
func CreateTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		topic := model.Topic{}
		c.Bind(&topic)

		// トピック名の重複チェック
		isDuplicated, message, err := topicService.CheckTopicName(topic.TopicName)
		if isDuplicated {
			// 重複していたらエラーメッセージを返す
			dt := DuplicatedTopic{
				IsDuplicated: isDuplicated,
				Message:      message,
			}
			return c.JSON(http.StatusBadRequest, dt)
		}

		createdTopic, err := topicService.CreateTopic(topic)

		if err != nil {
			return c.JSON(http.StatusBadRequest, createdTopic)
		}

		return c.JSON(http.StatusOK, createdTopic)
	}
}

// 全トピックを取得
func FindAllTopics() echo.HandlerFunc {
	return func(c echo.Context) error {

		topics, err := topicService.FindAllTopics()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, topics)
	}
}

// トピックを削除
func DeleteTopicByTopicID() echo.HandlerFunc {
	return func(c echo.Context) error {
		topicID, _ := strconv.Atoi(c.Param("topic_id"))

		// intをuintに変換
		var uintTopicID uint = uint(topicID)

		// トピックを削除
		err := topicService.DeleteTopicByTopicID(uintTopicID)

		if err != nil {
			return c.String(http.StatusBadRequest, "Cannot delete topic")
		}

		// ユーザの興味トピックを削除
		errUser := userInterestedTopicService.DeleteUserTopicByTopicID(topicID)

		fmt.Println(errUser)

		// トピックに紐づく記事トピックを削除
		errAT := articleTopicService.DeleteArticleTopicByTopicID(uintTopicID)

		fmt.Println(errAT)

		return c.String(http.StatusOK, "Success Delete!")
	}
}
