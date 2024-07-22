package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get(jwt.JwtPayloadKey)
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func GetDeptId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["deptId"] != nil {
		return int((data["deptId"]).(float64))
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetDeptId 缺少 deptId")
	return 0
}

func GetDeptName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["deptName"] != nil {
		return (data["deptName"]).(string)
	}
	fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetDeptName 缺少 deptName")
	return ""
}
