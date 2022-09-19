package tests

import (
	"github.com/adminwjp/infrastructure-go/utils"
	"testing"
)
//go test  -count=1 -v
//go test  -count=1 -v  test
//go test  -count=1 -v  test/file_test.go
func  TestFile(t *testing.T ) {

	f:=utils.FileUtil
	f.GetCurrentDirFileInfosAndDirInfos()
}
