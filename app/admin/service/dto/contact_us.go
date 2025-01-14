package dto

import (

	"fil-admin/app/admin/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type ContactUsGetPageReq struct {
	dto.Pagination     `search:"-"`
    Email string `form:"email"  search:"type:contains;column:email;table:contact_us" comment:"邮箱"`
    Subject string `form:"subject"  search:"type:contains;column:subject;table:contact_us" comment:"主题"`
    Message string `form:"message"  search:"type:contains;column:message;table:contact_us" comment:"内容"`
    ContactUsOrder
}

type ContactUsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:contact_us"`
    Email string `form:"emailOrder"  search:"type:order;column:email;table:contact_us"`
    Subject string `form:"subjectOrder"  search:"type:order;column:subject;table:contact_us"`
    Message string `form:"messageOrder"  search:"type:order;column:message;table:contact_us"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:contact_us"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:contact_us"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:contact_us"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:contact_us"`
    
}

func (m *ContactUsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ContactUsInsertReq struct {
    Id int `json:"-" comment:""` // 
    Email string `json:"email" comment:"邮箱"`
    Subject string `json:"subject" comment:"主题"`
    Message string `json:"message" comment:"内容"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *ContactUsInsertReq) Generate(model *models.ContactUs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Email = s.Email
    model.Subject = s.Subject
    model.Message = s.Message
    model.Status = s.Status
}

func (s *ContactUsInsertReq) GetId() interface{} {
	return s.Id
}

type ContactUsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Email string `json:"email" comment:"邮箱"`
    Subject string `json:"subject" comment:"主题"`
    Message string `json:"message" comment:"内容"`
    Status string `json:"status" comment:"状态"`
    common.ControlBy
}

func (s *ContactUsUpdateReq) Generate(model *models.ContactUs)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Email = s.Email
    model.Subject = s.Subject
    model.Message = s.Message
    model.Status = s.Status
}

func (s *ContactUsUpdateReq) GetId() interface{} {
	return s.Id
}

// ContactUsGetReq 功能获取请求参数
type ContactUsGetReq struct {
     Id int `uri:"id"`
}
func (s *ContactUsGetReq) GetId() interface{} {
	return s.Id
}

// ContactUsDeleteReq 功能删除请求参数
type ContactUsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ContactUsDeleteReq) GetId() interface{} {
	return s.Ids
}
