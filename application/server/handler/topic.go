package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// トピックを作成
func CreateTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		topic := model.Topic{}
		c.Bind(&topic)

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

		// ユーザのinterested_topicsを削除
		errUser := userInterestedTopicService.DeleteUserTopicByTopicID(topicID)

		fmt.Println(errUser)

		// トピックに紐づく記事トピックを削除
		errAT := articleTopicService.DeleteArticleTopicByTopicID(uintTopicID)

		fmt.Println(errAT)

		// ユーザのinterested_topicsにあるトピックを削除
		// userDeleteErr := userService.DeleteTopicFromInterestedTopics(uintTopicID)

		// if userDeleteErr != nil {
		// 	fmt.Println("cannot delete user's interested_topics")
		// }

		return c.String(http.StatusOK, "Success Delete!")
	}
}
