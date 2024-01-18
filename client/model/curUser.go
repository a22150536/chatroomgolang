package model

import (
	"net"
)

type CurUser struct {
	Conn       net.Conn
	UserName   string
	UserStatus int
}
