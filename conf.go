package main

//const 变量

const MAX_PACK_LENGTH int = 4096
const PORT int16 = 18600
const VERSION_V1 byte = 0x01

const WORK_ROUTIME_NUM int = 5 //工作协程数

const MessageQueue_LENGTH int = 500 //500条*4096=2M 的空间

const CMD_CGI_RETCODE byte = 0x01
const CMD_LOG_IMPORTANT byte = 0x02

const RETCODELOGNAME string = "./file/retcode.log"
const IMPORTANTLOGNAME string = "./file/important.log"

//消息定义
type Record_Req struct {
	ver  byte
	cmd  byte
	leng [4]byte
}

type Message_Block struct {
	cmd  byte
	buff []byte
}
