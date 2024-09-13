package jobs

import (
	"encoding/json"
	sysModels "fil-admin/app/admin/models"
	"fil-admin/app/filpool/models"
	"fil-admin/common/service"
	"fil-admin/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

// PushExec struct
type Apns2PushExec struct {
	service.Service
}

func (e Apns2PushExec) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore Apns2PushExec exec success"
	fmt.Println(str, arg.(string))

	db := GetDb()
	if db == nil {
		fmt.Println("db is nil")
		return nil
	}
	e.Orm = db
	list := []models.SendMsg{}
	e.FindNeedSendList(&list)
	if len(list) == 0 {
		log.Printf("No message to send. \r\n")
		return nil
	}
	sendMap := make(map[string][]sysModels.SysUser)
	for _, sendMsg := range list {
		if sendMap[sendMsg.Node] == nil {
			users, err := e.GetUserDeviceToken(sendMsg.Node)
			if err != nil {
				log.Printf("failed to get user device token: %s \r\n", err)
				continue
			}
			sendMap[sendMsg.Node] = users
		}
		users := sendMap[sendMsg.Node]
		e.UpdateSendStatus(&sendMsg)
		for _, user := range users {
			Apns2Push(user.DeviceToken, sendMsg.Title, sendMsg.Content)
		}
	}

	return nil
}

// GetPage 获取SendMsg列表
func (e *Apns2PushExec) FindNeedSendList(list *[]models.SendMsg) error {
	var err error
	var data models.SendMsg

	err = e.Orm.Model(&data).Where("send_status = 0").Order("id").Limit(1).Find(list).Error
	if err != nil {
		e.Log.Errorf("Apns2PushExec FindNeedSendList error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *Apns2PushExec) UpdateSendStatus(sendMsg *models.SendMsg) error {
	err := e.Orm.Model(&models.SendMsg{}).Where("id = ?", sendMsg.Id).Updates(map[string]interface{}{"send_status": 1, "send_time": time.Now()}).Error
	if err != nil {
		e.Log.Errorf("Apns2PushExec UpdateSendStatus error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *Apns2PushExec) GetUserDeviceToken(miner string) ([]sysModels.SysUser, error) {
	node := models.FilNodes{}
	e.Orm.Model(&models.FilNodes{}).Where("node = ?", miner).First(&node)
	if node.Id == 0 {
		log.Printf("Node %s not exist.:\n", miner)
		return nil, nil
	}

	//selSql := fmt.Sprintf("SELECT `sys_user`.`user_id`,`sys_user`.`username`,`sys_user`.`dept_id`,`sys_user`.`device_token`, `sys_user`.`token_status` FROM `sys_user` left join `sys_dept` on `sys_dept`.`dept_id` = `sys_user`.`dept_id` WHERE (`sys_dept`.`dept_path` like '%s' OR `sys_user`.`role_id` = 1) AND `sys_user`.`deleted_at` IS NULL AND `sys_user`.`token_status` = 1;", "%/"+strconv.Itoa(node.DeptId)+"/%")
	var users []sysModels.SysUser
	err := e.Orm.Model(&sysModels.SysUser{}).Where("(dept_id = ? OR role_id = 1) AND token_status = 1 AND device_token IS NOT NULL", node.DeptId).Find(&users).Error
	if err != nil {
		e.Log.Errorf("Apns2PushExec GetUserDeviceToken error:%s \r\n", err)
		return nil, err
	}
	return users, nil
}

type Aps struct {
	Aps Alert `json:"aps"`
}

type Alert struct {
	Alert string `json:"alert"`
}

// Apns2Push 推送消息
func Apns2Push(deviceToken string, title string, content string) {
	certPath := config.ExtConfig.Apns2.CertPath
	topic := config.ExtConfig.Apns2.Topic

	if certPath == "" || deviceToken == "" || topic == "" {
		log.Println("CertPath, DeviceToken, Topic is required")
		return
	}
	password := config.ExtConfig.Apns2.Password
	cert, err := certificate.FromP12File(certPath, password)
	if err != nil {
		log.Println("Cert Error:", err)
		return
	}
	path, _ := os.Getwd()
	log.Default().Printf("Cert Success，path=%s", path)
	if !config.ExtConfig.Apns2.Push {
		log.Println("Apns2 Push is disabled")
		return
	}

	a := Alert{
		Alert: content,
	}
	aps := Aps{
		Aps: a,
	}

	jsonStr, _ := json.Marshal(aps)
	notification := &apns2.Notification{}
	notification.DeviceToken = deviceToken
	notification.Topic = topic
	notification.Payload = []byte(jsonStr) // See Payload section below
	// Use Production() for apps published to the app store or installed as an ad-hoc distribution
	client := apns2.NewClient(cert)
	if config.ExtConfig.Apns2.Prod {
		client = client.Production()
	}

	res, err := client.Push(notification)
	if err != nil {
		log.Println("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

func PathWorkDir(certPath string) string {
	workDir, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", workDir, certPath)
}
