package apis

import (
	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"fil-admin/common/actions"
)

type FilMsig struct {
	api.Api
}

// GetPage 获取FilMsig列表
// @Summary 获取FilMsig列表
// @Description 获取FilMsig列表
// @Tags FilMsig
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FilMsig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-msig [get]
// @Security Bearer
func (e FilMsig) GetPage(c *gin.Context) {
	req := dto.FilMsigGetPageReq{}
	s := service.FilMsig{}
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
	list := make([]models.FilMsig, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		req.DeptId = fmt.Sprintf("/%d/", deptId)
	}
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilMsig失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取FilMsig
// @Summary 获取FilMsig
// @Description 获取FilMsig
// @Tags FilMsig
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FilMsig} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-msig/{id} [get]
// @Security Bearer
func (e FilMsig) Get(c *gin.Context) {
	req := dto.FilMsigGetReq{}
	s := service.FilMsig{}
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
	var object models.FilMsig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilMsig失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建FilMsig
// @Summary 创建FilMsig
// @Description 创建FilMsig
// @Tags FilMsig
// @Accept application/json
// @Product application/json
// @Param data body dto.FilMsigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fil-msig [post]
// @Security Bearer
func (e FilMsig) Insert(c *gin.Context) {
	req := dto.FilMsigInsertReq{}
	s := service.FilMsig{}
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
		e.Error(500, err, fmt.Sprintf("创建FilMsig失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改FilMsig
// @Summary 修改FilMsig
// @Description 修改FilMsig
// @Tags FilMsig
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FilMsigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fil-msig/{id} [put]
// @Security Bearer
func (e FilMsig) Update(c *gin.Context) {
	req := dto.FilMsigUpdateReq{}
	s := service.FilMsig{}
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
		e.Error(500, err, fmt.Sprintf("修改FilMsig失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除FilMsig
// @Summary 删除FilMsig
// @Description 删除FilMsig
// @Tags FilMsig
// @Param data body dto.FilMsigDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fil-msig [delete]
// @Security Bearer
func (e FilMsig) Delete(c *gin.Context) {
	s := service.FilMsig{}
	req := dto.FilMsigDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除FilMsig失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
