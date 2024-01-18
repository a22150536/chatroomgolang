package process

import (
	"encoding/json"
	"fmt"
	"lesson8/client/utils"
	"lesson8/common"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *common.Message) (err error) {
	var smsMes common.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("sms json unmarshal err =", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("sms data json marshal err =", err)
	}

	for Name, up := range userMgr.OnlineUsers {
		if Name == smsMes.UserName {
			continue
		}

		err = this.SendEachOnlineUser(data, up.Conn)
		if err != nil {
			fmt.Println(Name, " sms send content err =", err)
		}
	}
	return
}

func (this *SmsProcess) SendEachOnlineUser(content []byte, Conn net.Conn) (err error) {
	var tf utils.Transfer = utils.Transfer{
		Conn: Conn,
	}

	err = tf.WritePkg(content)
	if err != nil {
		fmt.Println("sms send content err =", err)
		return
	}
	return
}
