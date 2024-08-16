package router

import (
	"fil-admin/app/admin/apis"
	"fil-admin/common/middleware"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysLoginLogRouter)
}

// 需认证的路由代码
func registerSysLoginLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysLoginLog{}

	r := v1.Group("/sys-login-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
	}
}
