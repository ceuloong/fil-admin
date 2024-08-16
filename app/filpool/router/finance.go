package router

import (
	"fil-admin/common/middleware"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"

	"fil-admin/app/filpool/apis"
	"fil-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFinanceRouter)
}

// registerFinanceRouter
func registerFinanceRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Finance{}
	r := v1.Group("/finance").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetFinance)
		r.GET("/blockstats", actions.PermissionAction(), api.BlockStats)
	}
}
