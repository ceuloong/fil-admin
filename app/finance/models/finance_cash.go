package models

import (
	"fil-admin/common/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type FinanceCash struct {
	models.Model

	Name      string          `json:"name" gorm:"type:varchar(128);comment:货币名称"`
	Amount    decimal.Decimal `json:"amount" gorm:"type:decimal(20,4);comment:金额"`
	Balance   decimal.Decimal `json:"balance" gorm:"type:decimal(20,4);comment:余额"`
	TypeId    int             `json:"typeId" gorm:"type:int;comment:收支类型"`
	DictId    string          `json:"dictId" gorm:"type:int;comment:是账户还是现金"`
	Memo      string          `json:"memo" gorm:"type:varchar(255);comment:备注"`
	Employee  string          `json:"employee" gorm:"type:varchar(255);comment:员工名称"`
	Status    string          `json:"status" gorm:"type:int;comment:状态"`
	PublishAt time.Time       `json:"publishAt" gorm:"type:timestamp;comment:交易时间"`
	TypeIds   []int           `json:"typeIds" gorm:"-"`
	Type      *FinanceType    `json:"type"`
	models.ModelTime
	models.ControlBy
}

func (FinanceCash) TableName() string {
	return "finance_cash"
}

func (e *FinanceCash) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FinanceCash) GetId() interface{} {
	return e.Id
}

func AfterFind(e *FinanceCash, _ *gorm.DB) error {
	e.TypeIds = []int{e.TypeId}
	return nil
}
