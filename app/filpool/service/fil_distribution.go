package service

import (
	"errors"
	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"
	"fmt"
	"log"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/service"
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx"
	"gorm.io/gorm"

	"fil-admin/common/actions"
	cDto "fil-admin/common/dto"
)

type FilDistribution struct {
	service.Service
}

// GetPage 获取FilDistribution列表
func (e *FilDistribution) GetPage(c *dto.FilDistributionGetPageReq, p *actions.DataPermission, list *[]models.FilDistribution, count *int64) error {
	var err error
	var data models.FilDistribution

	err = e.Orm.Model(&data).Preload("FilNode").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("FilDistributionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FilDistribution对象
func (e *FilDistribution) Get(d *dto.FilDistributionGetReq, p *actions.DataPermission, model *models.FilDistribution) error {
	var data models.FilDistribution

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFilDistribution error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FilDistribution对象
func (e *FilDistribution) Insert(c *dto.FilDistributionInsertReq) error {
	var err error
	var data models.FilDistribution

	c.Generate(&data)

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FilDistributionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FilDistribution对象
func (e *FilDistribution) Update(c *dto.FilDistributionUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilDistribution{}
	db := e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("FilDistributionService Update error:%s \r\n", err)
		return err
	}

	err = e.Orm.Table(data.TableName()).Updates(c).Error
	if err != nil {
		e.Log.Errorf("FilDistributionService Update error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// UpdateStatus 修改FilDistribution对象
func (e *FilDistribution) UpdateStatus(c *dto.FilDistributionUpdateStatusReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilDistribution{}
	db := e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())

	if err = db.Error; err != nil {
		e.Log.Errorf("FilDistributionService Update error:%s \r\n", err)
		return err
	}

	err = e.Orm.Table(data.TableName()).Updates(c).Error
	if err != nil {
		e.Log.Errorf("FilDistributionService Update error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FilDistribution
func (e *FilDistribution) Remove(d *dto.FilDistributionDeleteReq, p *actions.DataPermission) error {
	var data models.FilDistribution

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFilDistribution error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// List 获取FilDistribution列表
func (e *FilDistribution) List(c *dto.FilDistributionGetPageReq, list *[]models.FilDistribution) error {
	var err error
	var data models.FilDistribution
	c.Status = 1
	fmt.Printf("req filDistribution:%d", c.GetIds())

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("FilDistributionService List error:%s \r\n", err)
		return err
	}
	return nil
}

// ExportXlsx 导出FilDistribution数据
func (e *FilDistribution) ExportXlsx(r *dto.FilDistributionGetPageReq) error {
	var err error
	var data models.FilDistribution
	var list []models.FilDistribution
	err = e.Orm.Model(&data).Preload("FilNode").
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
		).
		Find(&list).
		Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("FilDistributionService GetPage error:%s \r\n", err)
		return err
	}
	titleList := []string{"矿工", "类型", "可用餘額", "转币数量", "CC分成比例", "合作方分成比例", "上期质押", "本期质押", "参与分币数量=转币数量-(上期质押-本期质押)", "CC分币", "合作方分币", "质押返还", "接收地址", "创建时间"}
	e.saveXlsx(titleList, list, "fil_distribute")

	return nil
}

func (e *FilDistribution) saveXlsx(titleList []string, dataList []models.FilDistribution, fileName string) {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("分币记录")
	// 插入表头
	titleRow := sheet.AddRow()
	myStyle := xlsx.NewStyle()
	myStyle.Alignment = xlsx.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}
	myStyle.Font.Name = "宋体-简"
	myStyle.Font.Size = 11
	myStyle.Border = xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}

	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
		cell.SetStyle(myStyle)

	}
	// 插入内容
	for _, v := range dataList {
		row := sheet.AddRow()
		//row.WriteStruct(v, -1)
		cell := row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Node

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.FilNode.Type

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.0000;[RED]-#0.0000")
		cell.Value = v.AvailableBalance.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.0000;[RED]-#0.0000")
		cell.Value = v.HasTransfer.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.00;[RED]-#0.00")
		cell.Value = v.DistributePoint.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.00;[RED]-#0.00")
		cell.Value = decimal.NewFromInt(1).Sub(v.DistributePoint).String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.0000;[RED]-#0.0000")
		cell.Value = v.LastSectorPledge.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.0000;[RED]-#0.0000")
		cell.Value = v.CurSectorPledge.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.SetFormat("#0.00;[RED]-#0.00")
		cell.Value = v.EffectAmount.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.NumFmt = "0.00"
		cell.Value = v.DistributeAmount.String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.NumFmt = "0.00"
		cell.Value = v.EffectAmount.Sub(v.DistributeAmount).String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.NumFmt = "0.00"
		cell.Value = v.LastSectorPledge.Sub(v.CurSectorPledge).String()

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.AddressTo

		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = dateFormat(v.UpdatedAt)
		cell.NumFmt = "yyyy-MM-dd HH:mm:ss"
	}
	//fileName = fmt.Sprintf("%s-%s.xlsx", fileName, dateFormat(time.Now()))
	//err := file.Save("/Users/tech-air/Documents/Fil_distribute/" + fileName)
	//if err != nil {
	//	panic(err)
	//}
	if err := file.Save(fileName); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fileName:%s\n", fileName)
}

func dateFormat(t time.Time) string {
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := t.Format(timeTemplate1)
	return timeStr
}

func StringToTime(s string) time.Time {
	//loc, _ := time.LoadLocation("Local")
	ctime, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err == nil {
		//fmt.Println(ctime)
		return ctime
	}
	return time.Now()
}
