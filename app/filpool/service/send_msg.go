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

type SendMsg struct {
	service.Service
}

// GetPage 获取SendMsg列表
func (e *SendMsg) GetPage(c *dto.SendMsgGetPageReq, p *actions.DataPermission, list *[]models.SendMsg, count *int64) error {
	var err error
	var data models.SendMsg

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Order("Id DESC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SendMsgService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SendMsg对象
func (e *SendMsg) Get(d *dto.SendMsgGetReq, p *actions.DataPermission, model *models.SendMsg) error {
	var data models.SendMsg

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSendMsg error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SendMsg对象
func (e *SendMsg) Insert(c *dto.SendMsgInsertReq) error {
	var err error
	var data models.SendMsg
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SendMsgService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SendMsg对象
func (e *SendMsg) Update(c *dto.SendMsgUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SendMsg{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SendMsgService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SendMsg
func (e *SendMsg) Remove(d *dto.SendMsgDeleteReq, p *actions.DataPermission) error {
	var data models.SendMsg

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSendMsg error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
