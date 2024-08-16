package apis

import (
	"fil-admin/app/finance/models"
	"fil-admin/app/finance/service"
	"fil-admin/app/finance/service/dto"
	"fmt"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/common/actions"
)

type FinanceCash struct {
	api.Api
}

// GetPage 获取现金收支明细列表
// @Summary 获取现金收支明细列表
// @Description 获取现金收支明细列表
// @Tags 现金收支明细
// @Param name query string false "货币名称"
// @Param amount query string false "金额"
// @Param type query string false "收支类型"
// @Param memo query string false "备注"
// @Param employee query string false "员工名称"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FinanceCash}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance-cash [get]
// @Security Bearer
func (e FinanceCash) GetPage(c *gin.Context) {
	req := dto.FinanceCashGetPageReq{}
	s := service.FinanceCash{}
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
	list := make([]models.FinanceCash, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取现金收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取现金收支明细
// @Summary 获取现金收支明细
// @Description 获取现金收支明细
// @Tags 现金收支明细
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FinanceCash} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance-cash/{id} [get]
// @Security Bearer
func (e FinanceCash) Get(c *gin.Context) {
	req := dto.FinanceCashGetReq{}
	s := service.FinanceCash{}
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
	var object models.FinanceCash

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取现金收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建现金收支明细
// @Summary 创建现金收支明细
// @Description 创建现金收支明细
// @Tags 现金收支明细
// @Accept application/json
// @Product application/json
// @Param data body dto.FinanceCashInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance-cash [post]
// @Security Bearer
func (e FinanceCash) Insert(c *gin.Context) {
	req := dto.FinanceCashInsertReq{}
	s := service.FinanceCash{}
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
		e.Error(500, err, fmt.Sprintf("创建现金收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改现金收支明细
// @Summary 修改现金收支明细
// @Description 修改现金收支明细
// @Tags 现金收支明细
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FinanceCashUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance-cash/{id} [put]
// @Security Bearer
func (e FinanceCash) Update(c *gin.Context) {
	req := dto.FinanceCashUpdateReq{}
	s := service.FinanceCash{}
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
		e.Error(500, err, fmt.Sprintf("修改现金收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除现金收支明细
// @Summary 删除现金收支明细
// @Description 删除现金收支明细
// @Tags 现金收支明细
// @Param data body dto.FinanceCashDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance-cash [delete]
// @Security Bearer
func (e FinanceCash) Delete(c *gin.Context) {
	s := service.FinanceCash{}
	req := dto.FinanceCashDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除现金收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e FinanceCash) ExportXlsx(c *gin.Context) {
	req := dto.FinanceCashGetPageReq{}
	s := service.FinanceCash{}
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
	err = s.ExportXlsx(&req, e.Context)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建现金收支明细失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK("", "创建成功")
}
