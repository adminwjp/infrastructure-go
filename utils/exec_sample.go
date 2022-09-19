package utils

import "log"

func  TestCmd() {

	ex:=ExecUtil
	pId:=ex.FindPortToPId(1050)//sc list netstat -a

	log.Println("test  find port pid ,%d,==>",pId)

	ex.GetCmdResult("netstat -a")//sc list netstat -a
}
