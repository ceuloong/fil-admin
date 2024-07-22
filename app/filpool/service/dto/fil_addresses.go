package dto

import (
	"fil-admin/app/filpool/models"
	"github.com/shopspring/decimal"
	"time"

	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type FilAddressesGetPageReq struct {
	dto.Pagination `search:"-"`
	Node           string `form:"node"  search:"type:exact;column:node;table:fil_addresses" comment:""`
	AccountId      string `form:"accountId"  search:"type:exact;column:account_id;table:fil_addresses" comment:""`
	Address        string `form:"address"  search:"type:exact;column:address;table:fil_addresses" comment:""`
	Type           string `form:"type"  search:"type:exact;column:type;table:fil_addresses" comment:"controller, worker, other"`
	Status         string `form:"status"  search:"type:exact;column:status;table:fil_addresses" comment:"状态"`
	FilAddressesOrder
}

type FilAddressesOrder struct {
	Id               string `form:"idOrder"  search:"type:order;column:id;table:fil_addresses"`
	Node             string `form:"nodeOrder"  search:"type:order;column:node;table:fil_addresses"`
	AccountId        string `form:"accountIdOrder"  search:"type:order;column:account_id;table:fil_addresses"`
	Address          string `form:"addressOrder"  search:"type:order;column:address;table:fil_addresses"`
	Balance          string `form:"balanceOrder"  search:"type:order;column:balance;table:fil_addresses"`
	Message          string `form:"messageOrder"  search:"type:order;column:message;table:fil_addresses"`
	Type             string `form:"typeOrder"  search:"type:order;column:type;table:fil_addresses"`
	CreateTime       string `form:"createTimeOrder"  search:"type:order;column:create_time;table:fil_addresses"`
	CreatedTime      string `form:"createdTimeOrder"  search:"type:order;column:created_time;table:fil_addresses"`
	AccountType      string `form:"accountTypeOrder"  search:"type:order;column:account_type;table:fil_addresses"`
	LastTransferTime string `form:"lastTransferTimeOrder"  search:"type:order;column:last_transfer_time;table:fil_addresses"`
	Nonce            string `form:"nonceOrder"  search:"type:order;column:nonce;table:fil_addresses"`
	ReceiveAmount    string `form:"receiveAmountOrder"  search:"type:order;column:receive_amount;table:fil_addresses"`
	BurnAmount       string `form:"burnAmountOrder"  search:"type:order;column:burn_amount;table:fil_addresses"`
	SendAmount       string `form:"sendAmountOrder"  search:"type:order;column:send_amount;table:fil_addresses"`
	TransferCount    string `form:"transferCountOrder"  search:"type:order;column:transfer_count;table:fil_addresses"`
	TimeTag          string `form:"timeTagOrder"  search:"type:order;column:time_tag;table:fil_addresses"`
	Status           string `form:"statusOrder"  search:"type:order;column:status;table:fil_addresses"`
}

func (m *FilAddressesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FilAddressesInsertReq struct {
	Id               int             `json:"-" comment:""` //
	Node             string          `json:"node" comment:""`
	AccountId        string          `json:"accountId" comment:""`
	Address          string          `json:"address" comment:""`
	Balance          decimal.Decimal `json:"balance" comment:""`
	Message          string          `json:"message" comment:""`
	Type             string          `json:"type" comment:"controller, worker, other"`
	CreateTime       time.Time       `json:"createTime" comment:"地址创建时间"`
	CreatedTime      time.Time       `json:"createdTime" comment:"记录添加时间"`
	AccountType      string          `json:"accountType" comment:""`
	LastTransferTime time.Time       `json:"lastTransferTime" comment:"最后交易时间"`
	Nonce            int64           `json:"nonce" comment:""`
	ReceiveAmount    decimal.Decimal `json:"receiveAmount" comment:"总的接收数量"`
	BurnAmount       decimal.Decimal `json:"burnAmount" comment:"总的销毁数量"`
	SendAmount       decimal.Decimal `json:"sendAmount" comment:"总的发送数量"`
	TransferCount    int64           `json:"transferCount" comment:"转账交易数量"`
	TimeTag          int64           `json:"timeTag" comment:"时间标签"`
	Status           int             `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *FilAddressesInsertReq) Generate(model *models.FilAddresses) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Node = s.Node
	model.AccountId = s.AccountId
	model.Address = s.Address
	model.Balance = s.Balance
	model.Message = s.Message
	model.Type = s.Type
	model.CreateTime = s.CreateTime
	model.CreatedTime = s.CreatedTime
	model.AccountType = s.AccountType
	model.LastTransferTime = s.LastTransferTime
	model.Nonce = s.Nonce
	model.ReceiveAmount = s.ReceiveAmount
	model.BurnAmount = s.BurnAmount
	model.SendAmount = s.SendAmount
	model.TransferCount = s.TransferCount
	model.TimeTag = s.TimeTag
	model.Status = s.Status
}

func (s *FilAddressesInsertReq) GetId() interface{} {
	return s.Id
}

type FilAddressesUpdateReq struct {
	Id               int             `uri:"id" comment:""` //
	Node             string          `json:"node" comment:""`
	AccountId        string          `json:"accountId" comment:""`
	Address          string          `json:"address" comment:""`
	Balance          decimal.Decimal `json:"balance" comment:""`
	Message          string          `json:"message" comment:""`
	Type             string          `json:"type" comment:"controller, worker, other"`
	CreateTime       time.Time       `json:"createTime" comment:"地址创建时间"`
	CreatedTime      time.Time       `json:"createdTime" comment:"记录添加时间"`
	AccountType      string          `json:"accountType" comment:""`
	LastTransferTime time.Time       `json:"lastTransferTime" comment:"最后交易时间"`
	Nonce            int64           `json:"nonce" comment:""`
	ReceiveAmount    decimal.Decimal `json:"receiveAmount" comment:"总的接收数量"`
	BurnAmount       decimal.Decimal `json:"burnAmount" comment:"总的销毁数量"`
	SendAmount       decimal.Decimal `json:"sendAmount" comment:"总的发送数量"`
	TransferCount    int64           `json:"transferCount" comment:"转账交易数量"`
	TimeTag          int64           `json:"timeTag" comment:"时间标签"`
	Status           int             `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *FilAddressesUpdateReq) Generate(model *models.FilAddresses) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Node = s.Node
	model.AccountId = s.AccountId
	model.Address = s.Address
	model.Balance = s.Balance
	model.Message = s.Message
	model.Type = s.Type
	model.CreateTime = s.CreateTime
	model.CreatedTime = s.CreatedTime
	model.AccountType = s.AccountType
	model.LastTransferTime = s.LastTransferTime
	model.Nonce = s.Nonce
	model.ReceiveAmount = s.ReceiveAmount
	model.BurnAmount = s.BurnAmount
	model.SendAmount = s.SendAmount
	model.TransferCount = s.TransferCount
	model.TimeTag = s.TimeTag
	model.Status = s.Status
}

func (s *FilAddressesUpdateReq) GetId() interface{} {
	return s.Id
}

// FilAddressesGetReq 功能获取请求参数
type FilAddressesGetReq struct {
	Id int `uri:"id"`
}

func (s *FilAddressesGetReq) GetId() interface{} {
	return s.Id
}

// FilAddressesDeleteReq 功能删除请求参数
type FilAddressesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FilAddressesDeleteReq) GetId() interface{} {
	return s.Ids
}
