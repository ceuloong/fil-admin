package router

import (
	"fil-admin/app/finance/apis"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"

	"fil-admin/common/actions"
	"fil-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFinanceCoinRouter)
}

// registerFinanceCoinRouter
func registerFinanceCoinRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FinanceCoin{}
	r := v1.Group("/finance-coin").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}
