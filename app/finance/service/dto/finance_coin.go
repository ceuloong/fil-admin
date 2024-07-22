package dto

import (
	"fil-admin/app/finance/models"
	"github.com/shopspring/decimal"
	"time"

	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type FinanceCoinGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string          `form:"name"  search:"type:exact;column:name;table:finance_coin" comment:"币种名称"`
	Address        string          `form:"address"  search:"type:exact;column:address;table:finance_coin" comment:"收币地址"`
	Rate           decimal.Decimal `form:"rate"  search:"type:exact;column:rate;table:finance_coin" comment:"汇率"`
	Status         string          `form:"status"  search:"type:exact;column:status;table:finance_coin" comment:"状态"`
	FinanceCoinOrder
}

type FinanceCoinOrder struct {
	Id         string          `form:"idOrder"  search:"type:order;column:id;table:finance_coin"`
	Name       string          `form:"nameOrder"  search:"type:order;column:name;table:finance_coin"`
	CoinAmount decimal.Decimal `form:"coinAmountOrder"  search:"type:order;column:coin_amount;table:finance_coin"`
	Address    string          `form:"addressOrder"  search:"type:order;column:address;table:finance_coin"`
	Rate       decimal.Decimal `form:"rateOrder"  search:"type:order;column:rate;table:finance_coin"`
	CashAmount decimal.Decimal `form:"cashAmountOrder"  search:"type:order;column:cash_amount;table:finance_coin"`
	Status     string          `form:"statusOrder"  search:"type:order;column:status;table:finance_coin"`
	PublishAt  string          `form:"publishAtOrder"  search:"type:order;column:publish_at;table:finance_coin"`
	CreatedAt  string          `form:"createdAtOrder"  search:"type:order;column:created_at;table:finance_coin"`
	UpdatedAt  string          `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:finance_coin"`
	DeletedAt  string          `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:finance_coin"`
	CreateBy   string          `form:"createByOrder"  search:"type:order;column:create_by;table:finance_coin"`
	UpdateBy   string          `form:"updateByOrder"  search:"type:order;column:update_by;table:finance_coin"`
}

func (m *FinanceCoinGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FinanceCoinInsertReq struct {
	Id          int             `json:"-" comment:"编码"` // 编码
	Name        string          `json:"name" comment:"币种名称"`
	CoinAmount  decimal.Decimal `json:"coinAmount" comment:"币种数量"`
	Address     string          `json:"address" comment:"收币地址"`
	Rate        decimal.Decimal `json:"rate" comment:"汇率"`
	CashAmount  decimal.Decimal `json:"cashAmount" comment:"现金数量"`
	Status      string          `json:"status" comment:"状态"`
	PublishAt   time.Time       `json:"publishAt" comment:"交易时间"`
	Description string          `json:"description" comment:"备注"`
	common.ControlBy
}

func (s *FinanceCoinInsertReq) Generate(model *models.FinanceCoin) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CoinAmount = s.CoinAmount
	model.Address = s.Address
	model.Rate = s.Rate
	model.CashAmount = s.CashAmount
	model.Status = s.Status
	model.PublishAt = s.PublishAt
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Description = s.Description
}

func (s *FinanceCoinInsertReq) GetId() interface{} {
	return s.Id
}

type FinanceCoinUpdateReq struct {
	Id          int             `uri:"id" comment:"编码"` // 编码
	Name        string          `json:"name" comment:"币种名称"`
	CoinAmount  decimal.Decimal `json:"coinAmount" comment:"币种数量"`
	Address     string          `json:"address" comment:"收币地址"`
	Rate        decimal.Decimal `json:"rate" comment:"汇率"`
	CashAmount  decimal.Decimal `json:"cashAmount" comment:"现金数量"`
	Status      string          `json:"status" comment:"状态"`
	PublishAt   time.Time       `json:"publishAt" comment:"发布时间"`
	Description string          `json:"description" comment:"备注"`
	common.ControlBy
}

func (s *FinanceCoinUpdateReq) Generate(model *models.FinanceCoin) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.CoinAmount = s.CoinAmount
	model.Address = s.Address
	model.Rate = s.Rate
	model.CashAmount = s.CashAmount
	model.Status = s.Status
	model.PublishAt = s.PublishAt
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Description = s.Description
}

func (s *FinanceCoinUpdateReq) GetId() interface{} {
	return s.Id
}

// FinanceCoinGetReq 功能获取请求参数
type FinanceCoinGetReq struct {
	Id int `uri:"id"`
}

func (s *FinanceCoinGetReq) GetId() interface{} {
	return s.Id
}

// FinanceCoinDeleteReq 功能删除请求参数
type FinanceCoinDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FinanceCoinDeleteReq) GetId() interface{} {
	return s.Ids
}
