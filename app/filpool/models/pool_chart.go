package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type FilPoolChart struct {
	ID                  uint            `json:"id" gorm:"primarykey"`
	AvailableBalance    decimal.Decimal `json:"availableBalance" gorm:"type:decimal(20, 8)"`
	Balance             decimal.Decimal `json:"balance" gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance decimal.Decimal `json:"sectorPledgeBalance" gorm:"type:decimal(20, 8)"`
	VestingFunds        decimal.Decimal `json:"vestingFunds" gorm:"type:decimal(20, 8)"`
	LastTime            time.Time       `json:"lastTime" gorm:"type:datetime"`
	RewardValue         decimal.Decimal `json:"rewardValue" gorm:"type:decimal(20, 8)"`
	QualityAdjPower     decimal.Decimal `json:"qualityAdjPower" gorm:"type:decimal(20, 4);comment:有效算力"`
	PowerUnit           string          `json:"powerUnit" gorm:"type:varchar(50);comment:算力单位"`
	PowerPoint          decimal.Decimal `json:"powerPoint" gorm:"type:decimal(10,3);comment:算力占比"`
	ControlBalance      decimal.Decimal `json:"controlBalance" gorm:"type:decimal(20,8)"`
	DeptId              int             `json:"deptId" gorm:"type:int;comment:部门ID"`
	NodesCount          int             `json:"nodesCount" gorm:"type:int"`
}

func (table *FilPoolChart) TableName() string {
	return "pool_chart"
}

type PoolIndex struct {
	AvailableBalance    decimal.Decimal `json:"availableBalance" gorm:"type:decimal(20, 8)"`
	Balance             decimal.Decimal `json:"balance" gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance decimal.Decimal `json:"sectorPledgeBalance" gorm:"type:decimal(20, 8)"`
	VestingFunds        decimal.Decimal `json:"vestingFunds" gorm:"type:decimal(20, 8)"`
	LastTime            time.Time       `json:"lastTime" gorm:"type:datetime"`
	RewardValue         decimal.Decimal `json:"rewardValue" gorm:"type:decimal(20, 8)"`
	QualityAdjPower     decimal.Decimal `json:"qualityAdjPower" gorm:"type:decimal(20, 4);comment:有效算力"`
	PowerUnit           string          `json:"powerUnit" gorm:"type:varchar(50);comment:算力单位"`
	PowerPoint          decimal.Decimal `json:"powerPoint" gorm:"type:decimal(10,3);comment:算力占比"`
	ControlBalance      decimal.Decimal `json:"controlBalance" gorm:"type:decimal(20,8)"`
	MonthAvg            decimal.Decimal `json:"monthAvg"`
	WeekIncrease        decimal.Decimal `json:"weekIncrease"`
	DayIncrease         decimal.Decimal `json:"dayIncrease"`
	WeekTop             string          `json:"weekTop"`
	DayTop              string          `json:"dayTop"`
}

type AppPoolIndex struct {
	QualityAdjPower decimal.Decimal `json:"qualityAdjPower" gorm:"type:decimal(20, 8)"`
	NodesCount      int             `json:"nodesCount" gorm:"type:int"`
}

type PoolFinance struct {
	AvailableBalance    decimal.Decimal `json:"availableBalance" gorm:"type:decimal(20, 8)"`
	Balance             decimal.Decimal `json:"balance" gorm:"type:decimal(20, 8)"`
	SectorPledgeBalance decimal.Decimal `json:"sectorPledgeBalance" gorm:"type:decimal(20, 8)"`
	VestingFunds        decimal.Decimal `json:"vestingFunds" gorm:"type:decimal(20, 8)"`
	BlocksMined24h      int             `gorm:"type:int;24h报块数量" json:"blocksMined24H,omitempty"`
	TotalRewards24h     decimal.Decimal `gorm:"type:decimal(20, 8);24h出块奖励金额" json:"totalRewards24H"`
	NewlyPrice          decimal.Decimal `json:"newlyPrice"`
}

type BarChart struct {
	X string  `json:"x"`
	Y float64 `json:"y"`
}

type AppBarChart struct {
	X int64   `json:"x"`
	Y float64 `json:"y"`
}
