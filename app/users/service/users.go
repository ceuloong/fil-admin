package service

import (
	"errors"
	"fmt"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"gorm.io/gorm"

	"fil-admin/app/users/models"
	"fil-admin/app/users/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type Users struct {
	service.Service
}

// GetPage 获取Users列表
func (e *Users) GetPage(c *dto.UsersGetPageReq, p *actions.DataPermission, list *[]models.Users, count *int64) error {
	var err error
	var data models.Users

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("UsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Users对象
func (e *Users) Get(d *dto.UsersGetReq, p *actions.DataPermission, model *models.Users) error {
	var data models.Users

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Users对象
func (e *Users) Insert(c *dto.UsersInsertReq) error {
	var err error
	var data models.Users
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("UsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Users对象
func (e *Users) Update(c *dto.UsersUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Users{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("UsersService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Users
func (e *Users) Remove(d *dto.UsersDeleteReq, p *actions.DataPermission) error {
	var data models.Users

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveUsers error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Allocate 分配节点给用户
func (e *Users) Allocate(c *dto.UsersAllocateReq, p *actions.DataPermission) error {
	var err error
	var data models.Users

	// 开启事务
	tx := e.Orm.Begin()

	// 查找用户
	err = tx.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 将节点ID数组转换为逗号分隔的字符串
	nodeIdsStr := ""
	for i, id := range c.NodeIds {
		if i > 0 {
			nodeIdsStr += ","
		}
		nodeIdsStr += fmt.Sprintf("%d", id)
	}

	// 更新节点IDs
	if err = tx.Model(&data).Update("node_ids", nodeIdsStr).Error; err != nil {
		tx.Rollback()
		e.Log.Errorf("Update user node_ids error:%s \r\n", err)
		return err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
