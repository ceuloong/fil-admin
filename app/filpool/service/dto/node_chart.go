package dto

import (
	"fil-admin/common/dto"
	"time"
)

type NodeChartGetPageReq struct {
	dto.Pagination `search:"-"`
	Node           string `form:"node"  search:"type:exact;column:node;table:nodes_chart" comment:"账户名称"`
	MsigNode       string `form:"msigNode"  search:"type:exact;column:msig_node;table:nodes_chart" comment:"账户名称"`
	Status         string `form:"status"  search:"type:exact;column:status;table:nodes_chart" comment:"节点状态"`
	LastTime       string `form:"lastTime" search:"type:exact;column:last_time;table:nodes_chart" comment:"创建时间"`
	NodeChartOrder
}

type NodeChartOrder struct {
	Id       string `form:"idOrder"  search:"type:order;column:id;table:nodes_chart"`
	LastTime string `form:"lastTimeOrder"  search:"type:order;column:last_time;table:nodes_chart"`
}

func (m *NodeChartGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// NodeChartGetReq 功能获取请求参数
type NodeChartGetReq struct {
	LastTime time.Time `form:"lastTime" search:"type:gte;column:last_time;table:nodes_chart" comment:"时间"`
	Nodes    []string  `form:"nodes" search:"type:in;column:node;table:nodes_chart" comment:"node"`
	NodeChartOrder
}

func (s *NodeChartGetReq) GetId() interface{} {
	return s.Id
}

func (s *NodeChartGetReq) GetNeedSearch() interface{} {
	return *s
}
