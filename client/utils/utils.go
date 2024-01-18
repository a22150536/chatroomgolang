package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lesson8/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4], uint32(len(data)))
	n, err := this.Conn.Write(bytes[:4])

	if n != 4 || err != nil {
		fmt.Println("conn Write fail err :", err)
		fmt.Printf("首發送 %d 字節", n)
		return
	}

	n, err = this.Conn.Write(data)

	if n != len(data) || err != nil {
		fmt.Println("conn Write data fail err :", err)
		fmt.Printf("發送 %d 字節 %d 字節", n, len(data))
		return
	}

	return
}

func (this *Transfer) ReadPkg() (mes common.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("conn read err = ", err)
		err = errors.New("conn read err ")
		return
	}
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err:", err)
		err = errors.New("conn.Read fail err")
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("un json fail err:", err)
		err = errors.New("un json fail err")
		return
	}

	return
}
