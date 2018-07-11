package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"time"
	"unsafe"
)

var queue chan Message_Block
var lock chan int

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func recvUDPMsg(conn *net.UDPConn) {
	headlen := (int)(unsafe.Sizeof(Record_Req{}))

	for true {
		var buf [MAX_PACK_LENGTH]byte
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println("ReadFromUDP error")
			return
		}

		if n < headlen || n >= MAX_PACK_LENGTH {
			fmt.Println("ReadFromUDP size error")
			continue
		}

		stReq := (*Record_Req)(unsafe.Pointer(&buf))

		// fmt.Println("Ver:%d, cmd:%d", stReq.ver, stReq.cmd)

		if stReq.ver != VERSION_V1 {
			fmt.Println("ReadFromUDP version error")
			continue
		}

		l := stReq.leng[:]
		pkglen := (int)(binary.BigEndian.Uint32(l))
		//fmt.Println("length", pkglen)
		if pkglen != n-headlen {
			fmt.Println(n, headlen, pkglen)
			continue
		}

		msg := new(Message_Block)
		msg.cmd = stReq.cmd
		msg.buff = buf[headlen:n]

		//fmt.Println("produce msg is ", string(buf[6:n]))

		Producer(*msg, queue)

	}

	//WriteToUDP
	//func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	//_, err := conn.WriteToUDP([]byte("nice to see u"), raddr)
	//checkError(err)
}

func Producer(item Message_Block, queue chan Message_Block) {
	queue <- item
}

func Consumer(queue chan Message_Block, lock chan int) {
	for true {
		c_item := <-queue

		lk := <-lock
		//其中一个夺到了，

		switch c_item.cmd {
		case CMD_CGI_RETCODE:
			Write_File(RETCODELOGNAME, c_item.buff)
			break
		case CMD_LOG_IMPORTANT:
			Write_File(IMPORTANTLOGNAME, c_item.buff)
			break
		default:
			break
		}

		lock <- lk
		//写完之后，再往里写一个，让其他的routine去夺
	}
}

func Write_File(pre_filename string, buff []byte) {
	// fmt.Println("consume msg is ", string(buff))

	filename := fmt.Sprintf("%s.%s", pre_filename, time.Now().Format("20060102"))

	var err error
	var f *os.File

	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
		//fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		//fmt.Println("文件不存在")
	}

	defer f.Close()
	if err != nil {
		return
	}
	_, _ = io.WriteString(f, string(buff)+"\n") //写入文件(字符串)
}

func main() {

	addrinfo := fmt.Sprintf(":%d", PORT)
	udp_addr, err := net.ResolveUDPAddr("udp4", addrinfo)
	checkError(err)

	conn, err := net.ListenUDP("udp4", udp_addr)
	defer conn.Close()
	checkError(err)

	queue = make(chan Message_Block, MessageQueue_LENGTH)
	lock = make(chan int, 1)

	for i := 0; i < WORK_ROUTIME_NUM; i++ {
		go Consumer(queue, lock)
	}

	lock <- 2 //往channel中写一个数，让几个routine去夺，其实只是夺写文件的权限而已，消费管道不需要夺

	//go recvUDPMsg(conn)
	recvUDPMsg(conn)

}
