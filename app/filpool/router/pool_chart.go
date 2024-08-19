package router

import (
	"fil-admin/app/filpool/apis"
	"fil-admin/common/actions"
	"fil-admin/common/middleware"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFilPoolRouter)
}

// registerFilNodesRouter
func registerFilPoolRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FilPoolChart{}
	r := v1.Group("/fil-pool").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.ChartList)
		r.GET("/get", actions.PermissionAction(), api.Get)
		r.GET("/app-get", actions.PermissionAction(), api.AppGet)
		r.GET("/app-chart", actions.PermissionAction(), api.AppChartList)
	}
}
