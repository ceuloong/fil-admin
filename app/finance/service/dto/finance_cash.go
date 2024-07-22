package dto

import (
	"fil-admin/app/finance/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
	"github.com/shopspring/decimal"
	"time"
)

type FinanceCashGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:exact;column:name;table:finance_cash" comment:"货币名称"`
	Amount         string `form:"amount"  search:"type:exact;column:amount;table:finance_cash" comment:"金额"`
	Balance        string `form:"balance"  search:"type:exact;column:balance;table:finance_cash" comment:"余额"`
	Memo           string `form:"memo"  search:"type:exact;column:memo;table:finance_cash" comment:"备注"`
	DictId         string `form:"dictId"  search:"type:exact;column:dict_id;table:finance_cash" comment:"货币类型"`
	Employee       string `form:"employee"  search:"type:exact;column:employee;table:finance_cash" comment:"员工名称"`
	Status         string `form:"status"  search:"type:exact;column:status;table:finance_cash" comment:"状态"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:created_at;table:finance_cash" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:created_at;table:finance_cash" comment:"更新时间"`
	TypeJoin       `search:"type:left;on:type_id:type_id;table:finance_cash;join:finance_type"`
	FinanceCashOrder
}

type FinanceCashOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:finance_cash"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:finance_cash"`
	Amount    string `form:"amountOrder"  search:"type:order;column:amount;table:finance_cash"`
	Memo      string `form:"memoOrder"  search:"type:order;column:memo;table:finance_cash"`
	Employee  string `form:"employeeOrder"  search:"type:order;column:employee;table:finance_cash"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:finance_cash"`
	PublishAt string `form:"publishAtOrder"  search:"type:order;column:publish_at;table:finance_cash"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:finance_cash"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:finance_cash"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:finance_cash"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:finance_cash"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:finance_cash"`
}

type TypeJoin struct {
	TypeId string `search:"type:contains;column:type_path;table:finance_type" form:"typeId"`
}

func (m *FinanceCashGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FinanceCashInsertReq struct {
	Id        int             `json:"-" comment:"编码"` // 编码
	Name      string          `json:"name" comment:"货币名称"`
	Amount    decimal.Decimal `json:"amount" comment:"金额"`
	Balance   decimal.Decimal `json:"balance" comment:"余额"`
	TypeId    int             `json:"typeId" comment:"收支类型"`
	DictId    string          `json:"dictId" comment:"账户类型"`
	Memo      string          `json:"memo" comment:"备注"`
	Employee  string          `json:"employee" comment:"员工名称"`
	Status    string          `json:"status" comment:"状态"`
	PublishAt time.Time       `json:"publishAt" comment:"交易时间"`
	common.ControlBy
}

func (s *FinanceCashInsertReq) Generate(model *models.FinanceCash) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Amount = s.Amount
	model.Balance = s.Balance
	model.TypeId = s.TypeId
	model.DictId = s.DictId
	model.Memo = s.Memo
	model.Employee = s.Employee
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.PublishAt = s.PublishAt
}

func (s *FinanceCashInsertReq) GetId() interface{} {
	return s.Id
}

type FinanceCashUpdateReq struct {
	Id        int             `uri:"id" comment:"编码"` // 编码
	Name      string          `json:"name" comment:"货币名称"`
	Amount    decimal.Decimal `json:"amount" comment:"金额"`
	TypeId    int             `json:"typeId" comment:"收支类型"`
	DictId    string          `json:"dictId" comment:"账户类型"`
	Memo      string          `json:"memo" comment:"备注"`
	Employee  string          `json:"employee" comment:"员工名称"`
	Status    string          `json:"status" comment:"状态"`
	PublishAt time.Time       `json:"publishAt" comment:"交易时间"`
	common.ControlBy
}

func (s *FinanceCashUpdateReq) Generate(model *models.FinanceCash) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Amount = s.Amount
	model.TypeId = s.TypeId
	model.DictId = s.DictId
	model.Memo = s.Memo
	model.Employee = s.Employee
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.PublishAt = s.PublishAt
}

func (s *FinanceCashUpdateReq) GetId() interface{} {
	return s.Id
}

// FinanceCashGetReq 功能获取请求参数
type FinanceCashGetReq struct {
	Id int `uri:"id"`
}

func (s *FinanceCashGetReq) GetId() interface{} {
	return s.Id
}

// FinanceCashDeleteReq 功能删除请求参数
type FinanceCashDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FinanceCashDeleteReq) GetId() interface{} {
	return s.Ids
}
