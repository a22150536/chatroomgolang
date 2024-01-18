package process

import (
	"encoding/json"
	"fmt"
	"lesson8/client/utils"
	"lesson8/common"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes() {

	var tf utils.Transfer = utils.Transfer{
		Conn: this.Conn,
	}

Loop:
	for {

		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println(" server process mes read pkg err:", err)
			return
		}

		switch mes.Type {
		case common.NotifyUserStatusMesType:
			fmt.Println("有用戶上線")
			var notifyUserStatusMes common.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("notifyUserStatusMes unmarshal err:", err)
				break Loop
			}
			updateUserStatus(notifyUserStatusMes)
			outputOnlineUser()
			ShowLoginMenu()

		case common.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("server process mes Type unknown")
		}
	}

}
