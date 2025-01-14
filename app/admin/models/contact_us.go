package models

import (

	"fil-admin/common/models"

)

type ContactUs struct {
    models.Model
    
    Email string `json:"email" gorm:"type:varchar(255);comment:邮箱"` 
    Subject string `json:"subject" gorm:"type:varchar(255);comment:主题"` 
    Message string `json:"message" gorm:"type:text;comment:内容"` 
    Status string `json:"status" gorm:"type:enum('pending','replied');comment:状态"` 
    models.ModelTime
    models.ControlBy
}

func (ContactUs) TableName() string {
    return "contact_us"
}

func (e *ContactUs) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ContactUs) GetId() interface{} {
	return e.Id
}