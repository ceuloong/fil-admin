package utils

import (
	"github.com/shopspring/decimal"
	"time"
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
