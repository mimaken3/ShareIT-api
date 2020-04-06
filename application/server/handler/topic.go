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

		// ユーザのinterested_topicsにあるトピックを削除
		// userDeleteErr := userService.DeleteTopicFromInterestedTopics(uintTopicID)

		// if userDeleteErr != nil {
		// 	fmt.Println("cannot delete user's interested_topics")
		// }

		// articlesにあるトピックを削除

		// article_topicsにあるトピックを削除
		// willBeDeletedArticle := model.ArticleTopic{}
		// willBeDeletedArticle.TopicID = uintTopicID
		// at.articleTopicRepo.DeleteArticleTopic(willBeDeletedArticle)

		return c.String(http.StatusOK, "Success Delete!")
	}
}
