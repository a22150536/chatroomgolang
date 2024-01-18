package process

import (
	"fmt"
	"lesson8/client/model"
	"lesson8/common"
)

type UserListStatus struct {
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}

var onlineUsers map[string]*UserListStatus = make(map[string]*UserListStatus, 10)
var CurUser model.CurUser

func updateUserStatus(notifyUserStatusMes common.NotifyUserStatusMes) {

	if notifyUserStatusMes.UserStatus == common.UserOnline {
		user, ok := onlineUsers[notifyUserStatusMes.UserName]

		if !ok {
			user = &UserListStatus{
				UserName: notifyUserStatusMes.UserName,
			}
		}
		user.UserStatus = notifyUserStatusMes.UserStatus

		onlineUsers[notifyUserStatusMes.UserName] = user
	} else if notifyUserStatusMes.UserStatus == common.UserOffline {
		_, ok := onlineUsers[notifyUserStatusMes.UserName]
		if ok {
			delete(onlineUsers, notifyUserStatusMes.UserName)
		} else {
			fmt.Println("上線列表沒有", notifyUserStatusMes.UserName)
		}
	} else {
		fmt.Println("UserStatus錯誤")
	}

}

func outputOnlineUser() {
	fmt.Println("當前在線用戶列表:")
	for Name := range onlineUsers {
		fmt.Println("用戶帳號:", Name)
	}
}
