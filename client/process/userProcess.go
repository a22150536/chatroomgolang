package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"lesson8/client/utils"
	"lesson8/common"
	"net"
	"os"
)

type user struct {
	Name     string
	Password string
}

func NewUser(name string, pwd string) user {
	var user user
	user.Name = name
	user.Password = pwd
	return user
}

func (this *user) Checkuser() (err error) {
	var mes common.Message
	var loginMes common.LoginMes
	var utl utils.Transfer

	conn, err := net.Dial("tcp", "localhost:8889")
	utl.Conn = conn
	defer conn.Close()
	if err != nil {
		fmt.Println("ner.Dial失敗 err: ", err)
		return
	}

	mes.Type = common.LoginMesType
	loginMes.UserName = this.Name
	loginMes.UserPwd = this.Password
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes marshal失敗 err: ", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes marshal失敗 err: ", err)
		return
	}
	err = utl.WritePkg(data)

	if err != nil {
		fmt.Println("write pkg err =", err)
		return
	}

	mes, err = utl.ReadPkg()

	if err != nil {
		fmt.Println("conn Read data fail err :", err)
		return
	}

	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if err != nil {
		fmt.Println("json unmarshal data fail err :", err)
		return
	}

	if loginResMes.Code == 200 {
		CurUser.Conn = conn
		CurUser.UserName = this.Name
		CurUser.UserStatus = common.UserOnline
		fmt.Println("當前用戶列表如下:")
		for _, un := range loginResMes.UserList {
			fmt.Println(un)
			user := &UserListStatus{
				UserName:   this.Name,
				UserStatus: common.UserOnline,
			}
			onlineUsers[un] = user
		}
		fmt.Println()
		var pr Processor = Processor{
			Conn: conn,
		}
		go pr.ServerProcessMes()
		fmt.Printf("--------------恭喜%s登陸成功------------------\n", this.Name)
		for {
			ShowLoginMenu()
		}

	} else if loginResMes.Code == 500 {
		loginResMes.Error = "該用戶不存在，請註冊在使用"
		err = errors.New("該用戶不存在，請註冊在使用")
		return

	}
	return
}

func (this *user) Registeruser() (err error) {
	var mes common.Message
	var registerMes common.RegisterMes
	var utl utils.Transfer

	fmt.Println("註冊帳號:", this.Name)
	fmt.Println("註冊密碼:", this.Password)
	fmt.Println("註冊確認中")
	conn, err := net.Dial("tcp", "localhost:8889")
	utl.Conn = conn
	defer conn.Close()

	if err != nil {
		fmt.Println("ner.Dial失敗 err: ", err)
		return
	}
	fmt.Println("本機IP: " + conn.LocalAddr().String())

	mes.Type = common.RegisterMesType
	registerMes.UserName = this.Name
	registerMes.UserPwd = this.Password

	data, err := json.Marshal(registerMes)
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes marshal失敗 err: ", err)
		return
	}

	err = utl.WritePkg(data)

	if err != nil {
		fmt.Println("write pkg err =", err)
		return
	}

	mes, err = utl.ReadPkg()

	if err != nil {
		fmt.Println("conn Read data fail err :", err)
		return
	}

	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if err != nil {
		fmt.Println("register json unmarshal data fail err :", err)
		return
	}
	fmt.Println("json unmarshal register =", registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("註冊成功")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return
}

func ShowLoginMenu() {

	fmt.Println("--------------1. 顯示在線用戶列表------------------")
	fmt.Println("--------------2. 發送消息------------------")
	fmt.Println("--------------3. 退出系統------------------")
	fmt.Println("請選擇(1-3):")
	var key int
	var smsProcess *SmsProcess = &SmsProcess{}
	var content string
	fmt.Scan(&key)
	switch key {
	case 1:
		fmt.Println("顯示在線用戶列表")
		outputOnlineUser()
	case 2:

		fmt.Println("發送消息")
		fmt.Print("請輸入信息:")
		fmt.Scan(&content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("退出系統")
		os.Exit(0)
	default:
		fmt.Println("輸入錯誤，請輸入數字1-3")
		return
	}
}
