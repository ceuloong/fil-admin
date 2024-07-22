package models

import (
	"fil-admin/common/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type FilDistribution struct {
	models.Model
	NodeId           int             `json:"nodeId" gorm:"type:int(11);comment:NodeId"`
	Node             string          `json:"node" gorm:"type:varchar(128);comment:节点名称"`
	AvailableBalance decimal.Decimal `json:"availableBalance" gorm:"type:decimal(20,8);comment:可用余额"`
	HasTransfer      decimal.Decimal `json:"hasTransfer" gorm:"type:decimal(20,8);comment:本期分币前已转出数量"`
	DistributePoint  decimal.Decimal `json:"distributePoint" gorm:"type:decimal(20,3);comment:分成比例"`
	LastSectorPledge decimal.Decimal `json:"lastSectorPledge" gorm:"type:decimal(20,8);comment:上期质押数量"`
	CurSectorPledge  decimal.Decimal `json:"curSectorPledge" gorm:"type:decimal(20,8);comment:当前质押数量"`
	EffectAmount     decimal.Decimal `json:"effectAmount" gorm:"type:decimal(20,8);comment:参与分币数量=可用余额+转出数量"`
	DistributeAmount decimal.Decimal `json:"distributeAmount" gorm:"type:decimal(20,8);comment:应该分币数量=参与数量x分成比例"`
	AddressFrom      string          `json:"addressFrom" gorm:"type:varchar(255);comment:发送地址"`
	AddressTo        string          `json:"addressTo" gorm:"type:varchar(255);comment:接收地址"`
	Status           int             `json:"status" gorm:"type:int;comment:分币状态"`
	NodeIds          []int           `json:"nodeIds" gorm:"-"`
	FilNode          *FilNodes       `json:"filNode" gorm:"foreignKey:NodeId;references:Id;"`
	models.ModelTime
	models.ControlBy
}

func (FilDistribution) TableName() string {
	return "fil_distribution"
}

func (e *FilDistribution) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FilDistribution) GetId() interface{} {
	return e.Id
}

func (e *FilDistribution) AfterFind(_ *gorm.DB) error {
	e.NodeIds = []int{e.NodeId}
	return nil
}
