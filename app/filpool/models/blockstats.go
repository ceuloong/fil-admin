package models

import (
	"fil-admin/common/models"
	"time"
)

type BlockStats struct {
	models.Model

	Node                  string    `json:"node" gorm:"type:varchar(50);comment:Node"`
	BlocksGrowth          int       `json:"blocksGrowth" gorm:"type:int;comment:BlocksGrowth"`
	BlocksRewardGrowthFil string    `json:"blocksRewardGrowthFil" gorm:"type:decimal(20,8);comment:BlocksRewardGrowthFil"`
	HeightTimeStr         string    `json:"heightTimeStr" gorm:"type:varchar(50);comment:HeightTimeStr"`
	HeightTime            time.Time `json:"heightTime" gorm:"type:datetime;comment:HeightTime"`

	models.ModelTime
	models.ControlBy
}

func (BlockStats) TableName() string {
	return "block_stats"
}

func (e *BlockStats) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *BlockStats) GetId() interface{} {
	return e.Id
}

type PoolBlockStats struct {
	BlocksGrowth          int       `json:"blocksGrowth" gorm:"type:int;comment:BlocksGrowth"`
	BlocksRewardGrowthFil string    `json:"blocksRewardGrowthFil" gorm:"type:decimal(20,8);comment:BlocksRewardGrowthFil"`
	HeightTimeStr         string    `json:"heightTimeStr" gorm:"type:varchar(50);comment:HeightTimeStr"`
	HeightTime            time.Time `json:"heightTime" gorm:"type:datetime;comment:HeightTime"`
}
