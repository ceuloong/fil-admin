package dto

import (
	"time"

	"fil-admin/app/filpool/models"
	"fil-admin/common/dto"
	common "fil-admin/common/models"
)

type SendMsgGetPageReq struct {
	dto.Pagination `search:"-"`
	SendMsgOrder
}

type SendMsgOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:send_msg"`
	Title      string `form:"titleOrder"  search:"type:order;column:title;table:send_msg"`
	Node       string `form:"nodeOrder"  search:"type:order;column:node;table:send_msg"`
	Content    string `form:"contentOrder"  search:"type:order;column:content;table:send_msg"`
	CreateTime string `form:"createTimeOrder"  search:"type:order;column:create_time;table:send_msg"`
	SendTime   string `form:"sendTimeOrder"  search:"type:order;column:send_time;table:send_msg"`
	SendStatus string `form:"sendStatusOrder"  search:"type:order;column:send_status;table:send_msg"`
	Type       string `form:"typeOrder"  search:"type:order;column:type;table:send_msg"`
}

func (m *SendMsgGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SendMsgInsertReq struct {
	Id         int             `json:"-" comment:""` //
	Title      string          `json:"title" comment:""`
	Node       string          `json:"node" comment:""`
	Content    string          `json:"content" comment:""`
	CreateTime time.Time       `json:"createTime" comment:""`
	SendTime   *time.Time      `json:"sendTime" comment:""`
	SendStatus int             `json:"sendStatus" comment:""`
	Type       models.SendType `json:"type" comment:"消息类型"`
	common.ControlBy
}

func (s *SendMsgInsertReq) Generate(model *models.SendMsg) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Node = s.Node
	model.Content = s.Content
	model.CreateTime = s.CreateTime
	model.SendTime = s.SendTime
	model.SendStatus = s.SendStatus
	model.Type = s.Type
}

func (s *SendMsgInsertReq) GetId() interface{} {
	return s.Id
}

type SendMsgUpdateReq struct {
	Id         int `uri:"id" comment:""` //
	SendStatus int `json:"sendStatus" comment:""`
	common.ControlBy
}

func (s *SendMsgUpdateReq) Generate(model *models.SendMsg) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.SendStatus = s.SendStatus
}

func (s *SendMsgUpdateReq) GetId() interface{} {
	return s.Id
}

// SendMsgGetReq 功能获取请求参数
type SendMsgGetReq struct {
	Id int `uri:"id"`
}

func (s *SendMsgGetReq) GetId() interface{} {
	return s.Id
}

// SendMsgDeleteReq 功能删除请求参数
type SendMsgDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SendMsgDeleteReq) GetId() interface{} {
	return s.Ids
}
