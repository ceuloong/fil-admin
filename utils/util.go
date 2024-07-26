package utils

import (
	"time"

	"github.com/shopspring/decimal"
)

func DateFormatStr(t time.Time, timeTemplate string) string {
	//timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := t.Format(timeTemplate)
	return timeStr
}

func DecimalValue(str string) decimal.Decimal {
	value, _ := decimal.NewFromString(str)
	//v := value.Div(decimal.NewFromFloat(math.Pow10(18))).Round(4)
	return value
}

// SetTime 设置时间的小时数
func SetTime(now time.Time, hour int) time.Time {
	// 获取当前时间的年份和月份
	year, month, day := now.Date()
	return time.Date(year, month, day, hour, 0, 0, 0, now.Location())
}

func UnixTimeToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
