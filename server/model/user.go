package model

type User struct {
	Username   string `json:"userName"`
	UserPwd    string `json:"userPwd"`
	UserStatus int    `json:"userStatus"`
}
