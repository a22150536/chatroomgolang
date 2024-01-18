package main

import (
	"fmt"
	"lesson8/server/model"
	"lesson8/server/process"
	"net"
	"time"

	"github.com/redis/go-redis"
)

var rclient *redis.Client

func init() {
	initPool("localhost:6379", 16, 0, 15, 300*time.Second)
	initUserDao()
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(rclient)
}

func initPool(address string, maxIdle, maxActive, poolSize int, idleTimeout time.Duration) {
	rclient = redis.NewClient(&redis.Options{
		Addr:            address,
		Password:        "",
		DB:              0,
		PoolSize:        poolSize,
		ConnMaxIdleTime: idleTimeout,
		MaxActiveConns:  maxActive,
		MaxIdleConns:    maxIdle,
	})

}

func main() {
	defer rclient.Close()
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("開啟服務異常 err message:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("連接失敗 err message:", err)
		}
		var up process.Processor = process.Processor{
			Conn: conn,
		}
		go up.Process()
	}
}
