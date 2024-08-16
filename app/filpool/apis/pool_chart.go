package apis

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	"fil-admin/common/middleware"
	"fil-admin/utils"
)

type FilPoolChart struct {
	api.Api
}

// ChartList 获取FilPoolChart列表
// @Summary 获取FilPoolChart列表
// @Description 获取FilPoolChart列表
// @Tags FilPoolChart
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FilPoolChart}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-pool [get]
// @Security Bearer
func (e FilPoolChart) ChartList(c *gin.Context) {
	req := dto.FilPoolChartGetPageReq{}
	s := service.FilPoolChart{}
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

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = deptId
		}
	}

	p := actions.GetPermissionFromContext(c)
	sourceList := make([]models.FilPoolChart, 0)
	err = s.GetPage(&req, p, &sourceList)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilPoolChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	list := e.ListAddZero(sourceList)

	m := make(map[string][]models.BarChart)
	//var charts [100][2]string
	// 矿池算力数据
	var barData []models.BarChart
	// 控制地址余额数据
	var barData2 []models.BarChart
	// 矿池余额数据
	var balanceData []models.BarChart
	// 矿池扇区质押数据
	var sectorPledgeData []models.BarChart
	// 矿池奖励数据
	var rewardData []models.BarChart
	for i := len(list) - 12; i < len(list); i++ {
		f, _ := list[i].QualityAdjPower.Float64()
		cf, _ := list[i].ControlBalance.Float64()
		barData = append(barData, models.BarChart{
			X: list[i].LastTime.Format("1-02 15"),
			Y: f,
		})
		barData2 = append(barData2, models.BarChart{
			X: list[i].LastTime.Format("1-02 15"),
			Y: cf,
		})
		bf, _ := list[i].Balance.Float64()
		balanceData = append(balanceData, models.BarChart{
			X: list[i].LastTime.Format("1-02 15"),
			Y: bf,
		})
		sf, _ := list[i].SectorPledgeBalance.Float64()
		sectorPledgeData = append(sectorPledgeData, models.BarChart{
			X: list[i].LastTime.Format("1-02 15"),
			Y: sf,
		})
		rf, _ := list[i].RewardValue.Float64()
		rewardData = append(rewardData, models.BarChart{
			X: list[i].LastTime.Format("1-02 15"),
			Y: rf,
		})

	}
	m["barData"] = barData
	m["barData2"] = barData2
	m["balanceData"] = balanceData
	m["sectorPledgeData"] = sectorPledgeData
	m["rewardData"] = rewardData

	e.OK(m, "查询成功")
}

// AppChartList 获取FilPoolChart列表
// @Summary 获取FilPoolChart列表
// @Description 获取FilPoolChart列表
// @Tags FilPoolChart
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FilPoolChart}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-pool [get]
// @Security Bearer
func (e FilPoolChart) AppChartList(c *gin.Context) {
	req := dto.FilPoolChartGetPageReq{}
	s := service.FilPoolChart{}
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

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = deptId
		}
	}
	p := actions.GetPermissionFromContext(c)
	sourceList := make([]models.FilPoolChart, 0)
	err = s.GetPage(&req, p, &sourceList)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilPoolChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	list := e.ListAddZero(sourceList)

	m := make(map[string][]models.AppBarChart)
	var barData []models.AppBarChart
	for i := 0; i < len(list); i++ {
		f, _ := list[i].QualityAdjPower.Float64()
		barData = append(barData, models.AppBarChart{
			X: list[i].LastTime.Unix(),
			Y: f,
		})

	}
	m["barData"] = barData

	e.OK(m, "查询成功")
}

// Get 获取FilPoolChart
// @Summary 获取FilPoolChart
// @Description 获取FilPoolChart
// @Tags FilPoolChart
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FilPoolChart} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-pool/{id} [get]
// @Security Bearer
func (e FilPoolChart) Get(c *gin.Context) {
	req := dto.FilPoolChartGetReq{}
	s := service.FilPoolChart{}
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
	var poolChart models.FilPoolChart

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = deptId
		}
	}

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &poolChart)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilPoolChart失败，\r\n失败信息 %s", err.Error()))
		return
	}
	poolIndex := models.PoolIndex{
		AvailableBalance:    poolChart.AvailableBalance,
		Balance:             poolChart.Balance,
		SectorPledgeBalance: poolChart.SectorPledgeBalance,
		VestingFunds:        poolChart.VestingFunds,
		LastTime:            poolChart.LastTime,
		RewardValue:         poolChart.RewardValue,
		QualityAdjPower:     poolChart.QualityAdjPower,
		PowerUnit:           poolChart.PowerUnit,
		PowerPoint:          poolChart.PowerPoint,
		ControlBalance:      poolChart.ControlBalance,
	}

	var object models.FilPoolChart
	// 前一天
	object = e.GetOne(poolChart.LastTime.Add(-time.Hour*24), p, c)
	poolIndex.DayIncrease = poolIndex.QualityAdjPower.Sub(object.QualityAdjPower) //.Div(object.QualityAdjPower).RoundDown(2)
	if poolIndex.DayIncrease.LessThan(decimal.Zero) {
		poolIndex.DayTop = "bottom"
	} else {
		poolIndex.DayTop = "top"
	}

	// 上一周
	object = e.GetOne(poolChart.LastTime.Add(-time.Hour*24*7), p, c)
	poolIndex.WeekIncrease = poolIndex.QualityAdjPower.Sub(object.QualityAdjPower) //.Div(object.QualityAdjPower).RoundDown(2)
	if poolIndex.WeekIncrease.LessThan(decimal.Zero) {
		poolIndex.WeekTop = "bottom"
	} else {
		poolIndex.WeekTop = "top"
	}

	// 上个月
	object = e.GetOne(poolChart.LastTime.Add(-time.Hour*24*30), p, c)
	poolIndex.MonthAvg = poolIndex.RewardValue.Sub(object.RewardValue).Div(decimal.NewFromInt(30)).RoundDown(2)

	e.OK(poolIndex, "查询成功")
}

func (e FilPoolChart) GetOne(date time.Time, p *actions.DataPermission, c *gin.Context) models.FilPoolChart {
	req := dto.FilPoolChartGetReq{
		LastTime: date,
	}
	var object models.FilPoolChart
	s := service.FilPoolChart{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
	}

	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilPoolChart失败，\r\n失败信息 %s", err.Error()))
	}
	return object
}

// ChartAddZero 为了解决前端图表数据不全问题，添加前30天的数据
func (e FilPoolChart) ChartAddZero(list []models.FilPoolChart) []models.AppBarChart {
	barData := make([]models.AppBarChart, 0)
	now := time.Now()
	lastDay := utils.SetTime(now.AddDate(0, 0, -29), 12)
	// 先初始化30个点
	for i := 0; i < 30; i++ {
		barData = append(barData, models.AppBarChart{
			X: lastDay.AddDate(0, 0, i).Unix(),
			Y: 0,
		})
	}

	for index, v := range barData {
		for _, li := range list {
			t1 := time.Unix(v.X, 0)
			t2 := li.LastTime
			if t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day() {
				f, _ := li.QualityAdjPower.Float64()
				barData[index].Y = f
				break
			}
		}
	}

	return barData
}

// ChartAddZero 为了解决前端图表数据不全问题，添加前30天的数据
func (e FilPoolChart) ListAddZero(list []models.FilPoolChart) []models.FilPoolChart {
	newList := make([]models.FilPoolChart, 0)
	now := time.Now()
	lastDay := utils.SetTime(now.AddDate(0, 0, -29), 0)
	// 先初始化30个点
	for i := 0; i < 30; i++ {
		newList = append(newList, models.FilPoolChart{
			LastTime: lastDay.AddDate(0, 0, i),
		})
	}

	for i := 0; i < len(newList); i++ {
		for _, li := range list {
			t1 := time.Unix(newList[i].LastTime.Unix(), 0)
			t2 := li.LastTime
			if t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day() {
				newList[i] = li
				break
			}
		}
	}
	return newList
}

// ChartAddZero 为了解决前端图表数据不全问题，添加前一天的数据
/**
 * 数据格式：X:12:00，Y:算力
 *
 */
func (e FilPoolChart) DayChartAddZero(list []models.FilPoolChart) []models.BarChart {
	// 为了解决前端图表数据不全问题，添加前一天的数据
	// 1. 获取前一天的数据
	barData := make([]models.BarChart, 0)
	now := time.Now()
	lastDay := utils.SetTime(now.AddDate(0, 0, -1), now.Hour())
	// 先初始化24个点
	for i := 0; i < 24; i++ {
		barData = append(barData, models.BarChart{
			X: strconv.Itoa(lastDay.Add(time.Hour*time.Duration(i)).Hour()) + ":00",
			Y: 0,
		})
	}

	for c, li := range list {
		log.Printf("index:%d\n", c)
		for index, v := range barData {
			timeStr := strconv.Itoa(li.LastTime.Hour()) + ":00"
			if v.X == timeStr {
				f, _ := li.QualityAdjPower.Float64()
				barData[index].Y = f
			}

		}
	}

	return barData
}
