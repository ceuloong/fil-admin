package models

import (
	admin "fil-admin/cmd/migrate/migration/models"
	"log"
	"time"

	"github.com/shopspring/decimal"

	"fil-admin/common/models"
)

type FilNodes struct {
	models.Model

	Node                       string          `json:"node" gorm:"type:varchar(255);comment:账户名称"`
	MsigNode                   string          `json:"msigNode" gorm:"type:varchar(255);comment:所属地址"`
	Address                    string          `json:"address" gorm:"type:varchar(255);comment:地址"`
	MsgCount                   int             `json:"msgCount" gorm:"type:bigint;comment:消息数量"`
	SectorType                 string          `json:"sectorType" gorm:"type:varchar(50);comment:扇区类型"`
	CreateTime                 time.Time       `json:"createTime" gorm:"type:datetime;comment:账户创建时间"`
	AvailableBalance           decimal.Decimal `json:"availableBalance" gorm:"type:decimal(20,8);comment:可用余额"`
	Balance                    decimal.Decimal `json:"balance" gorm:"type:decimal(20,8);comment:账户余额"`
	SectorPledgeBalance        decimal.Decimal `json:"sectorPledgeBalance" gorm:"type:decimal(20,8);comment:扇区质押"`
	VestingFunds               decimal.Decimal `json:"vestingFunds" gorm:"type:decimal(20,8);comment:存储锁仓"`
	Height                     int             `json:"height" gorm:"type:bigint;comment:Height"`
	LastTime                   time.Time       `json:"lastTime" gorm:"type:datetime;comment:LastTime"`
	RewardValue                decimal.Decimal `json:"rewardValue" gorm:"type:decimal(20,8);comment:总奖励"`
	WeightedBlocks             int             `gorm:"type:int"`
	QualityAdjPower            decimal.Decimal `json:"qualityAdjPower" gorm:"type:decimal(20,4);comment:有效算力"`
	PowerUnit                  string          `json:"powerUnit" gorm:"type:varchar(50);comment:算力单位"`
	PowerPoint                 decimal.Decimal `json:"powerPoint" gorm:"type:decimal(10,3);comment:PowerPoint"`
	PowerGrade                 string          `json:"powerGrade" gorm:"type:varchar(50);comment:PowerGrade"`
	SectorSize                 string          `json:"sectorSize" gorm:"type:varchar(50);comment:扇区大小"`
	SectorStatus               string          `json:"sectorStatus" gorm:"type:varchar(255);comment:SectorStatus"`
	SectorTotal                int             `gorm:"type:int"`
	SectorEffective            int             `gorm:"type:int"`
	SectorError                int             `gorm:"type:int"`
	SectorRecovering           int             `gorm:"type:int"`
	ControlAddress             string          `json:"controlAddress" gorm:"type:varchar(255);comment:ControlAddress"`
	ControlBalance             decimal.Decimal `json:"controlBalance" gorm:"type:decimal(20,8);comment:ControlBalance"`
	Status                     string          `json:"status" gorm:"type:int;comment:节点状态默认 1可用   0下架"`
	Type                       string          `json:"type" gorm:"type:int;comment:节点类型 联合类型"`
	DistributePoint            decimal.Decimal `json:"distributePoint" gorm:"type:decimal(10,3);comment:联合挖矿分配比例"`
	HasTransfer                decimal.Decimal `json:"hasTransfer" gorm:"type:decimal(20,8);comment:上次分成后转出数量"`
	HasRealDistribute          decimal.Decimal `json:"hasRealDistribute" gorm:"type:decimal(20,8);comment:实际已分配的数量"`
	LastDisSectorPledgeBalance decimal.Decimal `json:"lastDisSectorPledgeBalance" gorm:"type:decimal(20,8);comment:记录上一次分币时的质押数量"`
	ReceiveAmount              decimal.Decimal `json:"receiveAmount" gorm:"type:decimal(20,8);comment:接收数量"`
	BurnAmount                 decimal.Decimal `json:"burnAmount" gorm:"type:decimal(20,8);comment:销毁数量"`
	SendAmount                 decimal.Decimal `json:"sendAmount" gorm:"type:decimal(20,8);comment:发送数量"`
	LastDistributeTime         time.Time       `json:"lastDistributeTime" gorm:"type:datetime;comment:最后一次分币时间"`
	EndTime                    time.Time       `json:"endTime" gorm:"type:datetime;comment:节点结束时间"`
	Tag                        string          `json:"tag" gorm:"-"`
	TimeTag                    int64           `gorm:"type:bigint" 时间标签`
	BlocksMined24h             int             `json:"blocksMined24h"`
	TotalRewards24h            decimal.Decimal `json:"totalRewards24h"`
	LuckyValue24h              decimal.Decimal `json:"luckyValue24h"`
	QualityAdjPowerDelta24h    decimal.Decimal `json:"qualityAdjPowerDelta24h"`
	BlocksMined7d              string          `json:"blocksMinedH7d"`
	TotalRewards7d             decimal.Decimal `json:"totalRewards7d"`
	LuckyValue7d               decimal.Decimal `json:"luckyValue7d"`
	QualityAdjPowerDelta7d     decimal.Decimal `json:"qualityAdjPowerDelta7d"`
	BlocksMined30d             string          `json:"blocksMinedH30d"`
	TotalRewards30d            decimal.Decimal `json:"totalRewards30d"`
	LuckyValue30d              decimal.Decimal `json:"luckyValue30d"`
	QualityAdjPowerDelta30d    decimal.Decimal `json:"qualityAdjPowerDelta30d"`
	MiningEfficiency           decimal.Decimal `gorm:"type:decimal(20,8)"` // 挖矿效率
	MiningEfficiency7d         decimal.Decimal `gorm:"type:decimal(20,8)"`
	MiningEfficiency30d        decimal.Decimal `gorm:"type:decimal(20,8)"`
	NodeChart                  *NodesChart     `json:"nodeChart" gorm:"foreignKey:node;references:node;"`
	DeptId                     int             `json:"deptId" gorm:"type:int;comment:部门ID"`
	Title                      string          `json:"title" gorm:"type:varchar(255);comment:节点标签"`
	SyncStatus                 string          `json:"syncStatus" gorm:"type:varchar(50);comment:同步状态"`
	Dept                       *admin.SysDept  `json:"dept"`
	ChartList                  *[]NodesChart   `json:"chartList" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (FilNodes) TableName() string {
	return "fil_nodes"
}

func (e *FilNodes) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FilNodes) GetId() interface{} {
	return e.Id
}

func (e *FilNodes) GetRewardAvailable() interface{} {
	log.Printf("调用了吗。。。。。。。。。")
	return e.RewardValue.Sub(e.VestingFunds)
}

//func (e *FilNodes) AfterFind(_ *gorm.DB) error {
//	f, _ := e.NodeChart.QualityAdjPower.Float64()
//	e.NodeChartData = append(e.NodeChartData, &NodeChartData{
//		X: e.NodeChart.LastTime.Format("1-02 15"),
//		Y: f,
//	})
//	return nil
//}
