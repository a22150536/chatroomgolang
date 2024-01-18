package process

import (
	"encoding/json"
	"fmt"
	"lesson8/client/utils"
	"lesson8/common"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {

	var mes common.Message
	mes.Type = common.SmsMesType

	var smsMes common.SmsMes
	smsMes.Content = content
	smsMes.UserName = CurUser.UserName
	smsMes.UserStatus = CurUser.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json marshal err :", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json marshal err :", err)
		return
	}

	var tf utils.Transfer = utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg err:", err)
		return
	}

	return
}
