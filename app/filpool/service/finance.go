package service

import (
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"

	"fil-admin/app/filpool/models"
)

type Finance struct {
	service.Service
}

// SumBlockStats 获取矿池的当天报块统计
func (e *Finance) SumBlockStats(nodes []string, lastDay time.Time, list *[]models.BlockStats) error {

	err := e.Orm.Model(&models.BlockStats{}).
		Select("height_time, height_time_str, SUM(Blocks_growth) as blocks_growth, SUM(blocks_reward_growth_fil) as blocks_reward_growth_fil").
		Where("height_time >= ? AND node in (?)", lastDay, nodes).
		Group("height_time, height_time_str").
		Find(list).Error

	if err != nil {
		e.Log.Errorf("SumBlockStats error:%s \r\n", err)
		return err
	}
	return nil
}
