package jobs

import (
	"encoding/json"
	"fil-admin/app/filpool/models"
	"fil-admin/common/redis"
	"fil-admin/utils"
	"fmt"
	"log"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"github.com/shopspring/decimal"
)

const (
	// BlockStatsKey : block stats key
	BlockStatsKey string = "block_stats"
	QueueLength   int64  = 720 // 720个小时，最近一个月数据
)

type BlockStatsExec struct {
	service.Service
}

func (t BlockStatsExec) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore BlockStatsExec exec success"
	fmt.Println(str, arg.(string))

	db := GetDb()
	if db == nil {
		fmt.Println("db is nil")
		return nil
	}
	t.Orm = db

	// 获取所有节点
	var nodes []models.FilNodes
	err := t.FindAllNode(&nodes)
	if err != nil {
		log.Printf("failed to get all nodes: %s \r\n", err)
		return err
	}
	for _, n := range nodes {
		bs, _ := t.Get24HBlockWinningStatsFromRedis(BlockStatsKey + "_" + n.Node)
		node := t.Generate(n, bs)
		// update node block stats
		if err = t.UpdateNodes(node); err != nil {
			log.Printf("failed to update node %s block stats: %s \r\n", n.Node, err)
		}
	}

	return nil
}

type BlockStats struct {
	BlocksGrowth          int             `json:"blocksGrowth"`
	BlocksRewardGrowthFil decimal.Decimal `json:"blocksRewardGrowthFil"`
	HeightTimeStr         string          `json:"heightTimeStr"`
	HeightTime            time.Time       `json:"heightTime"`
}

// Get24HBlockStatsFromRedis 从Redis获取最近24小时的统计数据
func (b BlockStatsExec) Get24HBlockStatsFromRedis(key string) ([]BlockStats, error) {
	list, err := redis.LRangeRedis(key, -24, -1)
	if err != nil {
		return nil, fmt.Errorf("failed to get block stats from Redis: %w", err)
	}
	var stats []BlockStats
	for _, v := range list {
		var stat BlockStats
		if err := json.Unmarshal([]byte(v), &stat); err != nil {
			return nil, fmt.Errorf("failed to unmarshal block stats: %w", err)
		}
		stats = append(stats, stat)
	}
	return stats, nil
}

func (b BlockStatsExec) Get24HBlockWinningStatsFromRedis(key string) (BlockStats, error) {
	stats := b.OneDayAddZero(key)
	var totalBlock int
	var totalReward decimal.Decimal
	for _, v := range stats {
		totalBlock += v.BlocksGrowth
		totalReward = totalReward.Add(v.BlocksRewardGrowthFil)
	}
	return BlockStats{
		BlocksGrowth:          totalBlock,
		BlocksRewardGrowthFil: totalReward,
		HeightTimeStr:         "24小时",
		HeightTime:            time.Now(),
	}, nil
}

// 初始化24小时数组
func (b BlockStatsExec) OneDayAddZero(key string) []BlockStats {

	var stats []BlockStats
	now := time.Now()
	lastDay := utils.SetTime(now.AddDate(0, 0, -1), now.Hour())
	// 先初始化24个点
	for i := 0; i < 24; i++ {
		newHour := lastDay.Add(time.Hour * time.Duration(i)).Hour()
		stats = append(stats, BlockStats{
			BlocksGrowth:          0,
			BlocksRewardGrowthFil: decimal.Zero,
			HeightTimeStr:         b.GetHourStr(newHour),
			HeightTime:            lastDay.Add(time.Hour * time.Duration(i)),
		})
	}

	var lastStats []BlockStats
	redisStats, err := b.Get24HBlockStatsFromRedis(key)
	if err != nil {
		return stats
	}
	for _, v := range stats {
		isEmpty := true
		for _, rv := range redisStats {
			if v.HeightTime.Equal(rv.HeightTime) {
				lastStats = append(lastStats, rv)
				isEmpty = false
				break
			}
		}
		if isEmpty {
			lastStats = append(lastStats, v)
		}
	}

	return lastStats
}

func (b BlockStatsExec) GetHourStr(hour int) string {
	if hour < 10 {
		return fmt.Sprintf("0%d:00", hour)
	}
	return fmt.Sprintf("%d:00", hour)

}

// FindAllNode 获取FilNodes列表 所有符合条件的记录，不分页
func (e *BlockStatsExec) FindAllNode(list *[]models.FilNodes) error {
	err := e.Orm.Model(&models.FilNodes{}).Where("status > 0").Find(list).Error

	if err != nil {
		log.Printf("FindAllNode error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *BlockStatsExec) UpdateNodes(node UpdateNodes) error {
	err := e.Orm.Model(&models.FilNodes{}).Where("node = ?", node.Node).Updates(node).Error

	if err != nil {
		log.Printf("update nodes %s error, %s", node.Node, err.Error())
		return err
	}
	return nil
}

func (e *BlockStatsExec) Generate(node models.FilNodes, bs BlockStats) UpdateNodes {
	updateNodes := UpdateNodes{
		Node:            node.Node,
		BlocksMined24h:  bs.BlocksGrowth,
		TotalRewards24h: bs.BlocksRewardGrowthFil,
	}
	if node.QualityAdjPower.GreaterThan(decimal.Zero) {
		// 计算24小时平均收益  FIL/PiB
		updateNodes.MiningEfficiency = updateNodes.TotalRewards24h.Div(node.QualityAdjPower).RoundDown(1)
		// 计算24小时Luck值
		if node.AverageWinRate.GreaterThan(decimal.Zero) {
			updateNodes.LuckyValue24h = decimal.NewFromInt(int64(bs.BlocksGrowth)).Div(node.AverageWinRate).Mul(decimal.NewFromInt(7)).RoundDown(4)
		}
	}

	return updateNodes
}

type UpdateNodes struct {
	Node             string          `json:"node"`
	BlocksMined24h   int             `json:"blocksMined24h"`
	TotalRewards24h  decimal.Decimal `json:"totalRewards24h"`
	MiningEfficiency decimal.Decimal `json:"miningEfficiency"`
	LuckyValue24h    decimal.Decimal `json:"luckyValue24h"`
}
