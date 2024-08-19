package apis

import (
	"fmt"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/api"
	"github.com/ceuloong/fil-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/ceuloong/fil-admin-core/sdk/pkg/response"
	"github.com/gin-gonic/gin"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	"fil-admin/common/middleware"
	"fil-admin/common/redis"
	"fil-admin/utils"
)

type Finance struct {
	api.Api
}

func (e Finance) GetFinance(c *gin.Context) {

	list := e.FindNodesList(c)

	poolIndex := new(models.PoolFinance)

	for _, filNodes := range list {
		poolIndex.AvailableBalance = poolIndex.AvailableBalance.Add(filNodes.AvailableBalance)
		poolIndex.Balance = poolIndex.Balance.Add(filNodes.Balance)
		poolIndex.SectorPledgeBalance = poolIndex.SectorPledgeBalance.Add(filNodes.SectorPledgeBalance)
		poolIndex.VestingFunds = poolIndex.VestingFunds.Add(filNodes.VestingFunds)
		poolIndex.BlocksMined24h = poolIndex.BlocksMined24h + filNodes.BlocksMined24h
		poolIndex.TotalRewards24h = poolIndex.TotalRewards24h.Add(filNodes.TotalRewards24h)
	}

	priceStr, _ := redis.GetRedis("ticker")
	poolIndex.NewlyPrice = utils.DecimalValue(priceStr)

	e.OK(poolIndex, "查询成功")
}

func (e Finance) BlockStats(c *gin.Context) {

	list := e.FindNodesList(c)
	var nodes []string
	for _, li := range list {
		nodes = append(nodes, li.Node)
	}

	now := time.Now()
	lastDay := utils.SetTime(now.AddDate(0, 0, -1), now.Hour())

	blockCharts := make([]models.BlockStats, 0)

	s := service.Finance{}
	err := s.SumBlockStats(nodes, lastDay, &blockCharts)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取报块数据失败，\r\n失败信息 %s", err.Error()))
		return
	}
	ChartAddZero(blockCharts)

	e.OK(blockCharts, "查询成功")
}

func (e Finance) FindNodesList(c *gin.Context) []models.FilNodes {
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
		return nil
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
		return nil
	}
	return list
}

func ChartAddZero(list []models.BlockStats) {
	// 为了解决前端图表数据不全问题，添加前一天的数据
	// 1. 获取前一天的数据
	if list == nil {
		// 全部补0
		now := time.Now()
		lastDay := utils.SetTime(now.AddDate(0, 0, -1), now.Hour())
		for i := 1; i <= 24; i++ {
			list = append(list, models.BlockStats{
				HeightTime:            lastDay.Add(time.Hour * time.Duration(i)),
				HeightTimeStr:         lastDay.Add(time.Hour * time.Duration(i)).Format("04:05"),
				BlocksGrowth:          0,
				BlocksRewardGrowthFil: "0",
			})
		}
	} else if len(list) < 24 {
		count := len(list)
		last := list[len(list)-1]
		lastTime := last.HeightTime
		for i := 1; i <= (24 - count); i++ {
			list = append(list, models.BlockStats{
				HeightTime:            lastTime.Add(time.Hour * time.Duration(i)),
				HeightTimeStr:         lastTime.Add(time.Hour * time.Duration(i)).Format("04:05"),
				BlocksGrowth:          0,
				BlocksRewardGrowthFil: "0",
			})
		}
	}
}
