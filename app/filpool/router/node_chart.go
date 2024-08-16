package router

import (
	"fil-admin/common/middleware"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"

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
