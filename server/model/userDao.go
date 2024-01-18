package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis"
)

var (
	MyUserDao *UserDao
)
var ctx = context.Background()

type UserDao struct {
	rclient *redis.Client
}

func NewUserDao(pool *redis.Client) (userDao *UserDao) {
	userDao = &UserDao{
		rclient: pool,
	}
	return
}

func (this *UserDao) getUserByName(name string) (user *User, err error) {
	res, err := this.rclient.HGet(ctx, "users", name).Result()
	if err != nil {

		fmt.Println("redis conn  err = ", err)
		return
	}

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("redis res unmarshal err = ", err)
		return
	}

	return
}

func (this *UserDao) Login(name string, pwd string) (user *User, err error) {

	user, err = this.getUserByName(name)

	if err != nil {
		return
	}

	if user.UserPwd != pwd {
		err = ERROR_USER_PWD
		return
	}
	return

}

func (this *UserDao) setUserByName(name string, user string) (err error) {

	err = this.rclient.HSet(ctx, "users", name, user).Err()

	if err != nil {
		fmt.Println("保存註冊用戶錯誤 err:", err)
		return
	}
	return
}

func (this *UserDao) Register(name string, pwd string) (err error) {

	_, err = this.getUserByName(name)
	if err == nil {
		err = ERROR_USER_EXIST
		fmt.Println("redis conn  err = ", err)
		return
	}

	var user *User = &User{
		Username: name,
		UserPwd:  pwd,
	}

	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println("register user json marshal err: ", err)
		return
	}

	err = this.setUserByName(name, string(data))

	return
}
