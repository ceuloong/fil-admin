package router

import (
	"fil-admin/app/finance/apis"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"

	"fil-admin/common/actions"
	"fil-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFinanceCashRouter)
}

// registerFinanceCashRouter
func registerFinanceCashRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FinanceCash{}
	r := v1.Group("/finance-cash").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		//r.DELETE("", api.Delete)
		r.POST("/export", api.ExportXlsx)
	}
}
