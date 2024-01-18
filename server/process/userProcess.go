package process

import (
	"encoding/json"
	"fmt"
	"lesson8/common"
	"lesson8/server/model"
	"lesson8/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	name string
}

func (this *UserProcess) ServerProcessLogin(mes *common.Message) (err error) {
	var loginMes common.LoginMes
	var resMes common.Message
	var loginResMes common.LoginResMes
	resMes.Type = common.LoginResMesType

	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("login unmarshal fail err=", err)
		return
	}

	user, err := model.MyUserDao.Login(loginMes.UserName, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = error.Error(err)
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = error.Error(err)
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服務器內部錯誤"
		}

	} else {
		loginResMes.Code = 200
		loginResMes.Error = "登入成功"
		this.name = loginMes.UserName
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser()
		for un := range userMgr.OnlineUsers {
			loginResMes.UserList = append(loginResMes.UserList, un)
		}

		fmt.Println(user, "登入成功")
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes marshal fail err=", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("ResMes marshal fail err=", err)
		return
	}
	var tf utils.Transfer = utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

func (this *UserProcess) ServerProcessRegister(mes *common.Message) (err error) {
	var registerMes common.RegisterMes
	var resMes common.Message
	var registerResMes common.RegisterResMes
	resMes.Type = common.RegisterResMesType

	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("register unmarshal fail err=", err)
		return
	}

	err = model.MyUserDao.Register(registerMes.UserName, registerMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 505
			registerResMes.Error = "註冊失敗，用戶已存在"
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "服務器內部錯誤"
		}
	} else {
		registerResMes.Code = 200
		registerResMes.Error = "註冊成功"
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("registerResMes marshal fail err=", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("registerResMes marshal fail err=", err)
		return
	}
	var tf utils.Transfer = utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)

	return
}

func (this *UserProcess) NotifyOthersOnlineUser() {

	for name, up := range userMgr.GetAllOnlineUser() {
		if this.name == name {
			continue
		}
		up.NotifyMeOnline(this.name)
	}
	return
}

func (this *UserProcess) NotifyMeOnline(user string) {
	var mes common.Message
	var notifyUserStatusMes common.NotifyUserStatusMes
	mes.Type = common.NotifyUserStatusMesType
	notifyUserStatusMes.UserName = user
	notifyUserStatusMes.UserStatus = common.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notify json marshal err:", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("notify mes json marshal err:", err)
		return
	}

	var tf utils.Transfer = utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notify writePkg fail err:", err)
		return
	}
}

func (this *UserProcess) NotifyOthersOfflineUser() {

	for name, up := range userMgr.GetAllOnlineUser() {
		if this.name == name {
			fmt.Println(name, "尚未清除連線信息")
			continue
		}
		up.NotifyMeOffline(this.name)
	}
	return
}

func (this *UserProcess) NotifyMeOffline(user string) {
	var mes common.Message
	var notifyUserStatusMes common.NotifyUserStatusMes
	mes.Type = common.NotifyUserStatusMesType
	notifyUserStatusMes.UserName = user
	notifyUserStatusMes.UserStatus = common.UserOffline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notify json marshal err:", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("notify mes json marshal err:", err)
		return
	}

	var tf utils.Transfer = utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notify writePkg fail err:", err)
		return
	}
}
