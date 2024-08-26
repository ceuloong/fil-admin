package models

import (
	"fil-admin/common/models"
	"time"
)

type Block struct {
	models.Model

	Height      int       `json:"height" gorm:"type:int;comment:高度"`
	Node        string    `json:"node" gorm:"type:varchar(255);comment:节点"`
	BlockTime   time.Time `json:"blockTime" gorm:"type:datetime;comment:区块时间"`
	NodeFrom    string    `json:"nodeFrom" gorm:"type:varchar(255);comment:NodeFrom"`
	NodeTo      string    `json:"nodeTo" gorm:"type:varchar(255);comment:NodeTo"`
	Message     string    `json:"message" gorm:"type:varchar(255);comment:区块哈希"`
	RewardValue string    `json:"rewardValue" gorm:"type:decimal(20,8);comment:奖励数量"`
	MsgCount    string    `json:"msgCount" gorm:"type:int;comment:消息数"`
	BlockSize   string    `json:"blockSize" gorm:"type:bigint;comment:区块大小"`
	Status      string    `json:"status" gorm:"type:int;comment:1正常   2孤块"`
	CreateTime  time.Time `json:"createTime" gorm:"type:datetime;comment:记录的创建时间"`
}

type BlockShow struct {
	Height      int    `json:"height" gorm:"type:int;comment:高度"`
	Node        string `json:"node" gorm:"type:varchar(255);comment:节点"`
	BlockTime   string `json:"blockTime" gorm:"type:datetime;comment:区块时间"`
	Message     string `json:"message" gorm:"type:varchar(255);comment:区块哈希"`
	RewardValue string `json:"rewardValue" gorm:"type:decimal(20,8);comment:奖励数量"`
	Status      string `json:"status" gorm:"type:int;comment:1正常   2孤块"`
}

func (Block) TableName() string {
	return "block"
}

func (e *Block) GetId() interface{} {
	return e.Id
}

func (e *Block) Generate() BlockShow {
	o := BlockShow{
		Height:      e.Height,
		Node:        e.Node,
		BlockTime:   e.BlockTime.Format("2006-01-02 15:04:05"),
		Message:     e.Message,
		RewardValue: e.RewardValue,
		Status:      e.Status,
	}
	return o
}
