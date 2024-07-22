package dto

import (
	"fil-admin/app/filpool/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type FilMsigGetPageReq struct {
	dto.Pagination `search:"-"`
	Address        string `form:"address"  search:"type:exact;column:address;table:fil_msig" comment:"账户名称"`
	DeptJoin       `search:"type:left;on:dept_id:dept_id;table:fil_msig;join:sys_dept"`
	FilMsigOrder
}

type FilMsigOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:fil_msig"`
	Address       string `form:"addressOrder"  search:"type:order;column:address;table:fil_msig"`
	RobustAddress string `form:"robustAddressOrder"  search:"type:order;column:robust_address;table:fil_msig"`
	Balance       string `form:"balanceOrder"  search:"type:order;column:balance;table:fil_msig"`
}

func (m *FilMsigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FilMsigInsertReq struct {
	Id            int    `json:"-" comment:""` //
	Address       string `json:"address" comment:""`
	RobustAddress string `json:"robustAddress" comment:""`
	Balance       string `json:"balance" comment:""`
	DeptId        int    `json:"deptId" comment:"所属部门"`
	common.ControlBy
}

func (s *FilMsigInsertReq) Generate(model *models.FilMsig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Address = s.Address
	model.RobustAddress = s.RobustAddress
	model.Balance = s.Balance
	model.DeptId = s.DeptId
}

func (s *FilMsigInsertReq) GetId() interface{} {
	return s.Id
}

type FilMsigUpdateReq struct {
	Id            int    `uri:"id" comment:""` //
	Address       string `json:"address" comment:""`
	RobustAddress string `json:"robustAddress" comment:""`
	Balance       string `json:"balance" comment:""`
	DeptId        int    `json:"deptId" comment:"所属部门"`
	common.ControlBy
}

func (s *FilMsigUpdateReq) Generate(model *models.FilMsig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Address = s.Address
	model.RobustAddress = s.RobustAddress
	model.Balance = s.Balance
	model.DeptId = s.DeptId
	model.ControlBy = s.ControlBy
}

func (s *FilMsigUpdateReq) GetId() interface{} {
	return s.Id
}

// FilMsigGetReq 功能获取请求参数
type FilMsigGetReq struct {
	Id int `uri:"id"`
}

func (s *FilMsigGetReq) GetId() interface{} {
	return s.Id
}

// FilMsigDeleteReq 功能删除请求参数
type FilMsigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FilMsigDeleteReq) GetId() interface{} {
	return s.Ids
}
