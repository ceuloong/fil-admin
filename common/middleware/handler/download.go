package handler

import (
	"encoding/json"
	"fil-admin/common/redis"

	"github.com/ceuloong/fil-admin-core/sdk/pkg/response"

	"github.com/360EntSecGroup-Skylar/excelize"
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
		//log.Fatal("Open file failed.")
		response.Error(c, 500, err, "Open file failed.")
		return
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

type TickerJson struct {
	NewlyPrice    float64 `json:"newlyPrice"`
	PercentChange float64 `json:"percentChange"`
	FlowTotal     string  `json:"flowTotal"`
	CnyRate       float64 `json:"cnyRate"`
	CnyPrice      float64 `json:"cnyPrice"`
}

func FilPrice(c *gin.Context) {
	lastJson, _ := redis.GetRedis("ticker_json")
	var ticker TickerJson
	json.Unmarshal([]byte(lastJson), &ticker)
	response.OK(c, ticker, "success")
}
