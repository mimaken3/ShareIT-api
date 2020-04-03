package handler

import (
	"net/http"

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
