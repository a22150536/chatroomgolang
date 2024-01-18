package process

import (
	"errors"
	"fmt"
	"io"
	"lesson8/common"
	"lesson8/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) Process() {

	defer this.Conn.Close()
	for {
		var tf utils.Transfer = utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println(this.Conn.RemoteAddr().String(), " 客戶端退出，服務器端也退出")

			}
			fmt.Println("read fail package err:", err)
			name, err := userMgr.DelOnlineUserByConn(this.Conn)
			if err != nil {

				return
			}

			up := UserProcess{
				Conn: this.Conn,
				name: name,
			}
			up.NotifyOthersOfflineUser()
			return
		}

		err = this.ServerProcessMes(&mes)
	}

}

func (this *Processor) ServerProcessMes(mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		var userP UserProcess = UserProcess{
			Conn: this.Conn,
		}
		err = userP.ServerProcessLogin(mes)
	case common.RegisterMesType:
		var userP UserProcess = UserProcess{
			Conn: this.Conn,
		}
		err = userP.ServerProcessRegister(mes)
	case common.SmsMesType:
		var userP SmsProcess = SmsProcess{}
		err = userP.SendGroupMes(mes)
	default:
		fmt.Println("類型不存在")
		err = errors.New("type not exit")
		return
	}
	return
}
