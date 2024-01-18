package process

import (
	"errors"
	"fmt"
	"net"
)

var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[string]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[string]*UserProcess, 1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess) {

	this.OnlineUsers[up.name] = up

}

func (this *UserMgr) DelOnlineUserByConn(Conn net.Conn) (name string, err error) {

	for name, up := range this.OnlineUsers {
		if Conn == up.Conn {
			delete(this.OnlineUsers, name)
			return name, nil
		}
	}
	err = errors.New("no user")
	fmt.Println("沒有找到客戶")
	return
}

func (this *UserMgr) GetAllOnlineUser() map[string]*UserProcess {
	return this.OnlineUsers
}

func (this *UserMgr) GetOnlineUserByName(name string) (up *UserProcess, err error) {
	up, ok := this.OnlineUsers[name]
	if !ok {
		err = fmt.Errorf("用戶 %s 不存在", name)
		return
	}
	return
}
