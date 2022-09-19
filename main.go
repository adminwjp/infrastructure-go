package main

import (
	"github.com/adminwjp/infrastructure-go/mqs/kafkas"
	"github.com/adminwjp/infrastructure-go/utils"
	"log"
	"time"
)

func main() {
	//
	currDir:=utils.FileUtil.GetCurrentDir()
	log.Println(currDir)
	kafkas.TestKafkaInit()

	kafkas.TestKafkaSubscribe()
	time.Sleep(time.Second*5)


	kafkas.TestKafkaPublish()
	kafkas.TestKafkaPublish()





	for  {
		time.Sleep(time.Millisecond*100)
	}
}
