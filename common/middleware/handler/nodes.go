package handler

import (
	models2 "fil-admin/app/filpool/models"
	"fil-admin/utils"
	"math"
	"time"

	"github.com/shopspring/decimal"

	"fil-admin/common/models"
)

type FilNodes struct {
	models.Model
	Node                    string          `json:"node" gorm:"type:varchar(255);comment:账户名称"`
	MsgCount                int             `json:"msgCount" gorm:"type:bigint;comment:消息数量"`
	SectorType              string          `json:"sectorType" gorm:"type:varchar(50);comment:扇区类型"`
	CreateTime              time.Time       `json:"createTime" gorm:"type:datetime;comment:账户创建时间"`
	AvailableBalance        decimal.Decimal `json:"availableBalance" gorm:"type:decimal(20,8);comment:可用余额"`
	Balance                 decimal.Decimal `json:"balance" gorm:"type:decimal(20,8);comment:账户余额"`
	SectorPledgeBalance     decimal.Decimal `json:"sectorPledgeBalance" gorm:"type:decimal(20,8);comment:扇区质押"`
	VestingFunds            decimal.Decimal `json:"vestingFunds" gorm:"type:decimal(20,8);comment:存储锁仓"`
	RewardValue             decimal.Decimal `json:"rewardValue" gorm:"type:decimal(20,8);comment:总奖励"`
	WeightedBlocks          int             `gorm:"type:int"`
	QualityAdjPower         decimal.Decimal `json:"qualityAdjPower" gorm:"type:decimal(20,4);comment:有效算力"`
	QualityAdjPowerDelta24h decimal.Decimal `json:"qualityAdjPowerDelta24h"`
	PowerUnit               string          `json:"powerUnit" gorm:"type:varchar(50);comment:算力单位"`
	PowerPoint              decimal.Decimal `json:"powerPoint" gorm:"type:decimal(10,3);comment:PowerPoint"`
	PowerGrade              string          `json:"powerGrade" gorm:"type:varchar(50);comment:PowerGrade"`
	SectorSize              string          `json:"sectorSize" gorm:"type:varchar(50);comment:扇区大小"`
	SectorStatus            string          `json:"sectorStatus" gorm:"type:varchar(255);comment:SectorStatus"`
	SectorTotal             int             `gorm:"type:int" json:"sectorTotal"`
	SectorEffective         int             `gorm:"type:int" json:"sectorEffective"`
	SectorError             int             `gorm:"type:int" json:"sectorError"`
	SectorRecovering        int             `gorm:"type:int" json:"sectorRecovering"`
	Status                  string          `json:"status" gorm:"type:int;comment:节点状态默认 1可用   0下架"`
	Type                    string          `json:"type" gorm:"type:int;comment:节点类型 联合类型"`
	EndTime                 time.Time       `json:"endTime" gorm:"type:datetime;comment:节点结束时间"`
	DeptId                  int             `json:"deptId" gorm:"type:int;comment:部门ID"`
	Title                   string          `json:"title" gorm:"type:varchar(255);comment:节点标签"`
	ChartList               *[]NodesChart   `json:"chartList" gorm:"-"`
	MiningEfficiency        decimal.Decimal `json:"miningEfficiency" gorm:"type:decimal(20,8)"`
	Height                  int             `json:"height" gorm:"type:int;comment:Height"`
	SyncStatus              bool            `json:"syncStatus" gorm:"type:int;comment:同步状态"`
	PowerDeltaShow          string          `json:"powerDeltaShow" gorm:"-"`
	PowerDeltaUnit          string          `json:"powerDeltaUnit" gorm:"-"`
}

func (FilNodes) TableName() string {
	return "fil_nodes"
}

func (s *FilNodes) Generate(node models2.FilNodes) FilNodes {
	v, str := utils.DecimalPowerValue(node.QualityAdjPowerDelta24h.String())
	v1, str1 := utils.DecimalPowerValue(node.QualityAdjPower.Mul(decimal.NewFromFloat(math.Pow10(6))).String())
	return FilNodes{
		Model:                   models.Model{Id: node.Id},
		Node:                    node.Node,
		MsgCount:                node.MsgCount,
		SectorType:              node.SectorType,
		CreateTime:              node.CreateTime,
		AvailableBalance:        node.AvailableBalance.RoundDown(2),
		Balance:                 node.Balance.RoundDown(2),
		SectorPledgeBalance:     node.SectorPledgeBalance.RoundDown(2),
		VestingFunds:            node.VestingFunds.RoundDown(2),
		RewardValue:             node.RewardValue.RoundDown(2),
		WeightedBlocks:          node.WeightedBlocks,
		QualityAdjPower:         v1,
		PowerUnit:               str1,
		PowerPoint:              node.PowerPoint,
		PowerGrade:              node.PowerGrade,
		SectorSize:              node.SectorSize,
		SectorStatus:            node.SectorStatus,
		SectorTotal:             node.SectorTotal,
		SectorEffective:         node.SectorEffective,
		SectorError:             node.SectorError,
		SectorRecovering:        node.SectorRecovering,
		QualityAdjPowerDelta24h: v,
		PowerDeltaUnit:          str,
		PowerDeltaShow:          node.GetPowerDeltaShow(),
		Status:                  node.Status,
		Type:                    node.Type,
		EndTime:                 node.EndTime,
		Title:                   node.Title,
		MiningEfficiency:        node.MiningEfficiency.Mul(decimal.NewFromInt(1000)).RoundDown(1),
		Height:                  node.Height,
		SyncStatus:              node.OnLine,
	}
}

type MinerSectors struct {
	Miner           string    `json:"miner"`
	SectorSize      string    `json:"sector_size"`
	SectorStatus    string    `json:"sector_status"`
	SectorEffective int       `json:"sector_effective"`
	Sectors         []Sectors `json:"sectors"`
}

type Sectors struct {
	Day       string `json:"day"`
	From      int    `json:"from"`
	To        int    `json:"to"`
	SectorNum int    `json:"sectorNum"`
	FromTo    string `json:"fromTo"`
	Power     string `json:"power"`
}
