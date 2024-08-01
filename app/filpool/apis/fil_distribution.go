package apis

import (
	"fil-admin/app/filpool/models"
	"fil-admin/app/filpool/service"
	"fil-admin/app/filpool/service/dto"
	"fil-admin/common/actions"
	cModels "fil-admin/common/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/shopspring/decimal"
)

type FilDistribution struct {
	api.Api
}

// GetPage 获取挖矿收益分币记录列表
// @Summary 获取挖矿收益分币记录列表
// @Description 获取挖矿收益分币记录列表
// @Tags 挖矿收益分币记录
// @Param node query string false "节点名称"
// @Param addressFrom query string false "发送地址"
// @Param addressTo query string false "接收地址"
// @Param status query string false "分币状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.FilDistribution}} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-distribution [get]
// @Security Bearer
func (e FilDistribution) GetPage(c *gin.Context) {
	req := dto.FilDistributionGetPageReq{}
	s := service.FilDistribution{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.FilDistribution, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取挖矿收益分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取挖矿收益分币记录
// @Summary 获取挖矿收益分币记录
// @Description 获取挖矿收益分币记录
// @Tags 挖矿收益分币记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.FilDistribution} "{"code": 200, "data": [...]}"
// @Router /api/v1/fil-distribution/{id} [get]
// @Security Bearer
func (e FilDistribution) Get(c *gin.Context) {
	req := dto.FilDistributionGetReq{}
	s := service.FilDistribution{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.FilDistribution

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取挖矿收益分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建挖矿收益分币记录
// @Summary 创建挖矿收益分币记录
// @Description 创建挖矿收益分币记录
// @Tags 挖矿收益分币记录
// @Accept application/json
// @Product application/json
// @Param data body dto.FilDistributionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/fil-distribution [post]
// @Security Bearer
func (e FilDistribution) Insert(c *gin.Context) {
	req := dto.FilDistributionInsertReq{}
	s := service.FilDistribution{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建挖矿收益分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改挖矿收益分币记录
// @Summary 修改挖矿收益分币记录
// @Description 修改挖矿收益分币记录
// @Tags 挖矿收益分币记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.FilDistributionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/fil-distribution/{id} [put]
// @Security Bearer
func (e FilDistribution) Update(c *gin.Context) {
	req := dto.FilDistributionUpdateReq{}
	s := service.FilDistribution{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改挖矿收益分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除挖矿收益分币记录
// @Summary 删除挖矿收益分币记录
// @Description 删除挖矿收益分币记录
// @Tags 挖矿收益分币记录
// @Param data body dto.FilDistributionDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/fil-distribution [delete]
// @Security Bearer
func (e FilDistribution) Delete(c *gin.Context) {
	s := service.FilDistribution{}
	req := dto.FilDistributionDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除挖矿收益分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// AddDistribute 生成分币记录
func (e FilDistribution) AddDistribute(c *gin.Context) {
	s := service.FilDistribution{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s2 := service.FilNodes{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s2.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)

	var unionList []models.FilNodes
	s2.FindUnionNode(&unionList)

	//查询节点的未分币记录
	var unSendList []models.FilDistribution
	s.List(&dto.FilDistributionGetPageReq{}, &unSendList)
	var m = make(map[string]models.FilDistribution)
	for _, distribute := range unSendList {
		m[distribute.Node] = distribute
	}

	for _, oneNode := range unionList {
		oneNode.AvailableBalance = DecimalRoundCountDownValue(oneNode.AvailableBalance, 0)
		effectAmount := oneNode.HasTransfer.Sub(oneNode.LastDisSectorPledgeBalance).Add(oneNode.SectorPledgeBalance)
		if distribute, ok := m[oneNode.Node]; ok {
			//更新
			updateReq := dto.FilDistributionUpdateReq{
				Id:               distribute.Id,
				NodeId:           oneNode.Id,
				Node:             oneNode.Node,
				AvailableBalance: oneNode.AvailableBalance,
				HasTransfer:      oneNode.HasTransfer,
				DistributePoint:  oneNode.DistributePoint,
				LastSectorPledge: oneNode.LastDisSectorPledgeBalance,
				CurSectorPledge:  oneNode.SectorPledgeBalance,
				EffectAmount:     effectAmount,
				DistributeAmount: effectAmount.Mul(oneNode.DistributePoint),
				AddressFrom:      "",
				AddressTo:        "",
				Status:           1,
			}
			updateReq.SetUpdateBy(user.GetUserId(c))
			err = s.Update(&updateReq, p)
		} else {
			//新增
			insertReq := dto.FilDistributionInsertReq{
				NodeId:           oneNode.Id,
				Node:             oneNode.Node,
				AvailableBalance: oneNode.AvailableBalance,
				HasTransfer:      oneNode.HasTransfer,
				DistributePoint:  oneNode.DistributePoint,
				LastSectorPledge: oneNode.LastDisSectorPledgeBalance,
				CurSectorPledge:  oneNode.SectorPledgeBalance,
				EffectAmount:     effectAmount,
				DistributeAmount: effectAmount.Mul(oneNode.DistributePoint),
				AddressFrom:      "",
				AddressTo:        "",
				Status:           1,
				ControlBy:        cModels.ControlBy{},
			}
			insertReq.SetCreateBy(user.GetUserId(c))
			err = s.Insert(&insertReq)
		}
	}
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建挖矿收益分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK("", "成功生成分币记录")
}

// DealDistribute 处理分币，更改分币记录状态，同时更新fil_node表数据
func (e FilDistribution) DealDistribute(c *gin.Context) {
	s := service.FilDistribution{}
	req := dto.FilDistributionDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)

	// 查询FilNode记录
	nodeS := service.FilNodes{}
	//nodeReq := dto.FilNodesUpdateReq{
	//	Ids: req.NodeIds,
	//}
	err = e.MakeContext(c).
		MakeOrm().
		//Bind(&nodeReq).
		MakeService(&nodeS.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var unionList []models.FilNodes
	nodeS.FindNeedDealDistributeNode(&unionList, req.Ids)
	//if len(unionList) > 0 {
	//	fmt.Printf("req.NodeIds:%d \n", req.GetId())
	//	e.OK(req.GetId(), strconv.Itoa(len(req.Ids)))
	//	return
	//}

	listReq := dto.FilDistributionGetPageReq{
		NodeIds: req.Ids,
	}
	var needSendList []models.FilDistribution
	s.List(&listReq, &needSendList)

	//if len(needSendList) > 0 {
	//	fmt.Printf("needSendList[0].Node:%d \n", len(needSendList))
	//	e.OK(req.GetId(), strconv.Itoa(len(needSendList)))
	//	return
	//}
	var m = make(map[string]models.FilNodes)
	for _, node := range unionList {
		m[node.Node] = node
	}

	if len(needSendList) == 0 {
		e.OK(req.GetId(), "当前没有需要分币的记录")
		return
	}

	for _, oneDis := range needSendList {
		updateReq := dto.FilDistributionUpdateStatusReq{
			Id:     oneDis.Id,
			Status: 2,
		}
		updateReq.SetUpdateBy(user.GetUserId(c))
		s.UpdateStatus(&updateReq, p)

		node := m[oneDis.Node]
		updateNodeReq := dto.FilNodesUpdateDistributeReq{
			Id:                         oneDis.NodeId,
			HasRealDistribute:          node.HasRealDistribute.Add(oneDis.DistributeAmount),
			LastDisSectorPledgeBalance: oneDis.CurSectorPledge,
			LastDistributeTime:         time.Now(),
			HasTransfer:                decimal.Zero,
		}
		updateNodeReq.SetUpdateBy(user.GetUserId(c))
		nodeS.UpdateDistribute(&updateNodeReq, p)
	}

	e.OK(req.GetId(), "分币记录更新成功")
}

func (e FilDistribution) ExportXlsx(c *gin.Context) {
	req := dto.FilDistributionGetPageReq{}
	s := service.FilDistribution{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.ExportXlsx(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("导出分币记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK("", "创建成功")
}

// DecimalRoundCountDownValue 保留到百位
func DecimalRoundCountDownValue(dec decimal.Decimal, round int32) decimal.Decimal {
	//value := decimal.NewFromFloat(math.Pow10(1))
	v := dec.RoundDown(round)
	return v
}
