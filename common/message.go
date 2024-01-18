package common

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code     int      `json:"code"`
	Error    string   `json:"error"`
	UserList []string `json:"userlist"`
}

type RegisterMes struct {
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}

type SmsMes struct {
	Content    string `json:"content"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}
