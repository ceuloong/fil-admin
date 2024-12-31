package service

import (
	"errors"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/app/users/models"
	"fil-admin/app/users/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type UserVerifications struct {
	service.Service
}

// GetPage 获取UserVerifications列表
func (e *UserVerifications) GetPage(c *dto.UserVerificationsGetPageReq, p *actions.DataPermission, list *[]models.UserVerifications, count *int64) error {
	var err error
	var data models.UserVerifications

	err = e.Orm.Model(&data).Preload("Users").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("UserVerificationsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取UserVerifications对象
func (e *UserVerifications) Get(d *dto.UserVerificationsGetReq, p *actions.DataPermission, model *models.UserVerifications) error {
	var data models.UserVerifications

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetUserVerifications error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建UserVerifications对象
func (e *UserVerifications) Insert(c *dto.UserVerificationsInsertReq) error {
	var err error
	var data models.UserVerifications
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("UserVerificationsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改UserVerifications对象
func (e *UserVerifications) Update(c *dto.UserVerificationsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.UserVerifications{}

	// 开启事务
	tx := e.Orm.Begin()

	err = tx.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 获取旧的状态，用于判断是否需要更新用户表
	oldStatus := data.Status

	// 更新认证信息
	c.Generate(&data)
	if err = tx.Save(&data).Error; err != nil {
		tx.Rollback()
		e.Log.Errorf("UserVerificationsService Save error:%s \r\n", err)
		return err
	}

	// 如果认证状态发生变化，同步更新用户表
	if oldStatus != data.Status {
		if err = tx.Model(&models.Users{}).
			Where("id = ?", data.UserId).
			Update("verify_status", data.Status).Error; err != nil {
			tx.Rollback()
			e.Log.Errorf("Update user verify_status error:%s \r\n", err)
			return err
		}
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Remove 删除UserVerifications
func (e *UserVerifications) Remove(d *dto.UserVerificationsDeleteReq, p *actions.DataPermission) error {
	var data models.UserVerifications

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveUserVerifications error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
