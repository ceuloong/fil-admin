package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
)

func GetDeptId(c *gin.Context) int {
	deptId, err := user.ExtractClaims(c).Int("deptId")
	if err != nil {
		fmt.Println(pkg.GetCurrentTimeStr() + " [WARING] " + c.Request.Method + " " + c.Request.URL.Path + " GetDeptId 缺少 deptid error: " + err.Error())
		return 0
	}

	return deptId
}

func GetDeptName(c *gin.Context) string {
	return user.ExtractClaims(c).String("deptName")
}
