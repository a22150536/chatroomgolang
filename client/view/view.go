package view

import (
	"fmt"
	"lesson8/client/process"
)

func View() {
	var option int
	var name, pwd string
	for {
		fmt.Println("-------------------歡迎登陸多人聊天系統---------------------")
		fmt.Printf("\t\t\t\t1 登陸聊天系統 \n")
		fmt.Printf("\t\t\t\t2 註冊用戶 \n")
		fmt.Printf("\t\t\t\t3 退出系統 \n")
		fmt.Print("請選擇(1-3): ")
		fmt.Scan(&option)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("輸入:", option)
		if option == 1 {
			fmt.Printf("\t\t\t\t1 登陸聊天系統 \n")
			fmt.Print("請輸入用戶帳號:")
			fmt.Scan(&name)
			fmt.Print("請輸入用戶密碼:")
			fmt.Scan(&pwd)
			user := process.NewUser(name, pwd)
			err := user.Checkuser()
			if err != nil {
				fmt.Println("登陸失敗 error message:", err)
			} else {
				fmt.Println("登陸成功")
			}
		} else if option == 2 {
			fmt.Printf("\t\t\t\t2 註冊用戶 \n")
			fmt.Print("請輸入用戶帳號:")
			fmt.Scan(&name)
			fmt.Print("請輸入用戶密碼:")
			fmt.Scan(&pwd)
			user := process.NewUser(name, pwd)
			err := user.Registeruser()
			if err != nil {
				fmt.Println("註冊失敗 error message:", err)
			}
		} else if option == 3 {
			fmt.Println("登出")
			break
		} else {
			fmt.Println("輸入錯誤，請輸入數字1-3")
		}
	}
}
