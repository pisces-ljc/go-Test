package utils

import (
	"ChatDemo/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// Transfer 将方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //传输时使用缓冲
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn,就不会阻塞
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		fmt.Printf("conn.Read(bytesLen), err = %v\n", err)
		return
	}

	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])

	//根据长度来读取消息内容
	//conn.Read将数据读到切片中去
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(bytes) fail =", err)
	}

	//将数据反序列化-> message.Message
	//注：&
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail =", err)
		return
	}

	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	//1.先发送一个长度给客户端
	// 先获取到data的长度 -> 转化成一个可以发送的[]byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))

	//注：PutUint32函数中类型支持的是（[]byte, uint32）
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)

	//开始发送长度
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(pkglen) fail, err = %v\n", err)
		return
	}

	//发送消息本身
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Printf("conn.Write(data) fail, err = %v\n", err)
		return
	}

	return
}
