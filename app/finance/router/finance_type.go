package router

import (
	"fil-admin/app/finance/apis"

	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"

	"fil-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFinanceTypeRouter)
}

// registerFinanceTypeRouter
func registerFinanceTypeRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.FinanceType{}
	r := v1.Group("/finance-type").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}

	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/typeTree", api.Get2Tree)
	}
}
