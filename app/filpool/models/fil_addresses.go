package models

import (
	"github.com/shopspring/decimal"
	"time"

	"fil-admin/common/models"
)

type FilAddresses struct {
	models.Model

	Node             string          `json:"node" gorm:"type:varchar(255);comment:Node"`
	AccountId        string          `json:"accountId" gorm:"type:varchar(50);comment:AccountId"`
	Address          string          `json:"address" gorm:"type:varchar(255);comment:Address"`
	Balance          decimal.Decimal `json:"balance" gorm:"decimal(20,8);comment:Balance"`
	Message          string          `json:"message" gorm:"type:varchar(255);comment:Message"`
	Type             string          `json:"type" gorm:"type:varchar(50);comment:controller, worker, other"`
	CreateTime       time.Time       `json:"createTime" gorm:"type:datetime;comment:地址创建时间"`
	CreatedTime      time.Time       `json:"createdTime" gorm:"type:datetime;comment:记录添加时间"`
	AccountType      string          `json:"accountType" gorm:"type:varchar(50);comment:AccountType"`
	LastTransferTime time.Time       `json:"lastTransferTime" gorm:"type:datetime;comment:最后交易时间"`
	Nonce            int64           `json:"nonce" gorm:"type:int;comment:Nonce"`
	ReceiveAmount    decimal.Decimal `json:"receiveAmount" gorm:"type:decimal(20,8);comment:总的接收数量"`
	BurnAmount       decimal.Decimal `json:"burnAmount" gorm:"type:decimal(20,8);comment:总的销毁数量"`
	SendAmount       decimal.Decimal `json:"sendAmount" gorm:"type:decimal(20,8);comment:总的发送数量"`
	TransferCount    int64           `json:"transferCount" gorm:"type:int;comment:转账交易数量"`
	TimeTag          int64           `json:"timeTag" gorm:"type:bigint;comment:时间标签"`
	Status           int             `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (FilAddresses) TableName() string {
	return "fil_addresses"
}

func (e *FilAddresses) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FilAddresses) GetId() interface{} {
	return e.Id
}
