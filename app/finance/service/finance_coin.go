package service

import (
	"errors"
	"fil-admin/app/finance/models"
	"fil-admin/app/finance/service/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type FinanceCoin struct {
	service.Service
}

// GetPage 获取FinanceCoin列表
func (e *FinanceCoin) GetPage(c *dto.FinanceCoinGetPageReq, p *actions.DataPermission, list *[]models.FinanceCoin, count *int64) error {
	var err error
	var data models.FinanceCoin

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("id desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FinanceCoinService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FinanceCoin对象
func (e *FinanceCoin) Get(d *dto.FinanceCoinGetReq, p *actions.DataPermission, model *models.FinanceCoin) error {
	var data models.FinanceCoin

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFinanceCoin error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FinanceCoin对象
func (e *FinanceCoin) Insert(c *dto.FinanceCoinInsertReq) error {
	var err error
	var data models.FinanceCoin
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FinanceCoinService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FinanceCoin对象
func (e *FinanceCoin) Update(c *dto.FinanceCoinUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FinanceCoin{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FinanceCoinService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FinanceCoin
func (e *FinanceCoin) Remove(d *dto.FinanceCoinDeleteReq, p *actions.DataPermission) error {
	var data models.FinanceCoin

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFinanceCoin error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
