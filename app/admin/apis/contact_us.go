package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"

	"fil-admin/app/admin/models"
	"fil-admin/app/admin/service"
	"fil-admin/app/admin/service/dto"
	"fil-admin/common/actions"
)

type ContactUs struct {
	api.Api
}

// GetPage 获取ContactUs列表
// @Summary 获取ContactUs列表
// @Description 获取ContactUs列表
// @Tags ContactUs
// @Param email query string false "邮箱"
// @Param subject query string false "主题"
// @Param message query string false "内容"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ContactUs}} "{"code": 200, "data": [...]}"
// @Router /api/v1/contact-us [get]
// @Security Bearer
func (e ContactUs) GetPage(c *gin.Context) {
    req := dto.ContactUsGetPageReq{}
    s := service.ContactUs{}
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
	list := make([]models.ContactUs, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ContactUs失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ContactUs
// @Summary 获取ContactUs
// @Description 获取ContactUs
// @Tags ContactUs
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ContactUs} "{"code": 200, "data": [...]}"
// @Router /api/v1/contact-us/{id} [get]
// @Security Bearer
func (e ContactUs) Get(c *gin.Context) {
	req := dto.ContactUsGetReq{}
	s := service.ContactUs{}
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
	var object models.ContactUs

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ContactUs失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建ContactUs
// @Summary 创建ContactUs
// @Description 创建ContactUs
// @Tags ContactUs
// @Accept application/json
// @Product application/json
// @Param data body dto.ContactUsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/contact-us [post]
// @Security Bearer
func (e ContactUs) Insert(c *gin.Context) {
    req := dto.ContactUsInsertReq{}
    s := service.ContactUs{}
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
		e.Error(500, err, fmt.Sprintf("创建ContactUs失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改ContactUs
// @Summary 修改ContactUs
// @Description 修改ContactUs
// @Tags ContactUs
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ContactUsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/contact-us/{id} [put]
// @Security Bearer
func (e ContactUs) Update(c *gin.Context) {
    req := dto.ContactUsUpdateReq{}
    s := service.ContactUs{}
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
		e.Error(500, err, fmt.Sprintf("修改ContactUs失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除ContactUs
// @Summary 删除ContactUs
// @Description 删除ContactUs
// @Tags ContactUs
// @Param data body dto.ContactUsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/contact-us [delete]
// @Security Bearer
func (e ContactUs) Delete(c *gin.Context) {
    s := service.ContactUs{}
    req := dto.ContactUsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ContactUs失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
