package apis

import (
	"fmt"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	"fil-admin/common/middleware"
)

type SendMsg struct {
	api.Api
}

// GetPage 获取SendMsg列表
// @Summary 获取SendMsg列表
// @Description 获取SendMsg列表
// @Tags SendMsg
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SendMsg}} "{"code": 200, "data": [...]}"
// @Router /api/v1/send-msg [get]
// @Security Bearer
func (e SendMsg) GetPage(c *gin.Context) {
	req := dto.SendMsgGetPageReq{}
	s := service.SendMsg{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			// 查询当前部门下的节点构建in查询条件
			nApi := FilNodes{}
			nodes := nApi.NodeIds(c, deptId)
			req.Node = nodes
		}
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SendMsg, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SendMsg失败，\r\n失败信息 %s", err.Error()))
		return
	}
	newList := make([]models.SendMsg, 0)
	for _, v := range list {
		v.TypeStr = v.GetTypeStr().(string)
		newList = append(newList, v)
	}

	e.PageOK(newList, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取SendMsg
// @Summary 获取SendMsg
// @Description 获取SendMsg
// @Tags SendMsg
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SendMsg} "{"code": 200, "data": [...]}"
// @Router /api/v1/send-msg/{id} [get]
// @Security Bearer
func (e SendMsg) Get(c *gin.Context) {
	req := dto.SendMsgGetReq{}
	s := service.SendMsg{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SendMsg

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取SendMsg失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建SendMsg
// @Summary 创建SendMsg
// @Description 创建SendMsg
// @Tags SendMsg
// @Accept application/json
// @Product application/json
// @Param data body dto.SendMsgInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/send-msg [post]
// @Security Bearer
func (e SendMsg) Insert(c *gin.Context) {
	req := dto.SendMsgInsertReq{}
	s := service.SendMsg{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建SendMsg失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改SendMsg
// @Summary 修改SendMsg
// @Description 修改SendMsg
// @Tags SendMsg
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SendMsgUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/send-msg/{id} [put]
// @Security Bearer
func (e SendMsg) Update(c *gin.Context) {
	req := dto.SendMsgUpdateReq{}
	s := service.SendMsg{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改SendMsg失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除SendMsg
// @Summary 删除SendMsg
// @Description 删除SendMsg
// @Tags SendMsg
// @Param data body dto.SendMsgDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/send-msg [delete]
// @Security Bearer
func (e SendMsg) Delete(c *gin.Context) {
	s := service.SendMsg{}
	req := dto.SendMsgDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除SendMsg失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
