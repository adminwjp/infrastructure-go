package utils

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)
var ExecUtil execUtil=execUtil{}
var LinuxShell="/bin/ls"
var WindowCmd="C:\\Windows\\System32\\cmd.exe"
//https://blog.csdn.net/weixin_44282540/article/details/109066976
type execUtil struct {

}
func (execUtil) FindPortToPId(port int)int {
	arch:=runtime.GOARCH
	os:=runtime.GOOS
	log.Printf("arch:%s,os:%s",arch,os)
	commandName:="cmd"
	commandPath:="/c"
	cmdStr:=""
	switch os {
		case "windows":
			cmdStr=fmt.Sprintf("netstat -ano -p tcp | findstr %d",port)
			break
		case "linux":
			commandName="bash"
			commandPath="-c"
			cmdStr=fmt.Sprintf("lsof  -i:%d",port)
			break
	}
	cmd := exec.Command(commandName,commandPath, cmdStr)
	output,err := cmd.Output()
	if err != nil {
		log.Println("error:"+err.Error())
		return -1
	}
	res:=string(output)
	log.Println(res)
	r:=regexp.MustCompile(`\s\d+\s`).FindAllString(res,-1)
	if len(r)>0{
		pId,err:=strconv.Atoi(strings.TrimSpace(r[0]))
		if err!=nil{
			return -1
		}
		return pId
	}
	return -1

}
func (execUtil) GetCmdResult(cmdOrShellStr string )(string,error){
	arch:=runtime.GOARCH
	os:=runtime.GOOS
	log.Printf("arch:%s,os:%s",arch,os)
	commandName:="cmd"
	commandPath:="/c"
	switch os {
	case "windows":break
	case "linux":
		commandName="bash"
		commandPath="-c"
		break
	}
	cmd := exec.Command(commandName,commandPath, cmdOrShellStr)
	// 执行命令，并返回结果
	output,err := cmd.Output()
	//err = cmd.Run() //执行命令，返回命令是否执行成功
	if err != nil {
		log.Println("error:"+err.Error())
		return "", err
	}
	log.Println(string(output))
	return string(output),nil

}
func (execUtil) GetCmd(params ...string )(string,error) {
	arch:=runtime.GOARCH
	os:=runtime.GOOS
	log.Printf("arch:%s,os:%s",arch,os)
	commandName:="cmd"
	switch os {
		case "windows":break
		case "linux":
			commandName="bash"
			break
	}
	// 通过exec.Command函数执行命令或者shell
	// 第一个参数是命令路径，当然如果PATH路径可以搜索到命令，可以不用输入完整的路径
	// 第二到第N个参数是命令的参数
	// 下面语句等价于执行命令: ls -l /var/
	//cmd := exec.Command("/bin/ls", "-l", "/var/")
	cmd := exec.Command(commandName, params...)
	// 执行命令，并返回结果
	output,err := cmd.Output()
	//err = cmd.Run() //执行命令，返回命令是否执行成功
	if err != nil {
		log.Println("error:"+err.Error())
		return "", err
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))
	return string(output),nil
}