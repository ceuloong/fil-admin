package dto

import (
	"fil-admin/app/filpool/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
	"time"

	"github.com/shopspring/decimal"
)

type FilNodesGetPageReq struct {
	dto.Pagination `search:"-"`
	Node           string `form:"node"  search:"type:exact;column:node;table:fil_nodes" comment:"账户名称"`
	MsigNode       string `form:"msigNode"  search:"type:exact;column:msig_node;table:fil_nodes" comment:"账户名称"`
	Status         string `form:"status"  search:"type:exact;column:status;table:fil_nodes" comment:"节点状态"`
	MinPower       int    `form:"minPower" search:"type:gt;column:quality_adj_power;table:fil_nodes" comment:"最小算力"`
	MaxPower       int    `form:"maxPower" search:"type:lte;column:quality_adj_power;table:fil_nodes" comment:"最大算力"`
	Type           string `form:"type"  search:"type:exact;column:type;table:fil_nodes" comment:"节点类型"`
	ControlAddress string `form:"controlAddress" search:"type:exact;column:control_address;table:fil_nodes" comment:"控制地址"`
	DeptJoin       `search:"type:left;on:dept_id:dept_id;table:fil_nodes;join:sys_dept"`
	FilNodesOrder
}

type DeptJoin struct {
	DeptId string `search:"type:contains;column:dept_path;table:sys_dept" form:"deptId"`
}

type FilNodesOrder struct {
	Id                         string `form:"idOrder"  search:"type:order;column:id;table:fil_nodes"`
	Node                       string `form:"nodeOrder"  search:"type:order;column:node;table:fil_nodes"`
	MsigNode                   string `form:"msigNodeOrder"  search:"type:order;column:msig_node;table:fil_nodes"`
	Address                    string `form:"addressOrder"  search:"type:order;column:address;table:fil_nodes"`
	MsgCount                   string `form:"msgCountOrder"  search:"type:order;column:msg_count;table:fil_nodes"`
	Type                       string `form:"typeOrder"  search:"type:order;column:type;table:fil_nodes"`
	DistributePoint            string `form:"distributePointOrder"  search:"type:order;column:distribute_point;table:fil_nodes"`
	CreateTime                 string `form:"createTimeOrder"  search:"type:order;column:create_time;table:fil_nodes"`
	AvailableBalance           string `form:"availableBalanceOrder"  search:"type:order;column:available_balance;table:fil_nodes"`
	Balance                    string `form:"balanceOrder"  search:"type:order;column:balance;table:fil_nodes"`
	SectorPledgeBalance        string `form:"sectorPledgeBalanceOrder"  search:"type:order;column:sector_pledge_balance;table:fil_nodes"`
	VestingFunds               string `form:"vestingFundsOrder"  search:"type:order;column:vesting_funds;table:fil_nodes"`
	Height                     string `form:"heightOrder"  search:"type:order;column:height;table:fil_nodes"`
	LastTime                   string `form:"lastTimeOrder"  search:"type:order;column:last_time;table:fil_nodes"`
	RewardValue                string `form:"rewardValueOrder"  search:"type:order;column:reward_value;table:fil_nodes"`
	QualityAdjPower            string `form:"qualityAdjPowerOrder"  search:"type:order;column:quality_adj_power;table:fil_nodes"`
	PowerUnit                  string `form:"powerUnitOrder"  search:"type:order;column:power_unit;table:fil_nodes"`
	PowerPoint                 string `form:"powerPointOrder"  search:"type:order;column:power_point;table:fil_nodes"`
	PowerGrade                 string `form:"powerGradeOrder"  search:"type:order;column:power_grade;table:fil_nodes"`
	SectorSize                 string `form:"sectorSizeOrder"  search:"type:order;column:sector_size;table:fil_nodes"`
	SectorStatus               string `form:"sectorStatusOrder"  search:"type:order;column:sector_status;table:fil_nodes"`
	ControlAddress             string `form:"controlAddressOrder"  search:"type:order;column:control_address;table:fil_nodes"`
	ControlBalance             string `form:"controlBalanceOrder"  search:"type:order;column:control_balance;table:fil_nodes"`
	EndTime                    string `form:"endTimeOrder"  search:"type:order;column:end_time;table:fil_nodes"`
	LastDisSectorPledgeBalance string `form:"lastSectorPledgeBalanceOrder"  search:"type:order;column:last_sector_pledge_balance;table:fil_nodes"`
	BlocksMined                string `form:"blocksMinedOrder"  search:"type:order;column:blocks_mined;table:fil_nodes"`
	TotalRewards24h            string `form:"totalRewards24hOrder"  search:"type:order;column:total_rewards24h;table:fil_nodes"`
	LuckyValue24h              string `form:"luckyValue24hOrder"  search:"type:order;column:lucky_value24h;table:fil_nodes"`
	ReceiveAmount              string `form:"receiveAmountOrder"  search:"type:order;column:receive_amount;table:fil_nodes"`
	SendAmount                 string `form:"sendAmountOrder"  search:"type:order;column:send_amount;table:fil_nodes"`
	BurnAmount                 string `form:"burnAmountOrder"  search:"type:order;column:burn_amount;table:fil_nodes"`
}

func (m *FilNodesGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FilNodesInsertReq struct {
	Id                         int             `json:"-" comment:""` //
	Node                       string          `json:"node" comment:"账户名称"`
	MsigNode                   string          `json:"msigNode" comment:"所属地址"`
	Address                    string          `json:"address" comment:"地址"`
	DistributePoint            decimal.Decimal `json:"distributePoint" comment:"分成比例"`
	HasTransfer                decimal.Decimal `json:"hasTransfer" comment:"上次分成后转出数量"`
	HasRealDistribute          decimal.Decimal `json:"hasRealDistribute" comment:"实际已分币数量"`
	LastDisSectorPledgeBalance decimal.Decimal `json:"lastSectorPledgeBalance" comment:"上次分币时质押数量"`
	LastDistributeTime         time.Time       `json:"lastDistributeTime" comment:"最后一次分币时间"`
	Status                     string          `json:"status"`
	Type                       string          `json:"type"`
	EndTime                    time.Time       `json:"endTime" comment:"节点结束时间"`
	DeptId                     int             `json:"deptId" comment:"所属部门"`
	common.ControlBy
	common.ModelTime
}

func (FilNodesInsertReq) TableName() string {
	return "fil_nodes"
}

func (s *FilNodesInsertReq) Generate(model *models.FilNodes) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Node = s.Node
	model.MsigNode = s.MsigNode
	model.Address = s.Address
	model.DistributePoint = s.DistributePoint
	model.HasTransfer = s.HasTransfer
	model.HasRealDistribute = s.HasRealDistribute
	model.LastDisSectorPledgeBalance = s.LastDisSectorPledgeBalance
	model.LastDistributeTime = s.LastDistributeTime
	model.Status = s.Status
	model.Type = s.Type
	model.EndTime = s.EndTime
	model.CreateTime = time.Now()
	model.LastTime = time.Now()
}

func (s *FilNodesInsertReq) GetId() interface{} {
	return s.Id
}

type FilNodesUpdateReq struct {
	Id                         int             `uri:"id" comment:""` //
	Node                       string          `json:"node" comment:"账户名称"`
	MsigNode                   string          `json:"msigNode" comment:"所属地址"`
	Address                    string          `json:"address" comment:"地址"`
	DistributePoint            decimal.Decimal `json:"distributePoint" comment:"分成比例"`
	HasTransfer                decimal.Decimal `json:"hasTransfer" comment:"上次分成后转出数量"`
	HasRealDistribute          decimal.Decimal `json:"hasRealDistribute" comment:"实际已分币数量"`
	LastDisSectorPledgeBalance decimal.Decimal `json:"lastSectorPledgeBalance" comment:"上次分币时质押数量"`
	LastDistributeTime         time.Time       `json:"lastDistributeTime" comment:"最后一次分币时间"`
	Status                     string          `json:"status"`
	Type                       string          `json:"type"`
	EndTime                    time.Time       `json:"endTime" comment:"节点结束时间"`
	Title                      string          `json:"title"`
	DeptId                     int             `json:"deptId"`
	common.ControlBy
	common.ModelTime
}

func (s *FilNodesUpdateReq) Generate(model *models.FilNodes) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Node = s.Node
	model.MsigNode = s.MsigNode
	model.Address = s.Address
	model.DistributePoint = s.DistributePoint
	model.HasTransfer = s.HasTransfer
	model.HasRealDistribute = s.HasRealDistribute
	model.LastDisSectorPledgeBalance = s.LastDisSectorPledgeBalance
	model.LastDistributeTime = s.LastDistributeTime
	model.Status = s.Status
	model.Type = s.Type
	model.EndTime = s.EndTime
	model.Title = s.Title
}

func (s *FilNodesUpdateReq) GetId() interface{} {
	return s.Id
}

type FilNodesUpdateTitleReq struct {
	Id    int    `uri:"id" comment:""` //
	Title string `json:"title"`
	common.ControlBy
}

func (s *FilNodesUpdateTitleReq) Generate(model *models.FilNodes) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.UpdateBy = s.UpdateBy
}

func (s *FilNodesUpdateTitleReq) GetId() interface{} {
	return s.Id
}

type FilNodesUpdateDistributeReq struct {
	Id                         int             `uri:"id" comment:""` //
	HasRealDistribute          decimal.Decimal `json:"hasRealDistribute" comment:"实际已分币数量"`
	LastDisSectorPledgeBalance decimal.Decimal `json:"lastDisSectorPledgeBalance" comment:"上次分币时质押数量"`
	LastDistributeTime         time.Time       `json:"lastDistributeTime" comment:"最后一次分币时间"`
	HasTransfer                decimal.Decimal `json:"hasTransfer" comment:"分币前已转出数量，分完币清0"`
	common.ControlBy
}

func (s *FilNodesUpdateDistributeReq) Generate(model *models.FilNodes) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.HasRealDistribute = s.HasRealDistribute
	model.LastDisSectorPledgeBalance = s.LastDisSectorPledgeBalance
	model.LastDistributeTime = s.LastDistributeTime
}

func (s *FilNodesUpdateDistributeReq) GetId() interface{} {
	return s.Id
}

// FilNodesGetReq 功能获取请求参数
type FilNodesGetReq struct {
	Id int `uri:"id"`
}

func (s *FilNodesGetReq) GetId() interface{} {
	return s.Id
}

// FilNodesDeleteReq 功能删除请求参数
type FilNodesDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FilNodesDeleteReq) GetId() interface{} {
	return s.Ids
}
