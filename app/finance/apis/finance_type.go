package apis

import (
	"fil-admin/app/finance/models"
	"fil-admin/app/finance/service"
	"fil-admin/app/finance/service/dto"
	"fmt"

	"github.com/gin-gonic/gin/binding"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"
)

type FinanceType struct {
	api.Api
}

// GetPage 获取收支类型列表
// @Summary 获取收支类型列表
// @Description 获取收支类型列表
// @Tags 收支类型
// @Param typeId query int false "编码"
// @Param parentId query string false "父类型"
// @Param name query string false "收支名称"
// @Param inOut query string false "收支  1收入   2支出"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FinanceType}} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance-type [get]
// @Security Bearer
func (e FinanceType) GetPage(c *gin.Context) {
	req := dto.FinanceTypeGetPageReq{}
	s := service.FinanceType{}
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
	list := make([]models.FinanceType, 0)
	list, err = s.SetTypePage(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取收支类型失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(list, "查询成功")
}

// Get 获取收支类型
// @Summary 获取收支类型
// @Description 获取收支类型
// @Tags 收支类型
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FinanceType} "{"code": 200, "data": [...]}"
// @Router /api/v1/finance-type/{id} [get]
// @Security Bearer
func (e FinanceType) Get(c *gin.Context) {
	req := dto.FinanceTypeGetReq{}
	s := service.FinanceType{}
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
	var object models.FinanceType

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取收支类型失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建收支类型
// @Summary 创建收支类型
// @Description 创建收支类型
// @Tags 收支类型
// @Accept application/json
// @Product application/json
// @Param data body dto.FinanceTypeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/finance-type [post]
// @Security Bearer
func (e FinanceType) Insert(c *gin.Context) {
	req := dto.FinanceTypeInsertReq{}
	s := service.FinanceType{}
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
		e.Error(500, err, fmt.Sprintf("创建收支类型失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改收支类型
// @Summary 修改收支类型
// @Description 修改收支类型
// @Tags 收支类型
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FinanceTypeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/finance-type/{id} [put]
// @Security Bearer
func (e FinanceType) Update(c *gin.Context) {
	req := dto.FinanceTypeUpdateReq{}
	s := service.FinanceType{}
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

	err = s.Update(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改收支类型失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除收支类型
// @Summary 删除收支类型
// @Description 删除收支类型
// @Tags 收支类型
// @Param data body dto.FinanceTypeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/finance-type [delete]
// @Security Bearer
func (e FinanceType) Delete(c *gin.Context) {
	s := service.FinanceType{}
	req := dto.FinanceTypeDeleteReq{}
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

	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除收支类型失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Get2Tree 收支明细 左侧部门树
func (e FinanceType) Get2Tree(c *gin.Context) {
	s := service.FinanceType{}
	req := dto.FinanceTypeGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list := make([]dto.TypeLabel, 0)
	list, err = s.SetTypeTree(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(list, "")
}
