package dto

import (
	"fil-admin/common/dto"
	"time"

	"github.com/shopspring/decimal"
)

type FilPoolChartGetPageReq struct {
	dto.Pagination  `search:"-"`
	Node            string          `form:"node"  search:"type:exact;column:node;table:pool_chart" comment:"账户名称"`
	MsigNode        string          `form:"msigNode"  search:"type:exact;column:msig_node;table:pool_chart" comment:"账户名称"`
	Status          string          `form:"status"  search:"type:exact;column:status;table:pool_chart" comment:"节点状态"`
	QualityAdjPower decimal.Decimal `form:"qualityAdjPower" search:"type:gt;column:status;table:pool_chart" comment:"有效算力"`
	Type            string          `form:"type"  search:"type:exact;column:type;table:pool_chart" comment:"节点类型"`
	DeptId          int             `json:"deptId" search:"type:exact;column:type;table:pool_chart" comment:部门ID"`
	FilPoolChartOrder
}

type FilPoolChartOrder struct {
	Id       string `form:"idOrder"  search:"type:order;column:id;table:pool_chart"`
	LastTime string `form:"lastTimeOrder"  search:"type:order;column:last_time;table:pool_chart"`
}

func (m *FilPoolChartGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// FilPoolChartGetReq 功能获取请求参数
type FilPoolChartGetReq struct {
	LastTime time.Time `form:"lastTime" search:"type:lte;column:last_time;table:pool_chart" comment:"时间"`
	DeptId   int       `json:"deptId" search:"type:exact;column:type;table:pool_chart" comment:"部门ID"`
	FilPoolChartOrder
}

func (s *FilPoolChartGetReq) GetId() interface{} {
	return s.Id
}

func (s *FilPoolChartGetReq) GetNeedSearch() interface{} {
	return *s
}
