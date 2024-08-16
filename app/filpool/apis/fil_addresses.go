package apis

import (
	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fmt"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/common/actions"
)

type FilAddresses struct {
	api.Api
}

// GetPage 获取FilAddresses列表
// @Summary 获取FilAddresses列表
// @Description 获取FilAddresses列表
// @Tags FilAddresses
// @Param node query string false ""
// @Param accountId query string false ""
// @Param address query string false ""
// @Param type query string false "controller, worker, other"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FilAddresses}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-addresses [get]
// @Security Bearer
func (e FilAddresses) GetPage(c *gin.Context) {
	req := dto.FilAddressesGetPageReq{}
	s := service.FilAddresses{}
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
	list := make([]models.FilAddresses, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilAddresses失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取FilAddresses
// @Summary 获取FilAddresses
// @Description 获取FilAddresses
// @Tags FilAddresses
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FilAddresses} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-addresses/{id} [get]
// @Security Bearer
func (e FilAddresses) Get(c *gin.Context) {
	req := dto.FilAddressesGetReq{}
	s := service.FilAddresses{}
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
	var object models.FilAddresses

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilAddresses失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建FilAddresses
// @Summary 创建FilAddresses
// @Description 创建FilAddresses
// @Tags FilAddresses
// @Accept application/json
// @Product application/json
// @Param data body dto.FilAddressesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fil-addresses [post]
// @Security Bearer
func (e FilAddresses) Insert(c *gin.Context) {
	req := dto.FilAddressesInsertReq{}
	s := service.FilAddresses{}
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
		e.Error(500, err, fmt.Sprintf("创建FilAddresses失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改FilAddresses
// @Summary 修改FilAddresses
// @Description 修改FilAddresses
// @Tags FilAddresses
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FilAddressesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fil-addresses/{id} [put]
// @Security Bearer
func (e FilAddresses) Update(c *gin.Context) {
	req := dto.FilAddressesUpdateReq{}
	s := service.FilAddresses{}
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
		e.Error(500, err, fmt.Sprintf("修改FilAddresses失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除FilAddresses
// @Summary 删除FilAddresses
// @Description 删除FilAddresses
// @Tags FilAddresses
// @Param data body dto.FilAddressesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fil-addresses [delete]
// @Security Bearer
func (e FilAddresses) Delete(c *gin.Context) {
	s := service.FilAddresses{}
	req := dto.FilAddressesDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除FilAddresses失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
