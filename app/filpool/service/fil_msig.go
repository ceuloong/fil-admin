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

type FilMsig struct {
	service.Service
}

// GetPage 获取FilMsig列表
func (e *FilMsig) GetPage(c *dto.FilMsigGetPageReq, p *actions.DataPermission, list *[]models.FilMsig, count *int64) error {
	var err error
	var data models.FilMsig

	err = e.Orm.Model(&data).Preload("Dept").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FilMsigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FilMsig对象
func (e *FilMsig) Get(d *dto.FilMsigGetReq, p *actions.DataPermission, model *models.FilMsig) error {
	var data models.FilMsig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFilMsig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FilMsig对象
func (e *FilMsig) Insert(c *dto.FilMsigInsertReq) error {
	var err error
	var data models.FilMsig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FilMsigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FilMsig对象
func (e *FilMsig) Update(c *dto.FilMsigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilMsig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("FilMsigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FilMsig
func (e *FilMsig) Remove(d *dto.FilMsigDeleteReq, p *actions.DataPermission) error {
	var data models.FilMsig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFilMsig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
