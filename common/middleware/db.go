package middleware

import (
	"github.com/ceuloong/fil-admin-core/sdk"
	"github.com/gin-gonic/gin"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", sdk.Runtime.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}
