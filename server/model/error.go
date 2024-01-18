package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("用戶不存在....")
	ERROR_USER_EXIST     = errors.New("用戶已經存在....")
	ERROR_USER_PWD       = errors.New("密碼不正確")
)
