package service

import (
	"errors"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
	"fil-admin/utils"
)

type FilPoolChart struct {
	service.Service
}

// GetPage 获取FilNodes列表
func (e *FilPoolChart) GetPage(c *dto.FilPoolChartGetPageReq, p *actions.DataPermission, list *[]models.FilPoolChart) error {
	var err error
	var data models.FilPoolChart

	// 一个月前的时间 转换为string
	lastTime := utils.SetTime(time.Now().AddDate(0, 0, -30), 0).Format(time.DateTime)
	tx := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		)
	if c.DeptId == 0 {
		tx.Where("dept_id = 0")
	}
	tx.Where("last_time >= ?", lastTime)

	err = tx.Find(list).Error
	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetPage 获取FilNodes列表
func (e *FilPoolChart) GetDayAvgPage(c *dto.FilPoolChartGetPageReq, p *actions.DataPermission, list *[]models.FilPoolChart) error {
	var err error
	var data models.FilPoolChart

	// 一个月前的时间 转换为string
	lastTime := utils.SetTime(time.Now().AddDate(0, 0, -30), 0).Format(time.DateTime)
	tx := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Select("avg(quality_adj_power) as quality_adj_power, AVG(available_balance) as available_balance, AVG(balance) as balance, AVG(sector_pledge_balance) as sector_pledge_balance, AVG(vesting_funds) as vesting_funds, AVG(control_balance) as control_balance, TO_DAYS(last_time) as lastDays, MIN(last_time) as last_time")
	if c.DeptId == 0 {
		tx.Where("dept_id = 0")
	}
	tx.Where("last_time >= ?", lastTime).Group("TO_DAYS(last_time)")

	err = tx.Find(list).Error
	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FilNodes对象
func (e *FilPoolChart) Get(d *dto.FilPoolChartGetReq, p *actions.DataPermission, model *models.FilPoolChart) error {
	var data models.FilPoolChart

	tx := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(d.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		)
	if d.DeptId == 0 {
		tx.Where("dept_id = 0")
	}
	err := tx.Order("id DESC").
		First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service FilPoolChart Get error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

/**
 * @Description: 获取最近24小时的数据，24个点
 * @receiver e FilPoolChart
 * @param c
 * @param p
 * @param list
 * @return error
 */
func (e *FilPoolChart) GetAppChart(deptId int, lastDay time.Time, list *[]models.FilPoolChart) error {

	tx := e.Orm.Model(&models.FilPoolChart{}).
		Where("last_time >= ?", lastDay)
	if deptId > 0 {
		tx.Where("dept_id = ?", deptId)
	} else {
		tx.Where("dept_id = 0")
	}
	err := tx.Find(list).Error

	if err != nil {
		e.Log.Errorf("GetAppChart error:%s \r\n", err)
		return err
	}
	return nil
}
