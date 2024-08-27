package dto

import (
	"time"

	"fil-admin/app/filpool/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type BlockGetPageReq struct {
	dto.Pagination `search:"-"`
	Height         int    `form:"height"  search:"type:exact;column:height;table:block" comment:"高度"`
	Node           string `form:"node"  search:"type:exact;column:node;table:block" comment:"节点"`
	Message        string `form:"message"  search:"type:exact;column:message;table:block" comment:"区块哈希"`
	Status         string `form:"status"  search:"type:exact;column:status;table:block" comment:"1正常   2孤块"`
	BlockOrder
}

type BlockOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:block"`
	Height      string `form:"heightOrder"  search:"type:order;column:height;table:block"`
	Node        string `form:"nodeOrder"  search:"type:order;column:node;table:block"`
	BlockTime   string `form:"blockTimeOrder"  search:"type:order;column:block_time;table:block"`
	NodeFrom    string `form:"nodeFromOrder"  search:"type:order;column:node_from;table:block"`
	NodeTo      string `form:"nodeToOrder"  search:"type:order;column:node_to;table:block"`
	Message     string `form:"messageOrder"  search:"type:order;column:message;table:block"`
	RewardValue string `form:"rewardValueOrder"  search:"type:order;column:reward_value;table:block"`
	MsgCount    string `form:"msgCountOrder"  search:"type:order;column:msg_count;table:block"`
	BlockSize   string `form:"blockSizeOrder"  search:"type:order;column:block_size;table:block"`
	Status      string `form:"statusOrder"  search:"type:order;column:status;table:block"`
	CreateTime  string `form:"createTimeOrder"  search:"type:order;column:create_time;table:block"`
}

func (m *BlockGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BlockInsertReq struct {
	Id          int       `json:"-" comment:"高度"`
	Height      int       `json:"height" comment:"高度"` // 高度
	Node        string    `json:"node" comment:"节点"`
	BlockTime   time.Time `json:"blockTime" comment:"区块时间"`
	NodeFrom    string    `json:"nodeFrom" comment:""`
	NodeTo      string    `json:"nodeTo" comment:""`
	Message     string    `json:"message" comment:"区块哈希"`
	RewardValue string    `json:"rewardValue" comment:"奖励数量"`
	MsgCount    string    `json:"msgCount" comment:"消息数"`
	BlockSize   string    `json:"blockSize" comment:"区块大小"`
	Status      string    `json:"status" comment:"1正常   2孤块"`
	CreateTime  time.Time `json:"createTime" comment:"记录的创建时间"`
	common.ControlBy
}

func (s *BlockInsertReq) Generate(model *models.Block) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Height = s.Height
	model.Node = s.Node
	model.BlockTime = s.BlockTime
	model.NodeFrom = s.NodeFrom
	model.NodeTo = s.NodeTo
	model.Message = s.Message
	model.RewardValue = s.RewardValue
	model.MsgCount = s.MsgCount
	model.BlockSize = s.BlockSize
	model.Status = s.Status
	model.CreateTime = s.CreateTime
}

func (s *BlockInsertReq) GetId() interface{} {
	return s.Height
}

type BlockUpdateReq struct {
	Id         int       `uri:"id" comment:""` //
	BlockSize  string    `json:"blockSize" comment:"区块大小"`
	Status     string    `json:"status" comment:"1正常   2孤块"`
	CreateTime time.Time `json:"createTime" comment:"记录的创建时间"`
	common.ControlBy
}

func (s *BlockUpdateReq) Generate(model *models.Block) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BlockSize = s.BlockSize
	model.Status = s.Status
	model.CreateTime = s.CreateTime
}

func (s *BlockUpdateReq) GetId() interface{} {
	return s.Id
}

// BlockGetReq 功能获取请求参数
type BlockGetReq struct {
	Id int `uri:"id"`
}

func (s *BlockGetReq) GetId() interface{} {
	return s.Id
}

// BlockDeleteReq 功能删除请求参数
type BlockDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BlockDeleteReq) GetId() interface{} {
	return s.Ids
}
