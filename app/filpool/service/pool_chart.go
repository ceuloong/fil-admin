package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type FilPoolChart struct {
	service.Service
}

// GetPage 获取FilNodes列表
func (e *FilPoolChart) GetPage(c *dto.FilPoolChartGetPageReq, p *actions.DataPermission, list *[]models.FilPoolChart) error {
	var err error
	var data models.FilPoolChart

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).
		Order("Id DESC").
		Find(list).Limit(100).Error
	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FilNodes对象
func (e *FilPoolChart) Get(d *dto.FilPoolChartGetReq, p *actions.DataPermission, model *models.FilPoolChart) error {
	var data models.FilPoolChart

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(d.GetNeedSearch()),
			actions.Permission(data.TableName(), p),
		).Order("id DESC").
		First(model).Error
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
