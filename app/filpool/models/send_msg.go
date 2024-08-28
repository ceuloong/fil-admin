package models

import (
	"fil-admin/common/models"
	"time"
)

type SendType int

const (
	SectorsError SendType = 101 // 扇区错误
	HeightDelay  SendType = 102 // 高度延迟
	LuckyLow     SendType = 103 // 幸运值过低
	OrphanBlock  SendType = 104 // 孤块
	OffLine      SendType = 105 // 离线
)

type SendMsg struct {
	models.Model
	Title      string     `gorm:"type:varchar(255)" json:"title"`
	Node       string     `gorm:"type:varchar(30)" json:"node"`
	Content    string     `gorm:"type:varchar(255)" json:"content"`
	CreateTime time.Time  `gorm:"type:datetime" json:"createTime"`
	SendTime   *time.Time `gorm:"type:datetime" json:"sendTime"`
	Type       SendType   `gorm:"type:int" json:"type"`
	TypeStr    string     `gorm:"-" json:"typeStr"`
	TimeShow   string     `gorm:"-" json:"timeShow"`
	SendStatus int        `gorm:"type:int" json:"sendStatus"`
	// models.ModelTime
	// models.ControlBy
}

type ShowSendMsg struct {
	Id         int      `json:"id"`
	Title      string   `gorm:"type:varchar(255)" json:"title"`
	Node       string   `gorm:"type:varchar(30)" json:"node"`
	Content    string   `gorm:"type:varchar(255)" json:"content"`
	Type       SendType `gorm:"type:int" json:"type"`
	TypeStr    string   `gorm:"-" json:"typeStr"`
	TimeShow   string   `gorm:"-" json:"timeShow"`
	SendStatus int      `gorm:"type:int" json:"sendStatus"`
}

func (SendMsg) TableName() string {
	return "send_msg"
}

func (e *SendMsg) GetId() interface{} {
	return e.Id
}

func (e *SendMsg) GetTypeStr() interface{} {
	switch e.Type {
	case SectorsError:
		return "算力异常"
	case HeightDelay:
		return "高度延迟"
	case LuckyLow:
		return "幸运值过低"
	case OrphanBlock:
		return "孤块"
	case OffLine:
		return "离线"
	default:
		return ""
	}
}

// 当天只显示时间，昨天显示昨天，其他显示日期
func (e *SendMsg) ShowTimeStr() string {
	now := time.Now()
	if e.CreateTime.Year() == now.Year() && e.CreateTime.Month() == now.Month() && e.CreateTime.Day() == now.Day() {
		return e.CreateTime.Format("15:04:05")
	}
	// if e.CreateTime.Year() == now.Year() && e.CreateTime.Month() == now.Month() && e.CreateTime.Day() == now.Day()-1 {
	// 	return "昨天"
	// }
	// if e.CreateTime.Year() == now.Year() && e.CreateTime.Month() == now.Month() && e.CreateTime.Day() == now.Day()-2 {
	// 	return "前天"
	// }
	// if e.CreateTime.Year() == now.Year() && e.CreateTime.Month() == now.Month() && e.CreateTime.Day() == now.Day()-7 {
	// 	return "一周前"
	// }
	return e.CreateTime.Format("2006-01-02 15:04:05")
}
