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
)

type Block struct {
	api.Api
}

// GetPage 获取Block列表
// @Summary 获取Block列表
// @Description 获取Block列表
// @Tags Block
// @Param height query int false "高度"
// @Param node query string false "节点"
// @Param message query string false "区块哈希"
// @Param status query string false "1正常   2孤块"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Block}} "{"code": 200, "data": [...]}"
// @Router /api/v1/block [get]
// @Security Bearer
func (e Block) GetPage(c *gin.Context) {
	req := dto.BlockGetPageReq{}
	s := service.Block{}
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
	list := make([]models.Block, 0)
	var count int64
	if req.Status == "" {
		req.Status = "2"
	}

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Block失败，\r\n失败信息 %s", err.Error()))
		return
	}

	newList := make([]models.BlockShow, 0)
	for _, b := range list {
		newList = append(newList, b.Generate())
	}

	e.PageOK(newList, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Block
// @Summary 获取Block
// @Description 获取Block
// @Tags Block
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Block} "{"code": 200, "data": [...]}"
// @Router /api/v1/block/{id} [get]
// @Security Bearer
func (e Block) Get(c *gin.Context) {
	req := dto.BlockGetReq{}
	s := service.Block{}
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
	var object models.Block

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Block失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Block
// @Summary 创建Block
// @Description 创建Block
// @Tags Block
// @Accept application/json
// @Product application/json
// @Param data body dto.BlockInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/block [post]
// @Security Bearer
func (e Block) Insert(c *gin.Context) {
	req := dto.BlockInsertReq{}
	s := service.Block{}
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
		e.Error(500, err, fmt.Sprintf("创建Block失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Block
// @Summary 修改Block
// @Description 修改Block
// @Tags Block
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BlockUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/block/{id} [put]
// @Security Bearer
func (e Block) Update(c *gin.Context) {
	req := dto.BlockUpdateReq{}
	s := service.Block{}
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
		e.Error(500, err, fmt.Sprintf("修改Block失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Block
// @Summary 删除Block
// @Description 删除Block
// @Tags Block
// @Param data body dto.BlockDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/block [delete]
// @Security Bearer
func (e Block) Delete(c *gin.Context) {
	s := service.Block{}
	req := dto.BlockDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除Block失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
