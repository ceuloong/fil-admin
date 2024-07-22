package dto

import (
	"fil-admin/app/finance/models"
	common "fil-admin/common/models"
)

type FinanceTypeGetPageReq struct {
	TypeId   int    `form:"typeId"  search:"type:exact;column:type_id;table:finance_type" comment:"编码"`
	ParentId string `form:"parentId"  search:"type:exact;column:parent_id;table:finance_type" comment:"父类型"`
	Name     string `form:"name"  search:"type:exact;column:name;table:finance_type" comment:"收支名称"`
	InOut    string `form:"inOut"  search:"type:exact;column:in_out;table:finance_type" comment:"收支  1收入   2支出"`
	Status   string `form:"status"  search:"type:exact;column:status;table:finance_type" comment:"状态"`
}

//type FinanceTypeOrder struct {
//	TypeId    string `form:"typeIdOrder"  search:"type:order;column:type_id;table:finance_type"`
//	ParentId  string `form:"parentIdOrder"  search:"type:order;column:parent_id;table:finance_type"`
//	Name      string `form:"nameOrder"  search:"type:order;column:name;table:finance_type"`
//	InOut     string `form:"inOutOrder"  search:"type:order;column:in_out;table:finance_type"`
//	TypePath  string `form:"typePathOrder"  search:"type:order;column:type_path;table:finance_type"`
//	Memo      string `form:"memoOrder"  search:"type:order;column:memo;table:finance_type"`
//	Sort      string `form:"sortOrder"  search:"type:order;column:sort;table:finance_type"`
//	Status    string `form:"statusOrder"  search:"type:order;column:status;table:finance_type"`
//	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:finance_type"`
//	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:finance_type"`
//	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:finance_type"`
//	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:finance_type"`
//	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:finance_type"`
//}

func (m *FinanceTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FinanceTypeInsertReq struct {
	TypeId   int    `json:"typeId" comment:"编码"` // 编码
	ParentId int    `json:"parentId" comment:"父类型"`
	Name     string `json:"name" comment:"收支名称"`
	InOut    int    `json:"inOut" comment:"收支  1收入   2支出"`
	TypePath string `json:"typePath" comment:""`
	Memo     string `json:"memo" comment:"备注"`
	Sort     int    `json:"sort" comment:"排序"`
	Status   string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *FinanceTypeInsertReq) Generate(model *models.FinanceType) {
	if s.TypeId == 0 {
		model.TypeId = s.TypeId
	}
	model.ParentId = s.ParentId
	model.Name = s.Name
	model.InOut = s.InOut
	model.TypePath = s.TypePath
	model.Memo = s.Memo
	model.Sort = s.Sort
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *FinanceTypeInsertReq) GetId() interface{} {
	return s.TypeId
}

type FinanceTypeUpdateReq struct {
	TypeId   int    `uri:"typeId" comment:"编码"` // 编码
	ParentId int    `json:"parentId" comment:"父类型"`
	Name     string `json:"name" comment:"收支名称"`
	InOut    int    `json:"inOut" comment:"收支  1收入   2支出"`
	TypePath string `json:"typePath" comment:""`
	Memo     string `json:"memo" comment:"备注"`
	Sort     int    `json:"sort" comment:"排序"`
	Status   string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *FinanceTypeUpdateReq) Generate(model *models.FinanceType) {
	if s.TypeId == 0 {
		model.TypeId = s.TypeId
	}
	model.ParentId = s.ParentId
	model.Name = s.Name
	model.InOut = s.InOut
	model.TypePath = s.TypePath
	model.Memo = s.Memo
	model.Sort = s.Sort
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *FinanceTypeUpdateReq) GetId() interface{} {
	return s.TypeId
}

// FinanceTypeGetReq 功能获取请求参数
type FinanceTypeGetReq struct {
	TypeId int `uri:"id"`
}

func (s *FinanceTypeGetReq) GetId() interface{} {
	return s.TypeId
}

// FinanceTypeDeleteReq 功能删除请求参数
type FinanceTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FinanceTypeDeleteReq) GetId() interface{} {
	return s.Ids
}

type TypeLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []TypeLabel `gorm:"-" json:"children"`
}
