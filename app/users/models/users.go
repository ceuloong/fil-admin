package models

import (
	"fil-admin/common/models"
)

type Users struct {
	models.Model

	Username     string `json:"username" gorm:"type:varchar(50);comment:用户名"`
	Email        string `json:"email" gorm:"type:varchar(100);comment:邮箱"`
	PasswordHash string `json:"-" gorm:"type:varchar(255);comment:密码"`
	Status       int    `json:"status" gorm:"type:tinyint(1);comment:状态"`
	VerifyStatus string `json:"verifyStatus" gorm:"type:enum('pending','verified','rejected');comment:认证状态"`
	NodeIds      string `json:"nodeIds" gorm:"type:text;comment:节点Id"`
	models.ModelTime
	models.ControlBy
}

func (Users) TableName() string {
	return "users"
}

func (e *Users) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Users) GetId() interface{} {
	return e.Id
}
