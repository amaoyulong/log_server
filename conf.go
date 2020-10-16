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
const CMD_LBB_CGI_RETCODE byte = 0x05
const CMD_LBB_IMPORTANT byte = 0x06
const CMD_DCC_IMPORTANT byte = 0x07
const CMD_DCC_CGI_RETCODE byte = 0x08

const CMD_END byte = 0x09

var CMD2File map[byte]string

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

func Init() {
	CMD2File = make(map[byte]string)

	CMD2File[CMD_BIDE2_CGI_RETCODE] = "./file/bide2_retcode.log"
	CMD2File[CMD_BIDE2_IMPORTANT] = "./file/bide2_important.log"

	CMD2File[CMD_DCR_CGI_RETCODE] = "./file/dcr_retcode.log"
	CMD2File[CMD_DCR_IMPORTANT] = "./file/dcr_important.log"

	CMD2File[CMD_LBB_CGI_RETCODE] = "./file/lbb_retcode.log"
	CMD2File[CMD_LBB_IMPORTANT] = "./file/lbb_important.log"

	CMD2File[CMD_DCC_IMPORTANT] = "./file/dcc_retcode.log"
	CMD2File[CMD_DCC_CGI_RETCODE] = "./file/dcc_important.log"
}
