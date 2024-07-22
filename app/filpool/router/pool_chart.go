package router

import (
	"fil-admin/app/filpool/apis"
	"fil-admin/common/actions"
	"fil-admin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
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
		r.GET("/app-chart", actions.PermissionAction(), api.AppChartList)
	}
}
