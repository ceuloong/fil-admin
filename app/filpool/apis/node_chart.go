package apis

import (
	"encoding/base64"
	"fil-admin/common/middleware/handler"
	"fmt"
	"path/filepath"
	"reflect"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
)

type NodesChart struct {
	api.Api
}

// GetSnapshotWithFilNodes 获取快照数据并关联filNodes信息
func (e NodesChart) GetSnapshotWithFilNodes(c *gin.Context) {
	req := dto.NodeChartGetPageReq{}
	s := service.NodesChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if req.LastTime == "" {
		req.LastTime = time.Now().Add(-time.Hour).Format("2006-01-02 15:00")
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.NodesChartWithFilNodes, 0)

	err = s.GetSnapshotWithFilNodes(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取快照数据失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(list, "查询成功")
}

func (e NodesChart) GetPage(c *gin.Context) {
	req := dto.NodeChartGetPageReq{}
	s := service.NodesChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.NodesChart, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取NodeChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// GetStats 获取节点统计信息
// @Summary 获取节点统计信息
// @Description 获取节点统计总算力、总余额等统计信息
// @Tags NodeChart
// @Success 200 {object} response.Response{data=models.NodeStats} "{"code": 200, "data": [...]}"
// @Router /api/v1/node-chart/stats [get]
// @Security Bearer
func (e NodesChart) GetStats(c *gin.Context) {
	req := dto.NodeChartGetPageReq{}
	s := service.NodesChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// Set default lastTime if not provided
	if req.LastTime == "" {
		req.LastTime = time.Now().Add(-time.Hour).Format("2006-01-02 15:00")
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.NodesChartWithFilNodes, 0)

	err = s.GetSnapshotWithFilNodes(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取NodeChart统计失败，\r\n失败信息 %s", err.Error()))
		return
	}

	stats := models.NodesChartWithFilNodes{}

	// Calculate totals
	for _, node := range list {
		stats.QualityAdjPower = stats.QualityAdjPower.Add(node.QualityAdjPower)
		stats.AvailableBalance = stats.AvailableBalance.Add(node.AvailableBalance)
		stats.Balance = stats.Balance.Add(node.Balance)
		stats.SectorPledgeBalance = stats.SectorPledgeBalance.Add(node.SectorPledgeBalance)
		stats.VestingFunds = stats.VestingFunds.Add(node.VestingFunds)
		stats.QualityAdjPowerDelta24h = stats.QualityAdjPowerDelta24h.Add(node.QualityAdjPowerDelta24h)
		stats.TotalRewards24h = stats.TotalRewards24h.Add(node.TotalRewards24h)
		monthRewardValue := node.RewardValue.Sub(node.LastMonthRewardValue)
		stats.LastMonthRewardValue = stats.LastMonthRewardValue.Add(monthRewardValue)
		stats.RealRewardValueMonth = stats.RealRewardValueMonth.Add(monthRewardValue.Mul(node.DistributePoint)).Round(4)
	}

	e.OK(stats, "统计查询成功")
}

// RanList 获取NodeChart列表
// @Summary 获取NodeChart列表
// @Description 获取NodeChart列表
// @Tags NodeChart
// @Success 200 {object} response.Response{data=response.Page{list=[]models.NodeChart}} "{"code": 200, "data": [...]}"
// @Router /api/v1/node-chart/ran-list [get]
// @Security Bearer
func (e NodesChart) RanList(c *gin.Context) {
	req := dto.NodeChartGetReq{}
	s := service.NodesChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.NodesChart, 0)

	err = s.GetList(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取NodeChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	m := make(map[string][]models.RankList)
	//var charts [100][2]string
	var barData []models.RankList
	for i := 12; i > 0; i-- {
		f, _ := list[i].QualityAdjPower.Float64()
		barData = append(barData, models.RankList{
			Name:  list[i].Node,
			Total: f,
		})
	}
	m["barData"] = barData

	e.OK(m, "查询成功")
}

// Get 获取NodeChart
// @Summary 获取NodeChart
// @Description 获取NodeChart
// @Tags NodeChart
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.NodeChart} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-pool/{id} [get]
// @Security Bearer
func (e NodesChart) Get(c *gin.Context) {
	req := dto.NodeChartGetReq{}
	s := service.NodesChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var list []models.NodesChart

	p := actions.GetPermissionFromContext(c)
	err = s.GetList(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取NodeChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	m := make(map[string][]models.NodeIndex)
	for _, nc := range list {
		nodeIndex := models.NodeIndex{
			AvailableBalance:    nc.AvailableBalance,
			Balance:             nc.Balance,
			SectorPledgeBalance: nc.SectorPledgeBalance,
			VestingFunds:        nc.VestingFunds,
			LastTime:            nc.LastTime,
			RewardValue:         nc.RewardValue,
			QualityAdjPower:     nc.QualityAdjPower,
			PowerUnit:           nc.PowerUnit,
			PowerPoint:          nc.PowerPoint,
			ControlBalance:      nc.ControlBalance,
		}

		m[nc.Node] = append(m[nc.Node], nodeIndex)
	}
	// 前一天
	//list = e.GetList(poolChart.LastTime.Add(-time.Hour*24), p, c)
	//poolIndex.DayIncrease = poolIndex.QualityAdjPower.Sub(object.QualityAdjPower) //.Div(object.QualityAdjPower).RoundDown(2)
	//
	//// 上一周
	//object = e.GetOne(poolChart.LastTime.Add(-time.Hour*24*7), p, c)
	//poolIndex.WeekIncrease = poolIndex.QualityAdjPower.Sub(object.QualityAdjPower) //.Div(object.QualityAdjPower).RoundDown(2)
	//
	//// 上个月
	//object = e.GetOne(poolChart.LastTime.Add(-time.Hour*24*30), p, c)
	//poolIndex.MonthAvg = poolIndex.RewardValue.Sub(object.RewardValue).Div(decimal.NewFromInt(30)).RoundDown(2)
	//
	//e.OK(poolIndex, "查询成功")
}

//func (e NodeChart) GetOne(date time.Time, p *actions.DataPermission, c *gin.Context) models.NodeChart {
//	req := dto.NodeChartGetReq{
//		LastTime: date,
//	}
//	var object models.NodeChart
//	s := service.NodeChart{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//	}
//
//	err = s.Get(&req, p, &object)
//	if err != nil {
//		e.Error(500, err, fmt.Sprintf("获取NodeChart失败，\r\n失败信息 %s", err.Error()))
//	}
//	return object
//}

func (e NodesChart) GetList(lastTime time.Time, nodes []string, c *gin.Context) map[string][]handler.NodesChart {
	req := dto.NodeChartGetReq{
		LastTime: lastTime,
		Nodes:    nodes,
	}
	var list []models.NodesChart
	s := service.NodesChart{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
	}

	p := actions.GetPermissionFromContext(c)
	//db, err := pkg.GetOrm(c)
	//s.GetChartList(db, lastTime, nodes, &list, err)
	err = s.GetList(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取NodeChart失败，\r\n失败信息 %s", err.Error()))
	}

	m := make(map[string][]handler.NodesChart)
	for i := len(list) - 1; i >= 0; i-- {
		nc := list[i]
		nChart := handler.NodesChart{
			Node:                    nc.Node,
			AvailableBalance:        nc.AvailableBalance,
			Balance:                 nc.Balance,
			SectorPledgeBalance:     nc.SectorPledgeBalance,
			VestingFunds:            nc.VestingFunds,
			RewardValue:             nc.RewardValue,
			QualityAdjPower:         nc.QualityAdjPower,
			PowerPoint:              nc.PowerPoint,
			BlocksMined24h:          nc.BlocksMined24h,
			TotalRewards24h:         nc.TotalRewards24h,
			LuckyValue24h:           nc.LuckyValue24h,
			QualityAdjPowerDelta24h: nc.QualityAdjPowerDelta24h,
			ReceiveAmount:           nc.ReceiveAmount,
			BurnAmount:              nc.BurnAmount,
			SendAmount:              nc.SendAmount,
			LastTime:                nc.LastTime,
		}

		if len(m[nc.Node]) < 100 {
			m[nc.Node] = append(m[nc.Node], nChart)
		}
	}

	return m
}

// ExportSnapshotWithFilNodes 导出快照数据并关联filNodes信息
func (e NodesChart) ExportSnapshotWithFilNodes(c *gin.Context) {
	req := dto.NodeChartGetPageReq{}
	s := service.NodesChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 读取模板文件
	templatePath := filepath.Join("static", "model.xlsx")
	f, err := excelize.OpenFile(templatePath)
	if err != nil {
		e.Logger.Error("打开模板文件失败:", err)
		e.Error(500, err, "打开模板文件失败")
		return
	}
	defer f.Close()

	if req.LastTime == "" {
		req.LastTime = time.Now().Add(-time.Hour).Format("2006-01-02 15:00")
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.NodesChartWithFilNodes, 0)

	err = s.GetSnapshotWithFilNodes(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取快照数据失败，\r\n失败信息 %s", err.Error()))
		return
	}

	list = e.ResetList(list)

	// 填充数据
	for i, node := range list {
		if err := fillRowWithReflection(f, i+1, node); err != nil {
			e.Error(500, err, "填充数据失败")
			return
		}
	}

	// 将文件保存到buffer
	buffer, err := f.WriteToBuffer()
	if err != nil {
		e.Logger.Error("写入Excel数据失败:", err)
		e.Error(500, err, "写入Excel数据失败")
		return
	}

	// 转换为base64
	base64Data := base64.StdEncoding.EncodeToString(buffer.Bytes())
	fileName := fmt.Sprintf("节点数据导出_%s.xlsx", time.Now().Format("20060102150405"))

	e.OK(map[string]interface{}{
		"file":     base64Data,
		"filename": fileName,
	}, "获取导出数据成功")
}

func (e NodesChart) ResetList(list []models.NodesChartWithFilNodes) []models.NodesChartWithFilNodes {
	newList := make([]models.NodesChartWithFilNodes, 0)
	for _, node := range list {
		node.QualityAdjPowerDelta24h = node.QualityAdjPower.Sub(node.LastQualityAdjPower)
		node.QualityAdjPowerDeltaMonth = node.QualityAdjPower.Sub(node.LastMonthQualityAdjPower)
		node.SectorPledgeBalanceDeltaMonth = node.SectorPledgeBalance.Sub(node.LastMonthSectorPledgeBalance)
		node.LastWeightedBlocks = node.WeightedBlocks - node.LastWeightedBlocks
		node.LastRewardValue = node.RewardValue.Sub(node.LastRewardValue)
		node.LastMonthWeightedBlocks = node.WeightedBlocks - node.LastMonthWeightedBlocks
		node.LastMonthRewardValue = node.RewardValue.Sub(node.LastMonthRewardValue)
		node.RealWeightedBlocks24h = decimal.NewFromInt(int64(node.WeightedBlocks - node.LastWeightedBlocks)).Mul(node.DistributePoint).Round(1)
		node.RealRewardValue24h = node.RewardValue.Sub(node.LastRewardValue).Mul(node.DistributePoint).Round(4)
		node.RealWeightedBlocksMonth = decimal.NewFromInt(int64(node.WeightedBlocks - node.LastMonthWeightedBlocks)).Mul(node.DistributePoint).Round(1)
		node.RealRewardValueMonth = node.RewardValue.Sub(node.LastMonthRewardValue).Mul(node.DistributePoint).Round(4)
		node.LastMonthReceiveAmount = node.ReceiveAmount.Sub(node.LastMonthReceiveAmount)
		node.LastMonthBurnAmount = node.BurnAmount.Sub(node.LastMonthBurnAmount)
		node.LastMonthSendAmount = node.SendAmount.Sub(node.LastMonthSendAmount)

		newList = append(newList, node)
	}
	return newList
}

func fillRowWithReflection(sheet *excelize.File, rowIndex int, data interface{}) error {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	sheetName := "data" // or whatever your sheet name is

	// Rest of the function remains similar, but use f.SetCellValue() instead:
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		//fieldName := typ.Field(i).Name

		colName := getColumnName(i)
		cellRef := fmt.Sprintf("%s%d", colName, rowIndex+1)

		// 根据字段类型直接设置值
		switch field.Kind() {
		case reflect.String:
			sheet.SetCellValue(sheetName, cellRef, field.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sheet.SetCellValue(sheetName, cellRef, field.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			sheet.SetCellValue(sheetName, cellRef, field.Uint())
		case reflect.Float32, reflect.Float64:
			sheet.SetCellValue(sheetName, cellRef, field.Float())
		case reflect.Bool:
			sheet.SetCellValue(sheetName, cellRef, field.Bool())
		case reflect.Struct:
			if t, ok := field.Interface().(time.Time); ok {
				sheet.SetCellValue(sheetName, cellRef, t)
				styleID, _ := sheet.GetCellStyle(sheetName, cellRef)
				style, _ := sheet.GetStyle(styleID)

				// 设置日期格式
				newStyleID, err := sheet.NewStyle(&excelize.Style{
					NumFmt: 22, // "m/d/yy h:mm" 格式
					Border: style.Border,
					Fill:   style.Fill,
					Font:   style.Font,
				})
				if err == nil {
					sheet.SetCellStyle(sheetName, cellRef, cellRef, newStyleID)
				}
			}
			if d, ok := field.Interface().(decimal.Decimal); ok {
				sheet.SetCellValue(sheetName, cellRef, d.String())
			}
		default:
			sheet.SetCellValue(sheetName, cellRef, field.Interface())
		}

	}

	return nil
}

// 常用的 Excel 日期格式
const (
	DateFormatDateTime = 22 // "m/d/yy h:mm"
	DateFormatDate     = 14 // "m/d/yy"
	DateFormatTime     = 21 // "h:mm:ss"
	// 自定义格式
	DateFormatCustom1 = "yyyy-mm-dd hh:mm:ss"
	DateFormatCustom2 = "yyyy年mm月dd日 hh时mm分ss秒"
)

// 设置时间格式的辅助函数
// func setTimeCell(f *excelize.File, sheetName string, cellRef string, t time.Time, format int) {
// 	// 设置单元格值
// 	f.SetCellValue(sheetName, cellRef, t)

// 	// 创建并应用样式
// 	style, _ := f.NewStyle(&excelize.Style{
// 		NumFmt: format,
// 	})

// 	f.SetCellStyle(cellRef, cellRef, cellRef, style)
// }

// 将列索引转换为 Excel 列名（A, B, C, ..., AA, AB, ...）
func getColumnName(colIndex int) string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""

	for colIndex >= 0 {
		result = string(alphabet[colIndex%26]) + result
		colIndex = colIndex/26 - 1
	}

	return result
}

// sheetName := "data"
// // 从第二行开始写入数据（假设第一行是表头）
// startRow := 2

// // 写入数据
// for i, node := range list {
// 	row := startRow + i
// 	// 按照模板的列顺序写入数据
// 	f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), node.Node)
// 	f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), node.MsigNode)
// 	f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), node.Type)
// 	f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), node.DistributePoint)
// 	f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), node.QualityAdjPower.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), node.LastQualityAdjPower.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), node.LastMonthQualityAdjPower.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), node.QualityAdjPowerDelta24h.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), node.QualityAdjPowerDeltaMonth.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), node.AvailableBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("K%d", row), node.LastAvailableBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), node.Balance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), node.LastBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), node.SectorPledgeBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("O%d", row), node.LastSectorPledgeBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("P%d", row), node.LastMonthSectorPledgeBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("Q%d", row), node.SectorPledgeBalanceDeltaMonth.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("R%d", row), node.VestingFunds.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("S%d", row), node.LastVestingFunds.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("T%d", row), node.BlocksMined24h)
// 	f.SetCellValue(sheetName, fmt.Sprintf("U%d", row), node.TotalRewards24h.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("V%d", row), node.LuckyValue24h.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("W%d", row), node.WeightedBlocks)
// 	f.SetCellValue(sheetName, fmt.Sprintf("X%d", row), node.RewardValue.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("Y%d", row), node.LastWeightedBlocks)
// 	f.SetCellValue(sheetName, fmt.Sprintf("Z%d", row), node.LastRewardValue.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AA%d", row), node.LastMonthWeightedBlocks)
// 	f.SetCellValue(sheetName, fmt.Sprintf("AB%d", row), node.LastMonthRewardValue.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AC%d", row), node.RealWeightedBlocks24h.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AD%d", row), node.RealRewardValue24h.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AE%d", row), node.RealWeightedBlocksMonth.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AF%d", row), node.RealRewardValueMonth.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AG%d", row), node.ReceiveAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AH%d", row), node.BurnAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AI%d", row), node.SendAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AJ%d", row), node.LastReceiveAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AK%d", row), node.LastBurnAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AL%d", row), node.LastSendAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AM%d", row), node.LastMonthReceiveAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AN%d", row), node.LastMonthBurnAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AO%d", row), node.LastMonthSendAmount.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AP%d", row), node.ControlBalance.String())
// 	f.SetCellValue(sheetName, fmt.Sprintf("AQ%d", row), node.SectorEffective)
// 	f.SetCellValue(sheetName, fmt.Sprintf("AR%d", row), node.SectorError)
// 	f.SetCellValue(sheetName, fmt.Sprintf("AS%d", row), node.SectorSize)
// }
