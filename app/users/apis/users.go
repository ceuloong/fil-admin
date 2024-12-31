package apis

import (
	"fmt"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/app/users/models"
	"fil-admin/app/users/service"
	"fil-admin/app/users/service/dto"
	"fil-admin/common/actions"
)

type Users struct {
	api.Api
}

// GetPage 获取Users列表
// @Summary 获取Users列表
// @Description 获取Users列表
// @Tags Users
// @Param username query string false "用户名"
// @Param email query string false "邮箱"
// @Param verifyStatus query string false "认证状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Users}} "{"code": 200, "data": [...]}"
// @Router /api/v1/users [get]
// @Security Bearer
func (e Users) GetPage(c *gin.Context) {
	req := dto.UsersGetPageReq{}
	s := service.Users{}
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Users, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Users失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Users
// @Summary 获取Users
// @Description 获取Users
// @Tags Users
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Users} "{"code": 200, "data": [...]}"
// @Router /api/v1/users/{id} [get]
// @Security Bearer
func (e Users) Get(c *gin.Context) {
	req := dto.UsersGetReq{}
	s := service.Users{}
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
	var object models.Users

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Users失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Users
// @Summary 创建Users
// @Description 创建Users
// @Tags Users
// @Accept application/json
// @Product application/json
// @Param data body dto.UsersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/users [post]
// @Security Bearer
func (e Users) Insert(c *gin.Context) {
	req := dto.UsersInsertReq{}
	s := service.Users{}
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
		e.Error(500, err, fmt.Sprintf("创建Users失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Users
// @Summary 修改Users
// @Description 修改Users
// @Tags Users
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.UsersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/users/{id} [put]
// @Security Bearer
func (e Users) Update(c *gin.Context) {
	req := dto.UsersUpdateReq{}
	s := service.Users{}
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
		e.Error(500, err, fmt.Sprintf("修改Users失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Users
// @Summary 删除Users
// @Description 删除Users
// @Tags Users
// @Param data body dto.UsersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/users [delete]
// @Security Bearer
func (e Users) Delete(c *gin.Context) {
	s := service.Users{}
	req := dto.UsersDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Users失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Allocate 分配节点
// @Summary 分配节点给用户
// @Description 分配节点给用户
// @Tags Users
// @Accept application/json
// @Product application/json
// @Param id path int true "用户ID"
// @Param data body dto.UsersAllocateReq true "节点数据"
// @Success 200 {object} response.Response "{"code": 200, "message": "分配成功"}"
// @Router /api/v1/users/{id}/allocate [post]
// @Security Bearer
func (e Users) Allocate(c *gin.Context) {
	req := dto.UsersAllocateReq{}
	s := service.Users{}
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

	err = s.Allocate(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("分配节点失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "分配成功")
}
