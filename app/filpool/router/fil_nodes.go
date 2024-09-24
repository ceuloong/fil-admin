package router

import (
	"fil-admin/common/middleware"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"

	"fil-admin/app/filpool/apis"
	"fil-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFilNodesRouter)
}

// registerFilNodesRouter
func registerFilNodesRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FilNodes{}
	r := v1.Group("/fil-nodes").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.POST("/export", api.ExportXlsx)
		r.POST("/rank-list", api.RankList)
	}

	r1 := v1.Group("/nodes-app").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.GET("", actions.PermissionAction(), api.AppPage)
		r1.GET("/chart", actions.PermissionAction(), api.ChartList)
		r1.GET("/total", actions.PermissionAction(), api.NodesTotal)
		r1.PUT("/:id", actions.PermissionAction(), api.UpdateTitle)
		r1.GET("/:id", actions.PermissionAction(), api.Get)
		r1.GET("/finance", actions.PermissionAction(), api.GetFinance)
		r1.GET("/blockstats", actions.PermissionAction(), api.BlockStats)
		r1.GET("/sectors", actions.PermissionAction(), api.GetSectors)
	}
}
