package service

import (
	"errors"
	"fil-admin/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx"
	"gorm.io/gorm"

	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"

	cDto "fil-admin/common/dto"
	"log"
)

type FilNodes struct {
	service.Service
}

// GetPage 获取FilNodes列表
func (e *FilNodes) GetPage(c *dto.FilNodesGetPageReq, p *actions.DataPermission, list *[]models.FilNodes, count *int64) error {
	var err error
	var data models.FilNodes

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// GetAll 获取FilNodes列表 所有符合条件的记录，不分页
func (e *FilNodes) GetAll(c *dto.FilNodesGetPageReq, p *actions.DataPermission, list *[]models.FilNodes, count *int64) error {
	var err error
	var data models.FilNodes

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Where("quality_adj_power > 0").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("GetAll error:%s \r\n", err)
		return err
	}
	return nil
}

// RankList 获取FilNodes列表
func (e *FilNodes) RankList(p *actions.DataPermission, list *[]models.FilNodes) error {
	var err error
	var data models.FilNodes

	err = e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Order("quality_adj_power DESC").
		Find(list).Error
	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// ControlList 获取FilNodes列表
func (e *FilNodes) ControlList(p *actions.DataPermission, list *[]models.FilNodes) error {
	var err error
	var data models.FilNodes

	err = e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Where("quality_adj_power > 0").Order("control_balance").
		Find(list).Error
	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取FilNodes对象
func (e *FilNodes) Get(d *dto.FilNodesGetReq, p *actions.DataPermission, model *models.FilNodes) error {
	var data models.FilNodes

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFilNodes error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建FilNodes对象
func (e *FilNodes) Insert(c *dto.FilNodesInsertReq) error {
	var err error
	var data = new(models.FilNodes)
	c.Generate(data)
	err = e.Orm.Create(&c).Error
	if err != nil {
		e.Log.Errorf("FilNodesService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改FilNodes对象
func (e *FilNodes) Update(c *dto.FilNodesUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilNodes{}
	db := e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("FilNodesService Update error:%s \r\n", err)
		return err
	}

	err = e.Orm.Table(data.TableName()).Updates(c).Error
	if err != nil {
		e.Log.Errorf("FilNodesService Update error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// UpdateTitle 修改FilNodes对象
func (e *FilNodes) UpdateTitle(c *dto.FilNodesUpdateTitleReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilNodes{}
	db := e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("FilNodesService Update error:%s \r\n", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Model(&data).Select("title", "update_by", "updated_at").Updates(data).Error
	if err != nil {
		e.Log.Errorf("FilNodesService Update error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// UpdateDistribute 修改FilNodes对象
func (e *FilNodes) UpdateDistribute(c *dto.FilNodesUpdateDistributeReq, p *actions.DataPermission) error {
	var err error
	var data = models.FilNodes{}
	db := e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("FilNodesService Update error:%s \r\n", err)
		return err
	}

	err = e.Orm.Table(data.TableName()).Updates(c).Error
	if err != nil {
		e.Log.Errorf("FilNodesService Update error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除FilNodes
func (e *FilNodes) Remove(d *dto.FilNodesDeleteReq, p *actions.DataPermission) error {
	var data models.FilNodes

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFilNodes error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// FindUnionNode 查询联合挖矿节点
func (e *FilNodes) FindUnionNode(list *[]models.FilNodes) error {
	var data models.FilNodes

	db := e.Orm.Model(&data).Where("balance > 0 AND distribute_point > 0 AND status > 0").Order("type").Find(list)
	if err := db.Error; err != nil {
		e.Log.Errorf("FindUnionNode error:%s \r\n", err)
		return err
	}
	return nil
}

// FindNeedDealDistributeNode 查询需要更新分币数据的节点
func (e *FilNodes) FindNeedDealDistributeNode(list *[]models.FilNodes, ids []int) error {
	var data models.FilNodes

	db := e.Orm.Model(&data).Where("id In (?)", ids).Order("id").Find(list)
	if err := db.Error; err != nil {
		e.Log.Errorf("FindNodeInIds error:%s \r\n", err)
		return err
	}
	return nil
}

// ExportXlsx 导出FilNode数据
func (e *FilNodes) ExportXlsx(r *dto.FilNodesGetPageReq, c *gin.Context) error {
	var err error
	var data models.FilNodes
	var list []models.FilNodes
	err = e.Orm.Model(&data).Preload("NodeChart").
		Scopes(
			cDto.MakeCondition(r.GetNeedSearch()),
		).Order("type, id").
		Find(&list).
		Limit(-1).Offset(-1).Error
	if err != nil {
		e.Log.Errorf("FilNodesService GetPage error:%s \r\n", err)
		//return err
	}

	titleList := []string{"Miner", "Owner", "類型", "收益比例", "算力(PiB)", "昨日算力(PiB)", "月初算力(PiB)", "24H算力增量(PiB)", "當月算力增量(PiB)", "可用餘額", "昨日可用", "總餘額", "昨日餘額", "當前質押", "昨日質押", "月初質押", "當月質押釋放", "存儲鎖倉", "昨日鎖倉", "24H爆塊數", "24H獎勵", "24H幸運值", "總爆塊數", "總獎勵", "昨日爆塊數", "昨日爆塊獎勵", "當月爆塊數", "當月爆塊獎勵", "昨日實際收益爆塊數", "昨日實際收益爆塊獎勵", "當月實際收益爆塊數", "當月實際收益爆塊獎勵", "充值總數量", "销毁總數量", "提現總數量", "昨日充值", "昨日銷燬", "昨日提現", "當月充值", "當月銷燬", "當月提現", "控制地址余额", "有效扇區", "錯誤扇區", "扇区大小", "创建时间", "结束时间"}
	e.saveXlsx(titleList, list, "fil_nodes")

	return nil
}

func (e *FilNodes) saveXlsx(titleList []string, dataList []models.FilNodes, fileName string) {
	// 打开文件
	// xlsx.NewFile() //
	// file = "/Users/hello/Documents/Fil_nodes/"
	//file, _ := xlsx.OpenFile("/Users/hello/Documents/Fil_nodes/fil_nodes.xlsx")
	file := xlsx.NewFile()
	// 添加sheet页
	//sheetName := fmt.Sprintf("%s", utils.DateFormatStr(time.Now(), "2006-01-02 15-04"))
	//fmt.Printf("list.size:%d\n", len(dataList))
	//fmt.Printf("sheetName:%s\n", sheetName)
	sheet, _ := file.AddSheet("data")
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

	timeTemplate1 := "2006-01-02 15:04:05"

	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
		cell.SetStyle(myStyle)

	}
	// 插入内容
	for _, n := range dataList {
		row := sheet.AddRow()
		//row.WriteStruct(v, -1)
		for _, v := range titleList {
			cell := row.AddCell()
			cell.SetStyle(myStyle)

			switch {
			case v == "Miner":
				cell.Value = n.Node
			case v == "Owner":
				cell.Value = n.MsigNode
			case v == "類型":
				if n.Type == "1" || n.Type == "4" || n.Type == "5" {
					cell.Value = "自有CC"
				} else if n.Type == "2" {
					cell.Value = "英链"
					if n.DistributePoint.String() == "0.2" {
						cell.Value = "Web3 PLUS"
					}
				} else if n.Type == "3" {
					cell.Value = "蜻蜓"
				} else if n.Type == "6" {
					cell.Value = "PLUS"
				} else {
					cell.Value = n.Type
				}
			case v == "收益比例":
				cell.Value = n.DistributePoint.String()
			case v == "算力(PiB)":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.QualityAdjPower.String()
			case v == "昨日算力(PiB)":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastQualityAdjPower.String()
			case v == "月初算力(PiB)":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastMonthQualityAdjPower.String()
			case v == "24H算力增量(PiB)":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.QualityAdjPower.Sub(n.NodeChart.LastQualityAdjPower).String()
			case v == "當月算力增量(PiB)":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.QualityAdjPower.Sub(n.NodeChart.LastMonthQualityAdjPower).String()
			case v == "可用餘額":
				cell.Value = n.AvailableBalance.String()
			case v == "昨日可用":
				cell.Value = n.NodeChart.LastAvailableBalance.String()
			case v == "總餘額":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.Balance.String()
			case v == "昨日餘額":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastBalance.String()
			case v == "當前質押":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.SectorPledgeBalance.String()
			case v == "昨日質押":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastSectorPledgeBalance.String()
			case v == "月初質押":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastMonthSectorPledgeBalance.String()
			case v == "當月質押釋放":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastMonthSectorPledgeBalance.Sub(n.SectorPledgeBalance).String()
			case v == "存儲鎖倉":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.VestingFunds.String()
			case v == "昨日鎖倉":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.LastVestingFunds.String()
			case v == "24H爆塊數":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = strconv.Itoa(n.BlocksMined24h)
			case v == "24H獎勵":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.TotalRewards24h.String()
			case v == "24H幸運值":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.LuckyValue24h.Mul(decimal.NewFromInt(100)).String()
			case v == "總爆塊數":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = strconv.Itoa(n.WeightedBlocks)
			case v == "總獎勵":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.RewardValue.String()
			case v == "昨日爆塊數":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = strconv.Itoa(n.NodeChart.WeightedBlocks - n.NodeChart.LastWeightedBlocks)
			case v == "昨日爆塊獎勵":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.RewardValue.Sub(n.NodeChart.LastRewardValue).String()
			case v == "當月爆塊數":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = strconv.Itoa(n.NodeChart.WeightedBlocks - n.NodeChart.LastMonthWeightedBlocks)
			case v == "當月爆塊獎勵":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.RewardValue.Sub(n.NodeChart.LastMonthRewardValue).String()
			case v == "昨日實際收益爆塊數":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = utils.DecimalValue(strconv.Itoa(n.NodeChart.WeightedBlocks - n.NodeChart.LastWeightedBlocks)).Mul(n.DistributePoint).String()
			case v == "昨日實際收益爆塊獎勵":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.RewardValue.Sub(n.NodeChart.LastRewardValue).Mul(n.DistributePoint).String()
			case v == "當月實際收益爆塊數":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = utils.DecimalValue(strconv.Itoa(n.NodeChart.WeightedBlocks - n.NodeChart.LastMonthWeightedBlocks)).Mul(n.DistributePoint).String()
			case v == "當月實際收益爆塊獎勵":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.NodeChart.RewardValue.Sub(n.NodeChart.LastMonthRewardValue).Mul(n.DistributePoint).String()
			case v == "充值總數量":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.ReceiveAmount.String()
			case v == "销毁總數量":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.BurnAmount.String()
			case v == "提現總數量":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.SendAmount.String()
			case v == "昨日充值":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.ReceiveAmount.Sub(n.NodeChart.LastReceiveAmount).String()
			case v == "昨日銷燬":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.BurnAmount.Sub(n.NodeChart.LastBurnAmount).String()
			case v == "昨日提現":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.SendAmount.Sub(n.NodeChart.LastSendAmount).String()
			case v == "當月充值":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.ReceiveAmount.Sub(n.NodeChart.LastMonthReceiveAmount).String()
			case v == "當月銷燬":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.BurnAmount.Sub(n.NodeChart.LastMonthBurnAmount).String()
			case v == "當月提現":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.SendAmount.Sub(n.NodeChart.LastMonthSendAmount).String()
			case v == "控制地址余额":
				cell.SetFormat("#0.0000;[RED]-#0.0000")
				cell.Value = n.ControlBalance.String()
			case v == "有效扇區":
				cell.Value = strconv.Itoa(n.SectorEffective)
			case v == "錯誤扇區":
				cell.Value = strconv.Itoa(n.SectorError)
			case v == "扇区大小":
				cell.Value = n.SectorSize
			case v == "创建时间":
				// cell.NumFmt = "yyyy-MM-dd HH:mm:ss"
				cell.Value = utils.DateFormatStr(n.CreateTime, timeTemplate1)
			case v == "结束时间":
				//cell.NumFmt = "yyyy-MM-dd"
				cell.Value = utils.DateFormatStr(n.EndTime, "2006-01-02")
			}
		}

	}
	//fileName = fmt.Sprintf("%s-%s.xlsx", fileName, dateFormat(time.Now()))
	//err := file.Save("/Users/tech-air/Documents/Fil_nodes/" + fileName)
	// 将文件保存为临时文件
	//tempFile := "temp.xlsx"
	if err := file.Save(fileName); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fileName:%s\n", fileName)

}

// func exportToExcel(w http.ResponseWriter, r *http.Request) {
// 	// 创建一个新的Excel文件
// 	f := excelize.NewFile()

// 	// 添加数据到表格中
// 	f.SetCellValue("Sheet1", "A1", "ID")
// 	f.SetCellValue("Sheet1", "B1", "Name")
// 	f.SetCellValue("Sheet1", "A2", "1")
// 	f.SetCellValue("Sheet1", "B2", "John Doe")
// 	f.SetCellValue("Sheet1", "A3", "2")
// 	f.SetCellValue("Sheet1", "B3", "Jane Smith")

// 	// 将文件保存为临时文件
// 	tempFile := "temp.xlsx"
// 	if err := f.SaveAs(tempFile); err != nil {
// 		log.Fatal(err)
// 	}

// 	// 设置响应头告诉浏览器文件的类型
// 	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
// 	// 设置响应头告诉浏览器文件的名称
// 	w.Header().Set("Content-Disposition", "attachment; filename=table.xlsx")

// 	// 将文件内容写入响应体中
// 	http.ServeFile(w, r, tempFile)
// }

// SumBlockStats 获取矿池的当天报块统计
func (e *FilNodes) SumBlockStats(nodes []string, lastDay time.Time, list *[]models.BlockStats) error {

	err := e.Orm.Model(&models.BlockStats{}).
		Select("height_time, height_time_str, SUM(Blocks_growth) as blocks_growth, SUM(blocks_reward_growth_fil) as blocks_reward_growth_fil").
		Where("height_time >= ? AND node in (?)", lastDay, nodes).
		Group("height_time, height_time_str").
		Find(list).Error

	if err != nil {
		e.Log.Errorf("SumBlockStats error:%s \r\n", err)
		return err
	}
	return nil
}
