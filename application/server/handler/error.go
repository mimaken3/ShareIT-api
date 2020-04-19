package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mimaken3/ShareIT-api/domain/model"
)

// ErrorCode is error type to be used shareIT service api
type ErrorCode string

func (ec ErrorCode) String() string {
	return string(ec)
}

type ArticleEchoContext struct {
	echo.Context
}

// Error codes managed by us.
const (
	AuthenticationParamMissing ErrorCode = "0"
	AuthenticationFailure      ErrorCode = "1"
	InvalidParameter           ErrorCode = "2"
	InternalError              ErrorCode = "3"
	NotFoundRecord             ErrorCode = "4"

	UnHandledError ErrorCode = "999"
)

var codeStatusMap = map[ErrorCode]int{
	AuthenticationFailure:      http.StatusForbidden,
	AuthenticationParamMissing: http.StatusBadRequest,
	InvalidParameter:           http.StatusBadRequest,
	InternalError:              http.StatusInternalServerError,
	NotFoundRecord:             http.StatusNotFound,
}

// GetHTTPStatus returns http status that corresponds to the given ErrorCode
func GetHTTPStatus(code ErrorCode) int {
	return codeStatusMap[code]
}

func (c *ArticleEchoContext) ErrorResponseFunc(err error) echo.HandlerFunc {
	return func(c echo.Context) error {
		er := &model.ErrorResponse{}
		errorMsg := err.Error()
		if errorMsg == "record not found" {
			er.Code = GetHTTPStatus("4")
			// er.Errors = errorMsg
			er.Errors = append(er.Errors, errorMsg)
			// return c.JSON(GetHTTPStatus("4"), er)
			// return c.JSON(404, er)
			return c.String(http.StatusBadRequest, "Request is failed because of "+err.Error())
		}
		// return c.String(http.StatusBadRequest, "不明なエラー")
		return err
	}
}
