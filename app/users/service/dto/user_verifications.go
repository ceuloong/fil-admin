package dto

import (

	"fil-admin/app/users/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type UserVerificationsGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId string `form:"userId"  search:"type:exact;column:user_id;table:user_verifications" comment:"用户Id"`
    RealName string `form:"realName"  search:"type:exact;column:real_name;table:user_verifications" comment:"真实姓名"`
    IdNumber string `form:"idNumber"  search:"type:exact;column:id_number;table:user_verifications" comment:"证件号码"`
    Status string `form:"status"  search:"type:exact;column:status;table:user_verifications" comment:"认证状态"`
    UserVerificationsOrder
}

type UserVerificationsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:user_verifications"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:user_verifications"`
    RealName string `form:"realNameOrder"  search:"type:order;column:real_name;table:user_verifications"`
    IdNumber string `form:"idNumberOrder"  search:"type:order;column:id_number;table:user_verifications"`
    IdFrontUrl string `form:"idFrontUrlOrder"  search:"type:order;column:id_front_url;table:user_verifications"`
    IdBackUrl string `form:"idBackUrlOrder"  search:"type:order;column:id_back_url;table:user_verifications"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:user_verifications"`
    RejectReason string `form:"rejectReasonOrder"  search:"type:order;column:reject_reason;table:user_verifications"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:user_verifications"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:user_verifications"`
    
}

func (m *UserVerificationsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type UserVerificationsInsertReq struct {
    Id int `json:"-" comment:""` // 
    UserId string `json:"userId" comment:"用户Id"`
    RealName string `json:"realName" comment:"真实姓名"`
    IdNumber string `json:"idNumber" comment:"证件号码"`
    IdFrontUrl string `json:"idFrontUrl" comment:"正面照片"`
    IdBackUrl string `json:"idBackUrl" comment:"背面照片"`
    Status string `json:"status" comment:"认证状态"`
    RejectReason string `json:"rejectReason" comment:"拒绝理由"`
    common.ControlBy
}

func (s *UserVerificationsInsertReq) Generate(model *models.UserVerifications)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.RealName = s.RealName
    model.IdNumber = s.IdNumber
    model.IdFrontUrl = s.IdFrontUrl
    model.IdBackUrl = s.IdBackUrl
    model.Status = s.Status
    model.RejectReason = s.RejectReason
}

func (s *UserVerificationsInsertReq) GetId() interface{} {
	return s.Id
}

type UserVerificationsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    UserId string `json:"userId" comment:"用户Id"`
    RealName string `json:"realName" comment:"真实姓名"`
    IdNumber string `json:"idNumber" comment:"证件号码"`
    IdFrontUrl string `json:"idFrontUrl" comment:"正面照片"`
    IdBackUrl string `json:"idBackUrl" comment:"背面照片"`
    Status string `json:"status" comment:"认证状态"`
    RejectReason string `json:"rejectReason" comment:"拒绝理由"`
    common.ControlBy
}

func (s *UserVerificationsUpdateReq) Generate(model *models.UserVerifications)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.RealName = s.RealName
    model.IdNumber = s.IdNumber
    model.IdFrontUrl = s.IdFrontUrl
    model.IdBackUrl = s.IdBackUrl
    model.Status = s.Status
    model.RejectReason = s.RejectReason
}

func (s *UserVerificationsUpdateReq) GetId() interface{} {
	return s.Id
}

// UserVerificationsGetReq 功能获取请求参数
type UserVerificationsGetReq struct {
     Id int `uri:"id"`
}
func (s *UserVerificationsGetReq) GetId() interface{} {
	return s.Id
}

// UserVerificationsDeleteReq 功能删除请求参数
type UserVerificationsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *UserVerificationsDeleteReq) GetId() interface{} {
	return s.Ids
}
