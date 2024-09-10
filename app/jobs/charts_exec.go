package jobs

import (
	sysModels "fil-admin/app/admin/models"
	"fil-admin/app/filpool/models"
	"fil-admin/utils"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ChartsExec struct {
	Orm *gorm.DB
}

func (t ChartsExec) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ChartsExec exec success"

	db := GetDb()
	if db == nil {
		fmt.Println("db is nil")
		return nil
	}
	t.Orm = db
	t.HandlerCharts()

	fmt.Println(str, arg.(string))

	return nil
}

func GetDb() *gorm.DB {

	dbs := sdk.Runtime.GetDb()
	for k, db := range dbs {
		if db.Name() == "mysql" {
			return db
		}
		fmt.Println(k, db.Name())
	}
	return dbs["*"]
}

/**
 * 保存矿池图表数据
 * 根据节点的所属部门分类保存各部门总算力
 */
func (t ChartsExec) HandlerCharts() {

	now := time.Now()
	lastTime := utils.SetTime(now, now.Hour())
	poolChart := &models.FilPoolChart{
		LastTime:  lastTime,
		PowerUnit: "PiB",
		DeptId:    0,
	}

	nodes := []models.FilNodes{}

	err := t.FindAllNode(&nodes)
	if err != nil {
		return // 错误处理
	}

	deptPoolChart := make(map[int]*models.FilPoolChart)
	hasPowerCount := 0
	noPowerCount := 0
	deptMap := t.getDeptMap() // 获取所有部门，保存各部门总算力

	for _, n := range nodes {
		if err := t.SaveNodesChart(n); err != nil {
			log.Printf("保存节点%s的图表数据失败：%s\n", n.Node, err)
		}

		updatePoolChart(poolChart, n) // 更新矿池图表数据

		if _, ok := deptPoolChart[n.DeptId]; !ok && n.DeptId > 0 {
			deptPoolChart[n.DeptId] = &models.FilPoolChart{
				LastTime:  lastTime,
				PowerUnit: "PiB",
				DeptId:    n.DeptId,
			}
		}
		if _, ok := deptPoolChart[n.DeptId]; ok {
			updatePoolChart(deptPoolChart[n.DeptId], n) // 更新部门矿池图表数据
		}
		if pId, ok := deptMap[n.DeptId]; ok && pId > 0 {
			if _, ok := deptPoolChart[pId]; !ok {
				deptPoolChart[pId] = &models.FilPoolChart{
					LastTime:  lastTime,
					PowerUnit: "PiB",
					DeptId:    pId,
				}
			}
			updatePoolChart(deptPoolChart[pId], n) // 更新父部门矿池图表数据
		}

		if n.QualityAdjPower.IsZero() {
			noPowerCount++
		} else {
			hasPowerCount++
		}
	}

	err = t.SavePoolChart(poolChart) // 保存矿池图表数据
	if err != nil {
		log.Printf("保存矿池图表数据失败：%s\n", err)
	}

	log.Printf("一共更新的 %d 个节点，其中有算力的节点 %d 个, 算力为0的节点 %d 个。\n", len(nodes), hasPowerCount, noPowerCount)

	for k, v := range deptPoolChart {
		log.Printf("保存部门%d的矿池数据\n", k)
		t.SavePoolChart(v)
	}
}

func updatePoolChart(poolChart *models.FilPoolChart, node models.FilNodes) {
	// 累加节点数据到矿池图表
	poolChart.Balance = poolChart.Balance.Add(node.Balance)
	poolChart.AvailableBalance = poolChart.AvailableBalance.Add(node.AvailableBalance)
	poolChart.SectorPledgeBalance = poolChart.SectorPledgeBalance.Add(node.SectorPledgeBalance)
	poolChart.VestingFunds = poolChart.VestingFunds.Add(node.VestingFunds)
	poolChart.QualityAdjPower = poolChart.QualityAdjPower.Add(node.QualityAdjPower)
	poolChart.PowerPoint = poolChart.PowerPoint.Add(node.PowerPoint)
	poolChart.ControlBalance = poolChart.ControlBalance.Add(node.ControlBalance)
	poolChart.RewardValue = poolChart.RewardValue.Add(node.RewardValue)
	poolChart.NodesCount++
}

// FindAllNode 获取FilNodes列表 所有符合条件的记录，不分页
func (e *ChartsExec) FindAllNode(list *[]models.FilNodes) error {
	err := e.Orm.Model(&models.FilNodes{}).Where("status > 0").Find(list).Error

	if err != nil {
		log.Printf("FindAllNode error:%s \r\n", err)
		return err
	}
	return nil
}

// SaveNodesChart 保存节点图表数据
func (e *ChartsExec) SaveNodesChart(nodes models.FilNodes) error {
	nodesChart, updateNodes := e.GetNodesChart(nodes)

	err := e.InsertNodesChart(nodesChart).Error
	if err != nil {
		log.Printf("save nodes chart error, %s", err.Error())
		return err
	}
	err = e.Orm.Model(&models.FilNodes{}).Where("node = ?", nodes.Node).Updates(updateNodes).Error
	if err != nil {
		log.Printf("update nodes error, %s", err.Error())
		return err
	}
	return nil
}

func (e *ChartsExec) InsertNodesChart(nodesChart models.NodesChart) *gorm.DB {
	return e.Orm.Create(&nodesChart)
}

func (e *ChartsExec) SavePoolChart(poolChart *models.FilPoolChart) error {
	return e.Orm.Create(&poolChart).Error
}

func (e *ChartsExec) GetLastOneByTime(node models.FilNodes, time time.Time) models.NodesChart {
	var lastOne models.NodesChart
	e.Orm.Model(&models.NodesChart{}).Where("last_time = ? AND node = ?", time, node.Node).First(&lastOne)
	return lastOne
}

// 处理重新封装图表昨日和上月数据
func (e *ChartsExec) GetNodesChart(nodes models.FilNodes) (models.NodesChart, UpdateNodesDelta) {
	currentTime := time.Now()
	lastDay := currentTime.AddDate(0, 0, -1).Truncate(time.Hour)
	lastOne := e.GetLastOneByTime(nodes, lastDay)

	//lastMonthLastDay := currentTime.AddDate(0, 0, -currentTime.Day()).Truncate(time.Hour)
	lastMonthLastDay := utils.SetTime(currentTime.AddDate(0, 0, -currentTime.Day()), 0)
	lastMonthLastOne := e.GetLastOneByTime(nodes, lastMonthLastDay)

	updateNodes := UpdateNodesDelta{}
	// 24小时内算力增量 保存的是单位为GiB的算力
	if lastOne.ID > 0 {
		updateNodes.QualityAdjPowerDelta24h = nodes.QualityAdjPower.Sub(lastOne.QualityAdjPower).Mul(decimal.NewFromFloat(math.Pow10(6)))
	} else {
		updateNodes.QualityAdjPowerDelta24h = decimal.Zero
	}
	// 30天内算力增量 保存的是单位为GiB的算力
	if lastMonthLastOne.ID > 0 {
		updateNodes.QualityAdjPowerDelta30d = nodes.QualityAdjPower.Sub(lastMonthLastOne.QualityAdjPower).Mul(decimal.NewFromFloat(math.Pow10(6)))
	}
	nodesChart := models.NodesChart{
		Node:                         nodes.Node,
		AvailableBalance:             nodes.AvailableBalance,
		Balance:                      nodes.Balance,
		SectorPledgeBalance:          nodes.SectorPledgeBalance,
		VestingFunds:                 nodes.VestingFunds,
		Height:                       nodes.Height,
		LastTime:                     currentTime.Truncate(time.Hour),
		RewardValue:                  nodes.RewardValue,
		WeightedBlocks:               nodes.WeightedBlocks,
		QualityAdjPower:              nodes.QualityAdjPower,
		PowerUnit:                    nodes.PowerUnit,
		PowerPoint:                   nodes.PowerPoint,
		ControlBalance:               nodes.ControlBalance,
		BlocksMined24h:               nodes.BlocksMined24h,
		TotalRewards24h:              nodes.TotalRewards24h,
		LuckyValue24h:                nodes.LuckyValue24h,
		QualityAdjPowerDelta24h:      nodes.QualityAdjPowerDelta24h,
		ReceiveAmount:                nodes.ReceiveAmount,
		BurnAmount:                   nodes.BurnAmount,
		SendAmount:                   nodes.SendAmount,
		LastAvailableBalance:         lastOne.AvailableBalance,
		LastBalance:                  lastOne.Balance,
		LastSectorPledgeBalance:      lastOne.SectorPledgeBalance,
		LastVestingFunds:             lastOne.VestingFunds,
		LastRewardValue:              lastOne.RewardValue,
		LastWeightedBlocks:           lastOne.WeightedBlocks,
		LastQualityAdjPower:          lastOne.QualityAdjPower,
		LastReceiveAmount:            lastOne.ReceiveAmount,
		LastBurnAmount:               lastOne.BurnAmount,
		LastSendAmount:               lastOne.SendAmount,
		LastMonthSectorPledgeBalance: lastMonthLastOne.SectorPledgeBalance,
		LastMonthRewardValue:         lastMonthLastOne.RewardValue,
		LastMonthWeightedBlocks:      lastMonthLastOne.WeightedBlocks,
		LastMonthQualityAdjPower:     lastMonthLastOne.QualityAdjPower,
		LastMonthReceiveAmount:       lastMonthLastOne.ReceiveAmount,
		LastMonthBurnAmount:          lastMonthLastOne.BurnAmount,
		LastMonthSendAmount:          lastMonthLastOne.SendAmount,
		TimeTag:                      nodes.TimeTag,
	}

	return nodesChart, updateNodes
}

// getDeptMap 获取部门map key:子部门id value:父部门id
func (e *ChartsExec) getDeptMap() map[int]int {
	deptMap := make(map[int]int)
	list := []sysModels.SysDept{}
	err := e.Orm.Model(&sysModels.SysDept{}).Where("parent_id > 0").Find(&list).Error
	if err != nil {
		log.Printf("getDeptMap error:%s \r\n", err)
		return nil
	}
	for _, v := range list {
		deptMap[v.DeptId] = v.ParentId
	}

	return deptMap
}

type UpdateNodesDelta struct {
	QualityAdjPowerDelta24h decimal.Decimal `json:"qualityAdjPowerDelta24h"`
	QualityAdjPowerDelta30d decimal.Decimal `json:"qualityAdjPowerDelta30d"`
}
