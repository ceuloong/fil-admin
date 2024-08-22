package service

import (
	"errors"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type Block struct {
	service.Service
}

// GetPage 获取Block列表
func (e *Block) GetPage(c *dto.BlockGetPageReq, p *actions.DataPermission, list *[]models.Block, count *int64) error {
	var err error
	var data models.Block

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("height desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("BlockService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Block对象
func (e *Block) Get(d *dto.BlockGetReq, p *actions.DataPermission, model *models.Block) error {
	var data models.Block

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetBlock error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Block对象
func (e *Block) Insert(c *dto.BlockInsertReq) error {
	var err error
	var data models.Block
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("BlockService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Block对象
func (e *Block) Update(c *dto.BlockUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Block{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("BlockService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Block
func (e *Block) Remove(d *dto.BlockDeleteReq, p *actions.DataPermission) error {
	var data models.Block

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveBlock error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
