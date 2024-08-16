package router

import (
	"fil-admin/app/other/apis"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFileRouter)
}

// 需认证的路由代码
func registerFileRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	var api = apis.File{}
	r := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r.POST("/public/uploadFile", api.UploadFile)
	}
}
