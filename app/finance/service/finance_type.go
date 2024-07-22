package service

import (
	"errors"
	"fil-admin/app/finance/models"
	"fil-admin/app/finance/service/dto"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	cDto "fil-admin/common/dto"
)

type FinanceType struct {
	service.Service
}

// Get 获取FinanceType对象
func (e *FinanceType) Get(d *dto.FinanceTypeGetReq, model *models.FinanceType) error {
	var data models.FinanceType

	err := e.Orm.Model(&data).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFinanceType error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysDept对象
func (e *FinanceType) Insert(c *dto.FinanceTypeInsertReq) error {
	var err error
	var data models.FinanceType
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	typePath := pkg.IntToString(data.TypeId) + "/"
	if data.ParentId != 0 {
		var typeP models.FinanceType
		tx.First(&typeP, data.ParentId)
		typePath = typeP.TypePath + typePath
	} else {
		typePath = "/0/" + typePath
	}
	var mp = map[string]string{}
	mp["type_path"] = typePath
	if err := tx.Model(&data).Update("type_path", typePath).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysDept对象
func (e *FinanceType) Update(c *dto.FinanceTypeUpdateReq) error {
	var err error
	var model = models.FinanceType{}
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.First(&model, c.GetId())
	c.Generate(&model)

	typePath := pkg.IntToString(model.TypeId) + "/"
	if model.ParentId != 0 {
		var typeP models.FinanceType
		tx.First(&typeP, model.ParentId)
		typePath = typeP.TypePath + typePath
	} else {
		typePath = "/0/" + typePath
	}
	model.TypePath = typePath
	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateFinanceType error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FinanceType
func (e *FinanceType) Remove(d *dto.FinanceTypeDeleteReq) error {
	var err error
	var data models.FinanceType

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}

// GetFinanceTypeList 获取组织数据
func (e *FinanceType) getList(c *dto.FinanceTypeGetPageReq, list *[]models.FinanceType) error {
	var err error
	var data models.FinanceType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// SetTypeTree 设置组织数据
func (e *FinanceType) SetTypeTree(c *dto.FinanceTypeGetPageReq) (m []dto.TypeLabel, err error) {
	var list []models.FinanceType
	err = e.getList(c, &list)

	m = make([]dto.TypeLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.TypeLabel{}
		e.Id = list[i].TypeId
		e.Label = list[i].Name
		typesInfo := typeTreeCall(&list, e)

		m = append(m, typesInfo)
	}
	return
}

// Call 递归构造组织数据
func typeTreeCall(typeList *[]models.FinanceType, label dto.TypeLabel) dto.TypeLabel {
	list := *typeList
	min := make([]dto.TypeLabel, 0)
	for j := 0; j < len(list); j++ {
		if label.Id != list[j].ParentId {
			continue
		}
		mi := dto.TypeLabel{Id: list[j].TypeId, Label: list[j].Name, Children: []dto.TypeLabel{}}
		ms := typeTreeCall(typeList, mi)
		min = append(min, ms)
	}
	label.Children = min
	return label
}

// SetTypePage 设置type页面数据
func (e *FinanceType) SetTypePage(c *dto.FinanceTypeGetPageReq) (m []models.FinanceType, err error) {
	var list []models.FinanceType
	err = e.getList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := e.typePageCall(&list, list[i])
		m = append(m, info)
	}
	return
}

func (e *FinanceType) typePageCall(typesList *[]models.FinanceType, menu models.FinanceType) models.FinanceType {
	list := *typesList
	min := make([]models.FinanceType, 0)
	for j := 0; j < len(list); j++ {
		if menu.TypeId != list[j].ParentId {
			continue
		}
		mi := models.FinanceType{}
		mi.TypeId = list[j].TypeId
		mi.ParentId = list[j].ParentId
		mi.TypePath = list[j].TypePath
		mi.Name = list[j].Name
		mi.Sort = list[j].Sort
		mi.Memo = list[j].Memo
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.FinanceType{}
		ms := e.typePageCall(typesList, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
