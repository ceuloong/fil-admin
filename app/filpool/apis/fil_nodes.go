package apis

import (
	"fil-admin/common/middleware"
	"fil-admin/common/middleware/handler"
	"fil-admin/common/redis"
	"fil-admin/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
)

type FilNodes struct {
	api.Api
}

// GetPage 获取FilNodes列表
// @Summary 获取FilNodes列表
// @Description 获取FilNodes列表
// @Tags FilNodes
// @Param node query string false "账户名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FilNodes}} "{"code": 200, "data": [...]}"
// @Router /api/v1/filpool-nodes [get]
// @Security Bearer
func (e FilNodes) GetPage(c *gin.Context) {
	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	list := make([]models.FilNodes, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = fmt.Sprintf("/%d/", deptId)
		}
	}
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	newList := make([]models.FilNodes, 0)
	poolIndex := new(models.NodesTotal)
	poolIndex.RoleId = user.GetRoleId(c)
	var luckyCount int64
	for _, filNodes := range list {
		if filNodes.EndTime.Before(time.Now()) {
			filNodes.Tag = "gray"
		} else if filNodes.EndTime.Before(time.Now().AddDate(0, 0, 30)) {
			filNodes.Tag = "red"
		} else {
			filNodes.Tag = "green"
		}
		newList = append(newList, filNodes)

		poolIndex.AvailableBalance = poolIndex.AvailableBalance.Add(filNodes.AvailableBalance)
		poolIndex.Balance = poolIndex.Balance.Add(filNodes.Balance)
		poolIndex.SectorPledgeBalance = poolIndex.SectorPledgeBalance.Add(filNodes.SectorPledgeBalance)
		poolIndex.VestingFunds = poolIndex.VestingFunds.Add(filNodes.VestingFunds)
		poolIndex.RewardValue = poolIndex.RewardValue.Add(filNodes.RewardValue)
		poolIndex.QualityAdjPower = poolIndex.QualityAdjPower.Add(filNodes.QualityAdjPower)
		poolIndex.PowerPoint = poolIndex.PowerPoint.Add(filNodes.PowerPoint)
		poolIndex.QualityAdjPowerDelta24h = poolIndex.QualityAdjPowerDelta24h.Add(filNodes.QualityAdjPowerDelta24h)
		poolIndex.BlocksMined24h = poolIndex.BlocksMined24h + filNodes.BlocksMined24h
		poolIndex.TotalRewards24h = poolIndex.TotalRewards24h.Add(filNodes.TotalRewards24h)
		if filNodes.QualityAdjPower.GreaterThan(decimal.Zero) {
			luckyCount++
		}
		poolIndex.LuckyValue24h = poolIndex.LuckyValue24h.Add(filNodes.LuckyValue24h)
	}
	if luckyCount > 0 {
		poolIndex.LuckyValue24h = poolIndex.LuckyValue24h.Div(decimal.NewFromInt(luckyCount))
	}
	poolIndex.NodesList = &newList

	e.PageOK(poolIndex, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e FilNodes) ChartList(c *gin.Context) {
	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	list := make([]models.FilNodes, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = fmt.Sprintf("/%d/", deptId)
		}
	}
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	newList := make([]handler.FilNodes, 0)
	var nodes []string
	for _, filNodes := range list {
		nodes = append(nodes, filNodes.Node)
		f := handler.FilNodes{}
		newList = append(newList, f.Generate(filNodes))
	}
	// TODO 从NodesChart里 in 查询返回map集合，node做为键
	ListWithChart := make([]handler.FilNodes, 0)
	if len(nodes) > 0 {
		ne := NodesChart{}
		lastTime := time.Now().Add(-time.Hour * 24 * 30)
		m := ne.GetList(lastTime, nodes, c)
		for _, node := range newList {
			if charts := m[node.Node]; charts != nil {
				node.ChartList = &charts
			} else {
				node.ChartList = new([]handler.NodesChart)
			}
			ListWithChart = append(ListWithChart, node)
		}
	}

	e.PageOK(ListWithChart, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e FilNodes) RankList(c *gin.Context) {
	s := service.FilNodes{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.FilNodes, 0)

	err = s.RankList(p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}
	m := make(map[string][]models.RankList)
	var rankList []models.RankList
	for i := 0; i < 10; i++ {
		f, _ := list[i].QualityAdjPower.Float64()
		rankList = append(rankList, models.RankList{
			Name:  list[i].Node,
			Total: f,
		})
	}
	m["rankList"] = rankList

	err = s.ControlList(p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}
	var controlList []models.RankList
	for i := 0; i < 10; i++ {
		cf, _ := list[i].ControlBalance.Float64()
		controlList = append(controlList, models.RankList{
			Name:  list[i].Node,
			Total: cf,
		})
	}
	m["controlList"] = controlList

	e.OK(m, "查询成功")
}

// Get 获取FilNodes
// @Summary 获取FilNodes
// @Description 获取FilNodes
// @Tags FilNodes
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FilNodes} "{"code": 200, "data": [...]}"
// @Router /api/v1/filpool-nodes/{id} [get]
// @Security Bearer
func (e FilNodes) Get(c *gin.Context) {
	req := dto.FilNodesGetReq{}
	s := service.FilNodes{}
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
	var object models.FilNodes

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建FilNodes
// @Summary 创建FilNodes
// @Description 创建FilNodes
// @Tags FilNodes
// @Accept application/json
// @Product application/json
// @Param data body dto.FilNodesInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/filpool-nodes [post]
// @Security Bearer
func (e FilNodes) Insert(c *gin.Context) {
	req := dto.FilNodesInsertReq{}
	s := service.FilNodes{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改FilNodes
// @Summary 修改FilNodes
// @Description 修改FilNodes
// @Tags FilNodes
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FilNodesUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/filpool-nodes/{id} [put]
// @Security Bearer
func (e FilNodes) Update(c *gin.Context) {
	req := dto.FilNodesUpdateReq{}
	s := service.FilNodes{}
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
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除FilNodes
// @Summary 删除FilNodes
// @Description 删除FilNodes
// @Tags FilNodes
// @Param data body dto.FilNodesDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/filpool-nodes [delete]
// @Security Bearer
func (e FilNodes) Delete(c *gin.Context) {
	s := service.FilNodes{}
	req := dto.FilNodesDeleteReq{}
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

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e FilNodes) ExportXlsx(c *gin.Context) {
	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	err = s.ExportXlsx(&req, c)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("导出节点信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK("", "创建成功")
}

func (e FilNodes) AppPage(c *gin.Context) {
	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	list := make([]models.FilNodes, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = fmt.Sprintf("/%d/", deptId)
		}
	}
	err = s.GetAll(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	newList := make([]handler.FilNodes, 0)
	f := handler.FilNodes{}
	for _, filNodes := range list {
		newList = append(newList, f.Generate(filNodes))
	}

	e.PageOK(newList, int(count), req.GetPageIndex(), 100, "查询成功")
}

func (e FilNodes) NodesTotal(c *gin.Context) {
	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	list := make([]models.FilNodes, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = fmt.Sprintf("/%d/", deptId)
		}
	}
	err = s.GetAll(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	poolIndex := new(models.NodesTotal)

	var luckyCount int64
	for _, filNodes := range list {
		poolIndex.AvailableBalance = poolIndex.AvailableBalance.Add(filNodes.AvailableBalance)
		poolIndex.Balance = poolIndex.Balance.Add(filNodes.Balance)
		poolIndex.SectorPledgeBalance = poolIndex.SectorPledgeBalance.Add(filNodes.SectorPledgeBalance)
		poolIndex.VestingFunds = poolIndex.VestingFunds.Add(filNodes.VestingFunds)
		poolIndex.RewardValue = poolIndex.RewardValue.Add(filNodes.RewardValue)
		poolIndex.QualityAdjPower = poolIndex.QualityAdjPower.Add(filNodes.QualityAdjPower)
		poolIndex.PowerPoint = poolIndex.PowerPoint.Add(filNodes.PowerPoint)
		poolIndex.QualityAdjPowerDelta24h = poolIndex.QualityAdjPowerDelta24h.Add(filNodes.QualityAdjPowerDelta24h)
		poolIndex.BlocksMined24h = poolIndex.BlocksMined24h + filNodes.BlocksMined24h
		poolIndex.TotalRewards24h = poolIndex.TotalRewards24h.Add(filNodes.TotalRewards24h)
		if filNodes.QualityAdjPower.GreaterThan(decimal.Zero) {
			luckyCount++
		}
		poolIndex.LuckyValue24h = poolIndex.LuckyValue24h.Add(filNodes.LuckyValue24h)
	}
	if luckyCount > 0 {
		poolIndex.LuckyValue24h = poolIndex.LuckyValue24h.Div(decimal.NewFromInt(luckyCount))
	}

	e.OK(poolIndex, "查询成功")
}

func (e FilNodes) UpdateTitle(c *gin.Context) {
	req := dto.FilNodesUpdateTitleReq{}
	s := service.FilNodes{}
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
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.UpdateTitle(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改title失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

/**
 * 获取矿池的财务数据
 */
func (e FilNodes) GetFinance(c *gin.Context) {

	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	list := make([]models.FilNodes, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = fmt.Sprintf("/%d/", deptId)
		}
	}
	err = s.GetAll(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	poolIndex := new(models.PoolFinance)

	for _, filNodes := range list {
		poolIndex.AvailableBalance = poolIndex.AvailableBalance.Add(filNodes.AvailableBalance)
		poolIndex.Balance = poolIndex.Balance.Add(filNodes.Balance)
		poolIndex.BlocksMined24h = poolIndex.BlocksMined24h + filNodes.BlocksMined24h
		poolIndex.TotalRewards24h = poolIndex.TotalRewards24h.Add(filNodes.TotalRewards24h)
	}

	priceStr, _ := redis.GetRedis("ticker")
	poolIndex.NewlyPrice = utils.DecimalValue(priceStr)

	e.OK(poolIndex, "查询成功")
}

/**
 * 获取矿池的报块数据
 */
func (e FilNodes) BlockStats(c *gin.Context) {

	req := dto.FilNodesGetPageReq{}
	s := service.FilNodes{}
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
	list := make([]models.FilNodes, 0)
	var count int64

	if user.GetRoleName(c) != "admin" && user.GetRoleName(c) != "系统管理员" {
		deptId := middleware.GetDeptId(c)
		if deptId > 0 {
			req.DeptId = fmt.Sprintf("/%d/", deptId)
		}
	}
	err = s.GetAll(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取FilNodes失败，\r\n失败信息 %s", err.Error()))
		return
	}

	var nodes []string
	for _, li := range list {
		nodes = append(nodes, li.Node)
	}

	now := time.Now()
	lastDay := utils.SetTime(now.AddDate(0, 0, -1), now.Hour())

	blockCharts := make([]models.BlockStats, 0)

	err = s.SumBlockStats(nodes, lastDay, &blockCharts)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取报块数据失败，\r\n失败信息 %s", err.Error()))
		return
	}
	poolBlockStats := e.ChartAddZero(blockCharts)

	e.OK(poolBlockStats, "查询成功")
}

func (e FilNodes) ChartAddZero(list []models.BlockStats) []models.PoolBlockStats {
	// 为了解决前端图表数据不全问题，添加前一天的数据
	// 1. 获取前一天的数据
	newList := make([]models.PoolBlockStats, 0)
	if list == nil {
		// 全部补0
		now := time.Now()
		lastDay := utils.SetTime(now.AddDate(0, 0, -1), now.Hour())
		for i := 0; i < 24; i++ {
			newHour := lastDay.Add(time.Hour * time.Duration(i)).Hour()
			newList = append(newList, models.PoolBlockStats{
				HeightTime:            lastDay.Add(time.Hour * time.Duration(i)),
				HeightTimeStr:         strconv.Itoa(newHour) + ":00",
				BlocksGrowth:          0,
				BlocksRewardGrowthFil: "0",
			})
		}
	} else if len(list) <= 24 {
		count := len(list)
		last := list[len(list)-1]
		lastTime := last.HeightTime
		for _, li := range list {
			newList = append(newList, models.PoolBlockStats{
				HeightTime:            li.HeightTime,
				HeightTimeStr:         li.HeightTimeStr,
				BlocksGrowth:          li.BlocksGrowth,
				BlocksRewardGrowthFil: li.BlocksRewardGrowthFil,
			})
		}
		for i := 1; i <= (24 - count); i++ {
			newList = append(newList, models.PoolBlockStats{
				HeightTime:            lastTime.Add(time.Hour * time.Duration(i)),
				HeightTimeStr:         strconv.Itoa(lastTime.Add(time.Hour*time.Duration(i)).Hour()) + ":00",
				BlocksGrowth:          0,
				BlocksRewardGrowthFil: "0",
			})
		}
	}

	return newList
}