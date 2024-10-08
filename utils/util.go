package utils

import (
	"math"
	"time"

	"github.com/shopspring/decimal"
)

const (
	Layout = "2006-01-02 15:04:05"
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

func DecimalPowerValue(str string) (decimal.Decimal, string) {
	value, _ := decimal.NewFromString(str)
	onethousand := decimal.NewFromFloat(math.Pow10(3))
	if value.Abs().LessThan(onethousand) {
		return value, "GiB"
	} else {
		v := value.Div(decimal.NewFromFloat(math.Pow10(3))).RoundDown(2)
		if v.Abs().LessThan(onethousand) {
			return v, "TiB"
		} else {
			v = v.Div(decimal.NewFromFloat(1000)).RoundDown(2)
			return v, "PiB"
		}
	}
}

// ListPagination 分页
func ListPagination(count, page, pageSize int) (int, int) {
	index := 0
	end := pageSize
	if page > 1 {
		index = (page - 1) * pageSize
	}
	if index > count {
		index = count - pageSize
	}
	end = pageSize + index
	if end > count {
		end = count
	}
	return index, end
}
