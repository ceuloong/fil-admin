package handler

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/ceuloong/fil-admin-core/logger"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileName := c.Query("fileName")
	if len(fileName) == 0 {
		fileName = "fil_nodes"
	}
	// 创建一个新的 Excel 文件
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal("Open file failed.")
	}

	// 设置响应头，告诉浏览器这是一个 Excel 文件
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+fileName+".xlsx")

	// 将 Excel 文件写入响应体，返回给浏览器进行下载
	file.Write(c.Writer)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
