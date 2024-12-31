package models

import (
	"fil-admin/common/models"
)

type UserVerifications struct {
	models.Model

	UserId       string `json:"userId" gorm:"type:int;comment:用户Id"`
	RealName     string `json:"realName" gorm:"type:varchar(50);comment:真实姓名"`
	IdNumber     string `json:"idNumber" gorm:"type:varchar(18);comment:证件号码"`
	IdFrontUrl   string `json:"idFrontUrl" gorm:"type:varchar(255);comment:正面照片"`
	IdBackUrl    string `json:"idBackUrl" gorm:"type:varchar(255);comment:背面照片"`
	Status       string `json:"status" gorm:"type:enum('pending','verified','rejected');comment:认证状态"`
	RejectReason string `json:"rejectReason" gorm:"type:text;comment:拒绝理由"`
	Users        Users  `json:"users" gorm:"foreignKey:UserId;references:Id"`
	models.ModelTime
	models.ControlBy
}

func (UserVerifications) TableName() string {
	return "user_verifications"
}

func (e *UserVerifications) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserVerifications) GetId() interface{} {
	return e.Id
}
