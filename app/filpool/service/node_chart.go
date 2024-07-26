package service

import (
	"errors"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type NodesChart struct {
	service.Service
}

// GetPage 获取NodeChart列表
func (e *NodesChart) GetPage(c *dto.NodeChartGetPageReq, p *actions.DataPermission, list *[]models.NodesChart, count *int64) error {
	var err error
	var data models.NodesChart

	lastTime := c.LastTime
	c.LastTime = ""

	tx := e.Orm.Model(&data)

	tx.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	)
	if lastTime != "" {
		tx.Where("TO_DAYS(last_time)=TO_DAYS(?) AND HOUR(last_time)=HOUR(?)", lastTime, lastTime)
	}
	tx.Order("id desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count)
	err = tx.Error
	if err != nil {
		e.Log.Errorf("NodeChartService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetList 获取FilNodes对象
func (e *NodesChart) GetList(d *dto.NodeChartGetReq, p *actions.DataPermission, list *[]models.NodesChart) error {
	var data models.NodesChart

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(d.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).Order("id desc").
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFilNodes error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *NodesChart) GetChartList(tx *gorm.DB, lastTime time.Time, nodes []string, list *[]models.NodesChart) (err error) {
	err = tx.Table("nodes_chart").Where("last_time < ? and node in (?)", lastTime, nodes).Order("id DESC").Find(&list).Error
	if err != nil {
		e.Log.Errorf("get nodes_chart error, %s", err.Error())
		return
	}
	return
}

func (e *NodesChart) GetLastOneByTime(node models.FilNodes, time time.Time) models.NodesChart {
	var lastOne models.NodesChart
	e.Orm.Model(&models.NodesChart{}).Where("TO_DAYS(last_time) = TO_DAYS(?) AND node = ?", time, node.Node).Order("last_time DESC").First(&lastOne)
	return lastOne
}
