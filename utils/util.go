package utils

import (
	"math"
	"time"

	"github.com/shopspring/decimal"
)

const (
	Layout = "2006-01-02 15:04:05"
)

// GetFirstDayOfMonth 获取当月第一天，如果当前是1号0点则返回上月第一天
func GetFirstDayOfMonth(now time.Time) time.Time {
	// 判断是否为当月1号0点
	isFirstDayMidnight := now.Day() == 1 && now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0

	if isFirstDayMidnight {
		// 如果是1号0点，返回上月第一天
		return time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())
	}

	// 否则返回当月第一天
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

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
