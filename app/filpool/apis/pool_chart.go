package apis

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.FilPoolChart, 0)

	err = s.GetPage(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilPoolChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	m := make(map[string][]models.BarChart)
	//var charts [100][2]string
	var barData []models.BarChart
	var barData2 []models.BarChart
	var balanceData []models.BarChart
	var sectorPledgeData []models.BarChart
	var rewardData []models.BarChart
	for i := 12; i >= 0; i-- {
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.FilPoolChart, 0)

	err = s.GetPage(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilPoolChart失败，\r\n失败信息 %s", err.Error()))
		return
	}

	m := make(map[string][]models.AppBarChart)
	var barData []models.AppBarChart
	for i := 29; i >= 0; i-- {
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