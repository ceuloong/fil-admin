package router

import (
	"fil-admin/common/middleware"
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"fil-admin/app/filpool/apis"
	"fil-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerNodeChartRouter)
}

// registerFilNodesRouter
func registerNodeChartRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.NodesChart{}
	r := v1.Group("/nodes-chart").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/ran-list", actions.PermissionAction(), api.RanList)
		r.GET("/get", actions.PermissionAction(), api.Get)
		r.GET("", actions.PermissionAction(), api.GetPage)
		//r.POST("/export", api.ExportXlsx)
	}
}
