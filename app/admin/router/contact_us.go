package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth"

	"fil-admin/app/admin/apis"
	"fil-admin/common/middleware"
	"fil-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerContactUsRouter)
}

// registerContactUsRouter
func registerContactUsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.ContactUs{}
	r := v1.Group("/contact-us").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}