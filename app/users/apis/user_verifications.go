package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"

	"fil-admin/app/users/models"
	"fil-admin/app/users/service"
	"fil-admin/app/users/service/dto"
	"fil-admin/common/actions"
)

type UserVerifications struct {
	api.Api
}

// GetPage 获取UserVerifications列表
// @Summary 获取UserVerifications列表
// @Description 获取UserVerifications列表
// @Tags UserVerifications
// @Param userId query string false "用户Id"
// @Param realName query string false "真实姓名"
// @Param idNumber query string false "证件号码"
// @Param status query string false "认证状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.UserVerifications}} "{"code": 200, "data": [...]}"
// @Router /api/v1/user-verifications [get]
// @Security Bearer
func (e UserVerifications) GetPage(c *gin.Context) {
    req := dto.UserVerificationsGetPageReq{}
    s := service.UserVerifications{}
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
	list := make([]models.UserVerifications, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取UserVerifications失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取UserVerifications
// @Summary 获取UserVerifications
// @Description 获取UserVerifications
// @Tags UserVerifications
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.UserVerifications} "{"code": 200, "data": [...]}"
// @Router /api/v1/user-verifications/{id} [get]
// @Security Bearer
func (e UserVerifications) Get(c *gin.Context) {
	req := dto.UserVerificationsGetReq{}
	s := service.UserVerifications{}
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
	var object models.UserVerifications

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取UserVerifications失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建UserVerifications
// @Summary 创建UserVerifications
// @Description 创建UserVerifications
// @Tags UserVerifications
// @Accept application/json
// @Product application/json
// @Param data body dto.UserVerificationsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/user-verifications [post]
// @Security Bearer
func (e UserVerifications) Insert(c *gin.Context) {
    req := dto.UserVerificationsInsertReq{}
    s := service.UserVerifications{}
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
		e.Error(500, err, fmt.Sprintf("创建UserVerifications失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改UserVerifications
// @Summary 修改UserVerifications
// @Description 修改UserVerifications
// @Tags UserVerifications
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.UserVerificationsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/user-verifications/{id} [put]
// @Security Bearer
func (e UserVerifications) Update(c *gin.Context) {
    req := dto.UserVerificationsUpdateReq{}
    s := service.UserVerifications{}
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
		e.Error(500, err, fmt.Sprintf("修改UserVerifications失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除UserVerifications
// @Summary 删除UserVerifications
// @Description 删除UserVerifications
// @Tags UserVerifications
// @Param data body dto.UserVerificationsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/user-verifications [delete]
// @Security Bearer
func (e UserVerifications) Delete(c *gin.Context) {
    s := service.UserVerifications{}
    req := dto.UserVerificationsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除UserVerifications失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
