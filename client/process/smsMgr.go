package process

import (
	"encoding/json"
	"fmt"
	"lesson8/common"
)

func outputGroupMes(mes *common.Message) {
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("sms unmarshal err=", err)
		return
	}

	info := fmt.Sprintf("用戶 %s:\t  %s", smsMes.UserName, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
