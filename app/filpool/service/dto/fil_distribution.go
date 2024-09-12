package dto

import (
	"fil-admin/app/filpool/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
	"time"

	"github.com/shopspring/decimal"
)

type FilDistributionGetPageReq struct {
	dto.Pagination `search:"-"`
	Node           string `form:"node"  search:"type:exact;column:node;table:fil_distribution" comment:"节点名称"`
	AddressFrom    string `form:"addressFrom"  search:"type:contains;column:address_from;table:fil_distribution" comment:"发送地址"`
	AddressTo      string `form:"addressTo"  search:"type:contains;column:address_to;table:fil_distribution" comment:"接收地址"`
	Status         int    `form:"status"  search:"type:exact;column:status;table:fil_distribution" comment:"分币状态"`
	NodeIds        []int  `form:"nodeIds"  search:"type:in;column:node_id;table:fil_distribution" comment:"节点ID"`
	FilNodeJoin    `search:"type:left;on:id:node_id;table:fil_distribution;join:fil_nodes"`
	FilDistributionOrder
}

type FilNodeJoin struct {
	Type string `search:"type:exact;column:type;table:fil_nodes" form:"type"`
}

type FilDistributionOrder struct {
	Id               string `form:"idOrder"  search:"type:order;column:id;table:fil_distribution"`
	Node             string `form:"nodeOrder"  search:"type:order;column:node;table:fil_distribution"`
	AvailableBalance string `form:"availableBalanceOrder"  search:"type:order;column:available_balance;table:fil_distribution"`
	LastSectorPledge string `form:"lastSectorPledgeOrder"  search:"type:order;column:last_sector_pledge;table:fil_distribution"`
	CurSectorPledge  string `form:"curSectorPledgeOrder"  search:"type:order;column:cur_sector_pledge;table:fil_distribution"`
	EffectAmount     string `form:"effectAmountOrder"  search:"type:order;column:effect_amount;table:fil_distribution"`
	DistributeAmount string `form:"distributeAmountOrder"  search:"type:order;column:distribute_amount;table:fil_distribution"`
	AddressFrom      string `form:"addressFromOrder"  search:"type:order;column:address_from;table:fil_distribution"`
	AddressTo        string `form:"addressToOrder"  search:"type:order;column:address_to;table:fil_distribution"`
	Status           string `form:"statusOrder"  search:"type:order;column:status;table:fil_distribution"`
	CreatedAt        string `form:"createdAtOrder"  search:"type:order;column:created_at;table:fil_distribution"`
	UpdatedAt        string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:fil_distribution"`
	DeletedAt        string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:fil_distribution"`
	CreateBy         string `form:"createByOrder"  search:"type:order;column:create_by;table:fil_distribution"`
	UpdateBy         string `form:"updateByOrder"  search:"type:order;column:update_by;table:fil_distribution"`
}

func (m *FilDistributionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FilDistributionInsertReq struct {
	Id               int             `json:"-" comment:"编码"` // 编码
	NodeId           int             `json:"nodeId" comment:"节点ID"`
	Node             string          `json:"node" comment:"节点名称"`
	AvailableBalance decimal.Decimal `json:"availableBalance" comment:"可用余额"`
	HasTransfer      decimal.Decimal `json:"hasTransfer" comment:"已转出数量"`
	DistributePoint  decimal.Decimal `json:"distributePoint" comment:"分成比例"`
	LastSectorPledge decimal.Decimal `json:"lastSectorPledge" comment:"上期质押数量"`
	CurSectorPledge  decimal.Decimal `json:"curSectorPledge" comment:"当前质押数量"`
	EffectAmount     decimal.Decimal `json:"effectAmount" comment:"参与分币数量"`
	DistributeAmount decimal.Decimal `json:"distributeAmount" comment:"应该分币数量"`
	AddressFrom      string          `json:"addressFrom" comment:"发送地址"`
	AddressTo        string          `json:"addressTo" comment:"接收地址"`
	Status           int             `json:"status" comment:"分币状态"`
	common.ControlBy
}

func (s *FilDistributionInsertReq) Generate(model *models.FilDistribution) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.NodeId = s.NodeId
	model.Node = s.Node
	model.AvailableBalance = s.AvailableBalance
	model.HasTransfer = s.HasTransfer
	model.DistributePoint = s.DistributePoint
	model.LastSectorPledge = s.LastSectorPledge
	model.CurSectorPledge = s.CurSectorPledge
	model.EffectAmount = s.EffectAmount
	model.DistributeAmount = s.DistributeAmount
	model.AddressFrom = s.AddressFrom
	model.AddressTo = s.AddressTo
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FilDistributionInsertReq) GetId() interface{} {
	return s.Id
}

type FilDistributionUpdateReq struct {
	Id               int             `json:"id" comment:"编码"` // 编码
	NodeId           int             `json:"nodeId" comment:"节点ID"`
	Node             string          `json:"node" comment:"节点名称"`
	AvailableBalance decimal.Decimal `json:"availableBalance" comment:"可用余额"`
	HasTransfer      decimal.Decimal `json:"hasTransfer" comment:"已转出数量"`
	DistributePoint  decimal.Decimal `json:"distributePoint" comment:"分成比例"`
	LastSectorPledge decimal.Decimal `json:"lastSectorPledge" comment:"上期质押数量"`
	CurSectorPledge  decimal.Decimal `json:"curSectorPledge" comment:"当前质押数量"`
	EffectAmount     decimal.Decimal `json:"effectAmount" comment:"参与分币数量"`
	DistributeAmount decimal.Decimal `json:"distributeAmount" comment:"应该分币数量"`
	AddressFrom      string          `json:"addressFrom" comment:"发送地址"`
	AddressTo        string          `json:"addressTo" comment:"接收地址"`
	Status           int             `json:"status" comment:"分币状态"`
	common.ControlBy
}

func (s *FilDistributionUpdateReq) Generate(model *models.FilDistribution) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.NodeId = s.NodeId
	model.Node = s.Node
	model.AvailableBalance = s.AvailableBalance
	model.HasTransfer = s.HasTransfer
	model.DistributePoint = s.DistributePoint
	model.LastSectorPledge = s.LastSectorPledge
	model.CurSectorPledge = s.CurSectorPledge
	model.EffectAmount = s.EffectAmount
	model.DistributeAmount = s.DistributeAmount
	model.AddressFrom = s.AddressFrom
	model.AddressTo = s.AddressTo
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FilDistributionUpdateReq) GetId() interface{} {
	return s.Id
}

type FilDistributionUpdateStatusReq struct {
	Id     int `json:"-" comment:"编码"` // 编码
	Status int `json:"status" comment:"分币状态"`
	common.ControlBy
	common.ModelTime
}

func (s *FilDistributionUpdateStatusReq) Generate(model *models.FilDistribution) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy
}

func (s *FilDistributionUpdateStatusReq) GetId() interface{} {
	return s.Id
}

// FilDistributionGetReq 功能获取请求参数
type FilDistributionGetReq struct {
	Id int `uri:"id"`
}

func (s *FilDistributionGetReq) GetId() interface{} {
	return s.Id
}

// FilDistributionDeleteReq 功能删除请求参数
type FilDistributionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FilDistributionDeleteReq) GetId() interface{} {
	return s.Ids
}

func (s *FilDistributionGetPageReq) GetIds() interface{} {
	return s.NodeIds
}

type FilDistributionExport struct {
	Node             string           `json:"node" gorm:"type:varchar(128);comment:节点名称"`
	FilNode          *models.FilNodes `json:"filNode" gorm:"foreignKey:NodeId;references:Id;"`
	AvailableBalance decimal.Decimal  `json:"availableBalance" gorm:"type:decimal(20,8);comment:可用余额"`
	HasTransfer      decimal.Decimal  `json:"hasTransfer" gorm:"type:decimal(20,8);comment:本期分币前已转出数量"`
	DistributePoint  decimal.Decimal  `json:"distributePoint" gorm:"type:decimal(20,3);comment:分成比例"`
	LastSectorPledge decimal.Decimal  `json:"lastSectorPledge" gorm:"type:decimal(20,8);comment:上期质押数量"`
	CurSectorPledge  decimal.Decimal  `json:"curSectorPledge" gorm:"type:decimal(20,8);comment:当前质押数量"`
	EffectAmount     decimal.Decimal  `json:"effectAmount" gorm:"type:decimal(20,8);comment:参与分币数量=可用余额+转出数量"`
	DistributeAmount decimal.Decimal  `json:"distributeAmount" gorm:"type:decimal(20,8);comment:应该分币数量=参与数量x分成比例"`
	AddressTo        string           `json:"addressTo" gorm:"type:varchar(255);comment:接收地址"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	//FilNodeJoin      `search:"type:left;on:id:node_id;table:fil_distribution;join:fil_nodes" :"filNodeJoin"`
}
