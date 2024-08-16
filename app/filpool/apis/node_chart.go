package apis

import (
	"fil-admin/common/middleware/handler"
	"fmt"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
)

type NodesChart struct {
	api.Api
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
