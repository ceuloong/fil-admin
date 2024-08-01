package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type NodesChart struct {
	ID                           uint            `gorm:"primarykey" json:"ID,omitempty"`
	Node                         string          `gorm:"type:varchar(255)" json:"node,omitempty"`
	AvailableBalance             decimal.Decimal `gorm:"type:decimal(20, 8)" json:"availableBalance"`
	LastAvailableBalance         decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastAvailableBalance"`
	Balance                      decimal.Decimal `gorm:"type:decimal(20, 8)" json:"balance"`
	LastBalance                  decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastBalance"`
	SectorPledgeBalance          decimal.Decimal `gorm:"type:decimal(20, 8)" json:"sectorPledgeBalance"`
	LastSectorPledgeBalance      decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastSectorPledgeBalance"`
	LastMonthSectorPledgeBalance decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastMonthSectorPledgeBalance"`
	VestingFunds                 decimal.Decimal `gorm:"type:decimal(20, 8)" json:"vestingFunds"`
	LastVestingFunds             decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastVestingFunds"`
	Height                       int             `gorm:"type:int" json:"height,omitempty"`
	LastTime                     time.Time       `gorm:"type:datetime" json:"lastTime"`
	RewardValue                  decimal.Decimal `gorm:"type:decimal(20, 8)" json:"rewardValue"`
	LastRewardValue              decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastRewardValue"`
	LastMonthRewardValue         decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastMonthRewardValue"`
	WeightedBlocks               int             `gorm:"type:int" json:"weightedBlocks,omitempty"`
	LastWeightedBlocks           int             `gorm:"type:int" json:"lastWeightedBlocks,omitempty"`
	LastMonthWeightedBlocks      int             `gorm:"type:int" json:"lastMonthWeightedBlocks,omitempty"`
	QualityAdjPower              decimal.Decimal `gorm:"type:decimal(20, 4);有效算力" json:"qualityAdjPower"`
	LastQualityAdjPower          decimal.Decimal `gorm:"type:decimal(20, 4);昨日有效算力" json:"lastQualityAdjPower"`
	LastMonthQualityAdjPower     decimal.Decimal `gorm:"type:decimal(20, 4);月初算力" json:"lastMonthQualityAdjPower"`
	PowerUnit                    string          `gorm:"type:varchar(50);算力单位" json:"powerUnit,omitempty"`
	PowerPoint                   decimal.Decimal `gorm:"type:decimal(10,3);算力占比" json:"powerPoint"`
	ControlBalance               decimal.Decimal `gorm:"type:decimal(20,8)" json:"controlBalance"`
	BlocksMined24h               int             `gorm:"type:int;24h报块数量" json:"blocksMined24H,omitempty"`
	TotalRewards24h              decimal.Decimal `gorm:"type:decimal(20, 8);24h出块奖励金额" json:"totalRewards24H"`
	LuckyValue24h                decimal.Decimal `gorm:"type:decimal(20, 8);24hLucky值" json:"luckyValue24H"`
	QualityAdjPowerDelta24h      decimal.Decimal `gorm:"type:decimal(20, 4);24h算力增量" json:"qualityAdjPowerDelta24H"`
	ReceiveAmount                decimal.Decimal `gorm:"type:decimal(20, 8);节点接收数量" json:"receiveAmount"`
	BurnAmount                   decimal.Decimal `gorm:"type:decimal(20, 8);节点销毁数量" json:"burnAmount"`
	SendAmount                   decimal.Decimal `gorm:"type:decimal(20, 8);节点发送数量" json:"sendAmount"`
	LastReceiveAmount            decimal.Decimal `gorm:"type:decimal(20, 8);前一天接收数量" json:"lastReceiveAmount"`
	LastBurnAmount               decimal.Decimal `gorm:"type:decimal(20, 4);前一天销毁数量" json:"lastBurnAmount"`
	LastSendAmount               decimal.Decimal `gorm:"type:decimal(20, 4);前一天提现数量" json:"lastSendAmount"`
	LastMonthReceiveAmount       decimal.Decimal `gorm:"type:decimal(20, 8);上月末节点接收数量" json:"lastMonthReceiveAmount"`
	LastMonthBurnAmount          decimal.Decimal `gorm:"type:decimal(20, 4);上月末销毁数量" json:"lastMonthBurnAmount"`
	LastMonthSendAmount          decimal.Decimal `gorm:"type:decimal(20, 4);上月末提现数量" json:"lastMonthSendAmount"`
	TimeTag                      int64           `gorm:"type:bigint" json:"timeTag"`
}

func (table *NodesChart) TableName() string {
	return "nodes_chart"
}

type NodeIndex struct {
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
}

type RankList struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type NodeChartData struct {
	X string  `json:"x"`
	Y float64 `json:"y"`
}

type NodesTotal struct {
	AvailableBalance        decimal.Decimal `json:"availableBalance"`
	Balance                 decimal.Decimal `json:"balance"`
	SectorPledgeBalance     decimal.Decimal `json:"sectorPledgeBalance"`
	VestingFunds            decimal.Decimal `json:"vestingFunds"`
	RewardValue             decimal.Decimal `json:"rewardValue"`
	WeightedBlocks          int             `json:"weightedBlocks"`
	QualityAdjPower         decimal.Decimal `json:"qualityAdjPower"`
	PowerPoint              decimal.Decimal `json:"powerPoint"`
	ControlBalance          decimal.Decimal `json:"controlBalance"`
	BlocksMined24h          int             `json:"blocksMined24H"`
	TotalRewards24h         decimal.Decimal `json:"totalRewards24H"`
	LuckyValue24h           decimal.Decimal `json:"luckyValue24H"`
	QualityAdjPowerDelta24h decimal.Decimal `json:"qualityAdjPowerDelta24H"`
	ReceiveAmount           decimal.Decimal `json:"receiveAmount"`
	BurnAmount              decimal.Decimal `json:"burnAmount"`
	SendAmount              decimal.Decimal `json:"sendAmount"`
	NodesList               *[]FilNodes     `json:"nodesList"`
	RoleId                  int             `json:"roleId"`
	TotalCount              int             `json:"totalCount"`
}
