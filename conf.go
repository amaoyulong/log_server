package main

//const 变量

const MAX_PACK_LENGTH int = 4096
const PORT int16 = 18600
const VERSION_V1 byte = 0x01

const WORK_ROUTIME_NUM int = 5 //工作协程数

const MessageQueue_LENGTH int = 500 //500条*4096=2M 的空间

const CMD_BIDE2_CGI_RETCODE byte = 0x01
const CMD_BIDE2_IMPORTANT byte = 0x02
const CMD_DCR_CGI_RETCODE byte = 0x03
const CMD_DCR_IMPORTANT byte = 0x04

const BIDE2_RETCODELOGNAME string = "./file/bide2_retcode.log"
const BIDE2_IMPORTANTLOGNAME string = "./file/bide2_important.log"
const DCR_RETCODELOGNAME string = "./file/dcr_retcode.log"
const DCR_IMPORTANTLOGNAME string = "./file/dcr_important.log"

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
