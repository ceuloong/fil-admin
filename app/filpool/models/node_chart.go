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
	PowerUnit               string          `json:"powerUnit"`
	PowerPoint              decimal.Decimal `json:"powerPoint"`
	ControlBalance          decimal.Decimal `json:"controlBalance"`
	BlocksMined24h          int             `json:"blocksMined24H"`
	TotalRewards24h         decimal.Decimal `json:"totalRewards24H"`
	LuckyValue24h           decimal.Decimal `json:"luckyValue24H"`
	QualityAdjPowerDelta24h decimal.Decimal `json:"qualityAdjPowerDelta24H"`
	MiningEfficiency        decimal.Decimal `json:"miningEfficiency"`
	ReceiveAmount           decimal.Decimal `json:"receiveAmount"`
	BurnAmount              decimal.Decimal `json:"burnAmount"`
	SendAmount              decimal.Decimal `json:"sendAmount"`
	NodesList               *[]FilNodes     `json:"nodesList"`
	RoleId                  int             `json:"roleId"`
	TotalCount              int             `json:"totalCount"`
	PowerDeltaUnit          string          `json:"powerDeltaUnit"`
	PowerDeltaShow          string          `json:"powerDeltaShow"`
}

func (e *NodesTotal) SetScale(total NodesTotal) NodesTotal {
	total.AvailableBalance = total.AvailableBalance.RoundDown(2)
	total.Balance = total.Balance.RoundDown(2)
	total.SectorPledgeBalance = total.SectorPledgeBalance.RoundDown(2)
	total.VestingFunds = total.VestingFunds.RoundDown(2)
	total.RewardValue = total.RewardValue.RoundDown(2)
	total.TotalRewards24h = total.TotalRewards24h.RoundDown(2)
	total.LuckyValue24h = total.LuckyValue24h.RoundDown(4)
	total.MiningEfficiency = total.MiningEfficiency.Mul(decimal.NewFromInt(1000)).RoundDown(1)
	total.QualityAdjPowerDelta24h = total.QualityAdjPowerDelta24h.RoundDown(2)

	return total
}

type NodesChartWithFilNodes struct {
	Node                          string          `gorm:"type:varchar(255)" json:"node,omitempty"`
	MsigNode                      string          `json:"msigNode" gorm:"type:varchar(255);comment:所属地址"`
	Type                          int             `json:"type" gorm:"type:int;comment:节点类型 联合类型"`
	DistributePoint               decimal.Decimal `json:"distributePoint" gorm:"type:decimal(10,3);comment:联合挖矿分配比例"`
	QualityAdjPower               decimal.Decimal `gorm:"type:decimal(20, 4);有效算力" json:"qualityAdjPower"`
	LastQualityAdjPower           decimal.Decimal `gorm:"type:decimal(20, 4);昨日有效算力" json:"lastQualityAdjPower"`
	LastMonthQualityAdjPower      decimal.Decimal `gorm:"type:decimal(20, 4);月初算力" json:"lastMonthQualityAdjPower"`
	QualityAdjPowerDelta24h       decimal.Decimal `gorm:"type:decimal(20, 4);24h算力增量" json:"qualityAdjPowerDelta24H"`
	QualityAdjPowerDeltaMonth     decimal.Decimal `gorm:"type:decimal(20, 4);月算力增量" json:"qualityAdjPowerDeltaMonth"`
	AvailableBalance              decimal.Decimal `gorm:"type:decimal(20, 8)" json:"availableBalance"`
	LastAvailableBalance          decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastAvailableBalance"`
	Balance                       decimal.Decimal `gorm:"type:decimal(20, 8)" json:"balance"`
	LastBalance                   decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastBalance"`
	SectorPledgeBalance           decimal.Decimal `gorm:"type:decimal(20, 8)" json:"sectorPledgeBalance"`
	LastSectorPledgeBalance       decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastSectorPledgeBalance"`
	LastMonthSectorPledgeBalance  decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastMonthSectorPledgeBalance"`
	SectorPledgeBalanceDeltaMonth decimal.Decimal `gorm:"type:decimal(20, 8);月质押量增量" json:"sectorPledgeBalanceDeltaMonth"`
	VestingFunds                  decimal.Decimal `gorm:"type:decimal(20, 8)" json:"vestingFunds"`
	LastVestingFunds              decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastVestingFunds"`
	BlocksMined24h                int             `gorm:"type:int;24h报块数量" json:"blocksMined24H,omitempty"`
	TotalRewards24h               decimal.Decimal `gorm:"type:decimal(20, 8);24h出块奖励金额" json:"totalRewards24H"`
	LuckyValue24h                 decimal.Decimal `gorm:"type:decimal(20, 8);24hLucky值" json:"luckyValue24H"`
	WeightedBlocks                int             `gorm:"type:int" json:"weightedBlocks,omitempty"`
	RewardValue                   decimal.Decimal `gorm:"type:decimal(20, 8)" json:"rewardValue"`
	LastWeightedBlocks            int             `gorm:"type:int" json:"lastWeightedBlocks,omitempty"`
	LastRewardValue               decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastRewardValue"`
	LastMonthWeightedBlocks       int             `gorm:"type:int" json:"lastMonthWeightedBlocks,omitempty"`
	LastMonthRewardValue          decimal.Decimal `gorm:"type:decimal(20, 8)" json:"lastMonthRewardValue"`
	RealWeightedBlocks24h         decimal.Decimal `gorm:"type:decimal(20, 8)" json:"realWeightedBlocks24h,omitempty"` //算上比例的实际24h报块数量
	RealRewardValue24h            decimal.Decimal `gorm:"type:decimal(20, 8);24h出块奖励金额" json:"realRewardValue24h"`    //算上比例的实际24h出块奖励金额
	RealWeightedBlocksMonth       decimal.Decimal `gorm:"type:decimal(20, 8);" json:"realWeightedBlocksMonth"`        //算上比例的实际月报块数量
	RealRewardValueMonth          decimal.Decimal `gorm:"type:decimal(20, 8);" json:"realRewardValueMonth"`           //算上比例的实际月出块奖励金额
	ReceiveAmount                 decimal.Decimal `gorm:"type:decimal(20, 8);节点接收数量" json:"receiveAmount"`
	BurnAmount                    decimal.Decimal `gorm:"type:decimal(20, 8);节点销毁数量" json:"burnAmount"`
	SendAmount                    decimal.Decimal `gorm:"type:decimal(20, 8);节点发送数量" json:"sendAmount"`
	LastReceiveAmount             decimal.Decimal `gorm:"type:decimal(20, 8);前一天接收数量" json:"lastReceiveAmount"`
	LastBurnAmount                decimal.Decimal `gorm:"type:decimal(20, 4);前一天销毁数量" json:"lastBurnAmount"`
	LastSendAmount                decimal.Decimal `gorm:"type:decimal(20, 4);前一天提现数量" json:"lastSendAmount"`
	LastMonthReceiveAmount        decimal.Decimal `gorm:"type:decimal(20, 8);上月末节点接收数量" json:"lastMonthReceiveAmount"`
	LastMonthBurnAmount           decimal.Decimal `gorm:"type:decimal(20, 4);上月末销毁数量" json:"lastMonthBurnAmount"`
	LastMonthSendAmount           decimal.Decimal `gorm:"type:decimal(20, 4);上月末提现数量" json:"lastMonthSendAmount"`
	ControlBalance                decimal.Decimal `gorm:"type:decimal(20,8)" json:"controlBalance"`
	SectorEffective               int             `gorm:"type:int" json:"sectorEffective"`
	SectorError                   int             `gorm:"type:int" json:"sectorError"`
	SectorSize                    string          `json:"sectorSize" gorm:"type:varchar(50);comment:扇区大小"`
	CreateTime                    time.Time       `json:"createTime" gorm:"type:datetime;comment:账户创建时间"`
	EndTime                       time.Time       `json:"endTime" gorm:"type:datetime;comment:节点结束时间"`
}
