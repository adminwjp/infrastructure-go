package tests

import (
	"testing"
	"github.com/adminwjp/infrastructure-go/utils"
)
//go test  -count=1 -v
//go test  -count=1 -v  test/cmd_test.go
func  TestCmd(t *testing.T ) {

	ex:=utils.ExecUtil
	pId:=ex.FindPortToPId(1050)//sc list netstat -a

	t.Logf("test  find port pid ,%d,==>",pId)

	ex.GetCmdResult("netstat -a")//sc list netstat -a
}
