package models

import (
	admin "fil-admin/cmd/migrate/migration/models"
	"fil-admin/common/models"
)

type FilMsig struct {
	models.Model

	Address       string         `json:"address" gorm:"type:varchar(50);comment:Address"`
	RobustAddress string         `json:"robustAddress" gorm:"type:varchar(255);comment:RobustAddress"`
	Balance       string         `json:"balance" gorm:"type:decimal(20,8);comment:Balance"`
	DeptId        int            `json:"deptId" gorm:"type:int;comment:所属部门"`
	Dept          *admin.SysDept `json:"dept"`
	models.ModelTime
	models.ControlBy
}

func (FilMsig) TableName() string {
	return "fil_msig"
}

func (e *FilMsig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *FilMsig) GetId() interface{} {
	return e.Id
}
