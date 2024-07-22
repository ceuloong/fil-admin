package handler

import (
	"github.com/shopspring/decimal"
	"time"
)

type NodesChart struct {
	ID                      uint            `gorm:"primarykey" json:"ID,omitempty"`
	Node                    string          `gorm:"type:varchar(255)" json:"node,omitempty"`
	AvailableBalance        decimal.Decimal `gorm:"type:decimal(20, 8)" json:"availableBalance"`
	Balance                 decimal.Decimal `gorm:"type:decimal(20, 8)" json:"balance"`
	SectorPledgeBalance     decimal.Decimal `gorm:"type:decimal(20, 8)" json:"sectorPledgeBalance"`
	VestingFunds            decimal.Decimal `gorm:"type:decimal(20, 8)" json:"vestingFunds"`
	RewardValue             decimal.Decimal `gorm:"type:decimal(20, 8)" json:"rewardValue"`
	WeightedBlocks          int             `gorm:"type:int" json:"weightedBlocks,omitempty"`
	QualityAdjPower         decimal.Decimal `gorm:"type:decimal(20, 4);有效算力" json:"qualityAdjPower"`
	PowerPoint              decimal.Decimal `gorm:"type:decimal(10,3);算力占比" json:"powerPoint"`
	BlocksMined24h          int             `gorm:"type:int;24h报块数量" json:"blocksMined24H,omitempty"`
	TotalRewards24h         decimal.Decimal `gorm:"type:decimal(20, 8);24h出块奖励金额" json:"totalRewards24H"`
	LuckyValue24h           decimal.Decimal `gorm:"type:decimal(20, 8);24hLucky值" json:"luckyValue24H"`
	QualityAdjPowerDelta24h decimal.Decimal `gorm:"type:decimal(20, 4);24h算力增量" json:"qualityAdjPowerDelta24H"`
	ReceiveAmount           decimal.Decimal `gorm:"type:decimal(20, 8);节点接收数量" json:"receiveAmount"`
	BurnAmount              decimal.Decimal `gorm:"type:decimal(20, 8);节点销毁数量" json:"burnAmount"`
	SendAmount              decimal.Decimal `gorm:"type:decimal(20, 8);节点发送数量" json:"sendAmount"`
	LastTime                time.Time       `json:"lastTime"`
}

func (table *NodesChart) TableName() string {
	return "nodes_chart"
}
