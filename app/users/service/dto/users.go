package dto

import (
	"fil-admin/app/users/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type UsersGetPageReq struct {
	dto.Pagination `search:"-"`
	Username       string `form:"username"  search:"type:contains;column:username;table:users" comment:"用户名"`
	Email          string `form:"email"  search:"type:exact;column:email;table:users" comment:"邮箱"`
	VerifyStatus   string `form:"verifyStatus"  search:"type:exact;column:verify_status;table:users" comment:"认证状态"`
	Status         int    `form:"status"  search:"type:exact;column:status;table:users" comment:"状态"`
	UsersOrder
}

type UsersOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:users"`
	Username     string `form:"usernameOrder"  search:"type:order;column:username;table:users"`
	Email        string `form:"emailOrder"  search:"type:order;column:email;table:users"`
	PasswordHash string `form:"passwordHashOrder"  search:"type:order;column:password_hash;table:users"`
	VerifyStatus string `form:"verifyStatusOrder"  search:"type:order;column:verify_status;table:users"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:users"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:users"`
}

func (m *UsersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type UsersInsertReq struct {
	Id           int    `json:"-" comment:""` //
	Username     string `json:"username" comment:"用户名"`
	Email        string `json:"email" comment:"邮箱"`
	PasswordHash string `json:"passwordHash" comment:"密码"`
	VerifyStatus string `json:"verifyStatus" comment:"认证状态"`
	common.ControlBy
}

func (s *UsersInsertReq) Generate(model *models.Users) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Email = s.Email
	model.PasswordHash = s.PasswordHash
	model.VerifyStatus = s.VerifyStatus
}

func (s *UsersInsertReq) GetId() interface{} {
	return s.Id
}

type UsersUpdateReq struct {
	Id           int    `uri:"id" comment:""` //
	Username     string `json:"username" comment:"用户名"`
	Email        string `json:"email" comment:"邮箱"`
	VerifyStatus string `json:"verifyStatus" comment:"认证状态"`
	Status       int    `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *UsersUpdateReq) Generate(model *models.Users) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Username = s.Username
	model.Email = s.Email
	model.VerifyStatus = s.VerifyStatus
	model.Status = s.Status
}

func (s *UsersUpdateReq) GetId() interface{} {
	return s.Id
}

// UsersGetReq 功能获取请求参数
type UsersGetReq struct {
	Id int `uri:"id"`
}

func (s *UsersGetReq) GetId() interface{} {
	return s.Id
}

// UsersDeleteReq 功能删除请求参数
type UsersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *UsersDeleteReq) GetId() interface{} {
	return s.Ids
}

// UsersAllocateReq 分配节点请求结构
type UsersAllocateReq struct {
	Id      int   `uri:"id" comment:"用户ID"`
	NodeIds []int `json:"nodeIds" comment:"节点IDs"`
	common.ControlBy
}

func (s *UsersAllocateReq) GetId() interface{} {
	return s.Id
}
