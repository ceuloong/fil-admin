package router

import (
	"fil-admin/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

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
