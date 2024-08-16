package service

import (
	"bytes"
	"errors"
	"fil-admin/app/finance/models"
	"fil-admin/app/finance/service/dto"
	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx"
	"gorm.io/gorm"
)

type FinanceCash struct {
	service.Service
}

// GetPage 获取FinanceCash列表
func (e *FinanceCash) GetPage(c *dto.FinanceCashGetPageReq, p *actions.DataPermission, list *[]models.FinanceCash, count *int64) error {
	var err error
	var data models.FinanceCash

	err = e.Orm.Model(&data).Preload("Type").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).Order("id desc").
		Find(list).
		//Order(clause.OrderByColumn{Column: clause.Column{Table: clause.CurrentTable, Name: clause.PrimaryKey}, Desc: true}).
		Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FinanceCashService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FinanceCash对象
func (e *FinanceCash) Get(d *dto.FinanceCashGetReq, p *actions.DataPermission, model *models.FinanceCash) error {
	var data models.FinanceCash

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFinanceCash error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
func DecimalValue(str float64) decimal.Decimal {
	value := decimal.NewFromFloat(str)
	return value
}

// Insert 创建FinanceCash对象
func (e *FinanceCash) Insert(c *dto.FinanceCashInsertReq) error {
	var err error
	var data models.FinanceCash
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()

	//判断收支类型如果是支出，amount为负数
	var fType models.FinanceType
	tx.First(&fType, data.TypeId)
	amount, _ := data.Amount.Float64()
	e.Log.Errorf("amount:%s", data.Name)
	amount = math.Abs(amount)
	data.Amount = decimal.NewFromFloat(amount)
	if fType.InOut == 2 {
		data.Amount = decimal.Zero.Sub(data.Amount)
	}
	e.Log.Errorf("amount:%s", amount)
	//查询最后一条保存的余额，更新当前余额
	var lastCash models.FinanceCash
	tx.Model(&models.FinanceCash{}).Where("dict_id = ?", c.DictId).Order("id desc").Limit(1).Find(&lastCash)
	if lastCash.Id == 0 {
		data.Balance = data.Amount
	} else {
		data.Balance = lastCash.Balance.Add(data.Amount)
	}

	if data.Balance.LessThan(decimal.Zero) {
		e.Log.Errorf("可用余额不足以支付:%s", err)
		return err
	}

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
	return nil

}

// Update 修改FinanceCash对象
func (e *FinanceCash) Update(c *dto.FinanceCashUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FinanceCash{}
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.First(&data, c.GetId())
	c.Generate(&data)

	var nextCash models.FinanceCash
	tx.Model(&models.FinanceCash{}).Where("id > ? AND dict_id = ?", c.GetId(), c.DictId).Limit(1).Find(&nextCash)
	if nextCash.Id > 0 {
		err = errors.New("只允许修改最后一条数据。")
		e.Log.Errorf("UpdateFinanceCash error:%s\n。", err)
		return err
	}

	//判断收支类型如果是支出，amount为负数
	var fType models.FinanceType
	tx.First(&fType, data.TypeId)
	amount, _ := data.Amount.Float64()
	e.Log.Errorf("amount:%s", data.Amount.String())
	amount = math.Abs(amount)
	data.Amount = decimal.NewFromFloat(amount)
	if fType.InOut == 2 {
		data.Amount = decimal.Zero.Sub(data.Amount)
	}
	//查询最后一条保存的余额，更新当前余额
	var lastCash models.FinanceCash
	tx.Model(&models.FinanceCash{}).Where("id < ? AND dict_id = ?", c.GetId(), c.DictId).Order("id desc").Limit(1).Find(&lastCash)
	if lastCash.Id == 0 {
		data.Balance = data.Amount
	} else {
		data.Balance = lastCash.Balance.Add(data.Amount)
	}

	if data.Balance.LessThan(decimal.Zero) {
		err = errors.New("可用余额不足以支付。")
		e.Log.Errorf("UpdateFinanceCash error:%s\n", err)
		return err
	}

	db := tx.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateFinanceCash error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FinanceCash
func (e *FinanceCash) Remove(d *dto.FinanceCashDeleteReq, p *actions.DataPermission) error {
	var data models.FinanceCash

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFinanceCash error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// ExportXlsx 导出FinanceCash数据
func (e *FinanceCash) ExportXlsx(r *dto.FinanceCashGetPageReq, c *gin.Context) error {
	var err error
	var data models.FinanceCash
	var list []models.FinanceCash
	err = e.Orm.Model(&data).Preload("Type").
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
		).
		Find(&list).
		Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("FinanceCashService GetPage error:%s \r\n", err)
		return err
	}

	e.saveXlsx(list)

	return nil
}

func (e *FinanceCash) saveXlsx(list []models.FinanceCash) {
	file, err := xlsx.OpenFile("template/Cash Book2023-Coffee Cloud.xlsx")
	if err != nil {
		panic(err)
	}
	first := file.Sheets[0]
	for i, cash := range list {
		row := first.Row(5 + i)
		cell := row.Cells[0]
		cell.Value = cash.CreatedAt.Month().String()
		cell = row.Cells[1]
		cell.Value = strconv.Itoa(cash.CreatedAt.Day())
		cell = row.Cells[2]
		cell.Value = fmt.Sprintf("PT000%d", i+1)
		cell = row.Cells[3]
		cell.Value = cash.Memo
		index := 4
		if cash.Amount.LessThan(decimal.Zero) {
			index = 5
			cash.Amount = decimal.Zero.Sub(cash.Amount)
		}
		cell = row.Cells[index]
		cell.Value = cash.Amount.String()
		cell = row.Cells[6]
		cell.Value = cash.Balance.String()
	}
	fileName := "Cash Book2023-Coffee Cloud.xlsx"
	/*fileName = fmt.Sprintf("%s.xlsx", fileName)

	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	println(content)
	strByte, _ := json.Marshal(content)
	str := string(strByte)
	println(strconv.Itoa(len(str)))
	w.Header().Set("Content-Length", strconv.Itoa(len(str)))
	//http.ServeContent(w, r, fileName, time.Now(), content)
	w.WriteHeader(http.StatusOK)*/

	err = file.Save("/Users/hello/Documents/Cash Book/" + fileName)
	if err != nil {
		panic(err)
	}
}

// DataToExcel 数据导出excel, dataList里面的对象为指针
func DataToExcel(w http.ResponseWriter, r *http.Request, titleList []string, dataList []interface{}, fileName string) {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
		cell.GetStyle().Font.Color = "00FF0000"
	}
	// 插入内容
	for _, v := range dataList {
		row := sheet.AddRow()
		row.WriteStruct(v, -1)
	}
	fileName = fmt.Sprintf("%s.xlsx", fileName)
	//_ = file.Save(fileName)
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	http.ServeContent(w, r, fileName, time.Now(), content)
}
