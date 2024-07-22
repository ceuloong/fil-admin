package models

import (
	"fil-admin/common/models"
)

type FinanceType struct {
	TypeId   int    `json:"typeId" gorm:"primaryKey;autoIncrement;"` //部门编码
	ParentId int    `json:"parentId" gorm:""`                        //上级部门
	Name     string `json:"name" gorm:"type:varchar(128);comment:收支名称"`
	InOut    int    `json:"inOut" gorm:"type:int;comment:收支  1收入   2支出"`
	TypePath string `json:"typePath" gorm:"type:varchar(255);comment:TypePath"`
	Memo     string `json:"memo" gorm:"type:varchar(255);comment:备注"`
	Sort     int    `json:"sort" gorm:"type:int;comment:排序"`
	Status   string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
	Children []FinanceType `json:"children" gorm:"-"`
}

func (FinanceType) TableName() string {
	return "finance_type"
}

func (e *FinanceType) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FinanceType) GetId() interface{} {
	return e.TypeId
}
