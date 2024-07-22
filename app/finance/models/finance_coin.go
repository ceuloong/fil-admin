package models

import (
	"github.com/shopspring/decimal"
	"time"

	"fil-admin/common/models"
)

type FinanceCoin struct {
	models.Model

	Name        string          `json:"name" gorm:"type:varchar(128);comment:币种名称"`
	CoinAmount  decimal.Decimal `json:"coinAmount" gorm:"type:decimal(20,8);comment:币种数量"`
	Address     string          `json:"address" gorm:"type:varchar(255);comment:收币地址"`
	Rate        decimal.Decimal `json:"rate" gorm:"type:decimal(20,4);comment:汇率"`
	CashAmount  decimal.Decimal `json:"cashAmount" gorm:"type:decimal(20,4);comment:现金数量"`
	Status      string          `json:"status" gorm:"type:int;comment:状态"`
	PublishAt   time.Time       `json:"publishAt" gorm:"type:timestamp;comment:发布时间"`
	Description string          `json:"description" gorm:"type:varchar(255);comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (FinanceCoin) TableName() string {
	return "finance_coin"
}

func (e *FinanceCoin) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FinanceCoin) GetId() interface{} {
	return e.Id
}
