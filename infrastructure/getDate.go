package infrastructure

import (
	"time"

	"google.golang.org/appengine"
)

// 現在の日付とデフォの削除日を取得
func getDate() (currentDate time.Time, defaultDeletedDate time.Time) {
	var _nowTime time.Time
	if appengine.IsAppEngine() {
		// GAEだとタイムゾーンがUTCなのでJSTにする
		_nowTime = time.Now().Add(time.Hour * 9)
	} else {
		// Local
		_nowTime = time.Now()
	}

	const dateFormat = "2006-01-02 15:04:05"
	nowTime := _nowTime.Format(dateFormat)
	currentDate, _ = time.Parse(dateFormat, nowTime)

	const defaultDeletedDateStr = "9999-12-31 23:59:59"
	defaultDeletedDate, _ = time.Parse(dateFormat, defaultDeletedDateStr)

	return
}
