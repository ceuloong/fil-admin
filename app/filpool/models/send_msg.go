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
	SendStatus int        `gorm:"type:int" json:"sendStatus"`
	// models.ModelTime
	// models.ControlBy
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
		return "扇区错误"
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
