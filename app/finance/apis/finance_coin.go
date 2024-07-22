package apis

import (
	"fil-admin/app/finance/models"
	"fil-admin/app/finance/service"
	"fil-admin/app/finance/service/dto"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"fil-admin/common/actions"
)

type FinanceCoin struct {
	api.Api
}

// GetPage 获取数字货币收支明细列表
// @Summary 获取数字货币收支明细列表
// @Description 获取数字货币收支明细列表
// @Tags 数字货币收支明细
// @Param name query string false "币种名称"
// @Param address query string false "收币地址"
// @Param rate query string false "汇率"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FinanceCoin}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance-coin [get]
// @Security Bearer
func (e FinanceCoin) GetPage(c *gin.Context) {
	req := dto.FinanceCoinGetPageReq{}
	s := service.FinanceCoin{}
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
	list := make([]models.FinanceCoin, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取数字货币收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取数字货币收支明细
// @Summary 获取数字货币收支明细
// @Description 获取数字货币收支明细
// @Tags 数字货币收支明细
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FinanceCoin} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance-coin/{id} [get]
// @Security Bearer
func (e FinanceCoin) Get(c *gin.Context) {
	req := dto.FinanceCoinGetReq{}
	s := service.FinanceCoin{}
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
	var object models.FinanceCoin

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取数字货币收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建数字货币收支明细
// @Summary 创建数字货币收支明细
// @Description 创建数字货币收支明细
// @Tags 数字货币收支明细
// @Accept application/json
// @Product application/json
// @Param data body dto.FinanceCoinInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance-coin [post]
// @Security Bearer
func (e FinanceCoin) Insert(c *gin.Context) {
	req := dto.FinanceCoinInsertReq{}
	s := service.FinanceCoin{}
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
		e.Error(500, err, fmt.Sprintf("创建数字货币收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改数字货币收支明细
// @Summary 修改数字货币收支明细
// @Description 修改数字货币收支明细
// @Tags 数字货币收支明细
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FinanceCoinUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance-coin/{id} [put]
// @Security Bearer
func (e FinanceCoin) Update(c *gin.Context) {
	req := dto.FinanceCoinUpdateReq{}
	s := service.FinanceCoin{}
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
		e.Error(500, err, fmt.Sprintf("修改数字货币收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除数字货币收支明细
// @Summary 删除数字货币收支明细
// @Description 删除数字货币收支明细
// @Tags 数字货币收支明细
// @Param data body dto.FinanceCoinDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance-coin [delete]
// @Security Bearer
func (e FinanceCoin) Delete(c *gin.Context) {
	s := service.FinanceCoin{}
	req := dto.FinanceCoinDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除数字货币收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
