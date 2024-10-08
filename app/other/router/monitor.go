package router

import (
	"net/http"

	"github.com/ceuloong/fil-admin-core/tools/transfer"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerMonitorRouter)
}

// 需认证的路由代码
func registerMonitorRouter(v1 *gin.RouterGroup) {
	v1.GET("/metrics", transfer.Handler(promhttp.Handler()))
	//健康检查
	v1.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

}
