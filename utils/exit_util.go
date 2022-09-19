package utils

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ListenApplicationExit(exit func())  {
	var stopLock *sync.Mutex=&sync.Mutex{}
	stop:=false
	stopChan:=make(chan struct{},1)
	signalChan:=make(chan os.Signal,1)
	go func() {
		<-signalChan
		stopLock.Lock()
		exit()
		stop=true
		stopLock.Unlock()
		log.Println("listen application exit ....")
		stopChan<- struct{}{}
		os.Exit(0)
	}()
	signal.Notify(signalChan,syscall.SIGINT,syscall.SIGTERM)
	time.Sleep(10*time.Second)
}
