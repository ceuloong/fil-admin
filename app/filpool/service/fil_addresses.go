package service

import (
	"errors"
	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type FilAddresses struct {
	service.Service
}

// GetPage 获取FilAddresses列表
func (e *FilAddresses) GetPage(c *dto.FilAddressesGetPageReq, p *actions.DataPermission, list *[]models.FilAddresses, count *int64) error {
	var err error
	var data models.FilAddresses

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FilAddressesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FilAddresses对象
func (e *FilAddresses) Get(d *dto.FilAddressesGetReq, p *actions.DataPermission, model *models.FilAddresses) error {
	var data models.FilAddresses

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFilAddresses error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FilAddresses对象
func (e *FilAddresses) Insert(c *dto.FilAddressesInsertReq) error {
	var err error
	var data models.FilAddresses
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FilAddressesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FilAddresses对象
func (e *FilAddresses) Update(c *dto.FilAddressesUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilAddresses{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FilAddressesService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FilAddresses
func (e *FilAddresses) Remove(d *dto.FilAddressesDeleteReq, p *actions.DataPermission) error {
	var data models.FilAddresses

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFilAddresses error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
