package consul_locks

import (
	"github.com/hashicorp/consul/api"
	"sync"
	"time"
)

type ConsulLock struct {
	Client               *api.Client
	LockDates map[string]time.Time
}



func (lock *ConsulLock)StartCleanTimeoutThread(){
	go func() {
		for  {
			lock.CleanTimeoutPath()
			time.Sleep(time.Millisecond*20)
		}
	}()
}
func (lock *ConsulLock)CleanTimeoutPath()  {
	for k, v := range lock.LockDates {
		if v.Unix()<time.Now().Unix()-5*1000*1000{
			delete(lock.LockDates,k)
		}
	}
}
func(lock *ConsulLock)Lock(key string)(bool,error){
	l,err:=lock.Client.LockKey(key)
	if err!=nil{
		return false, err
	}
	if l==nil{
		return false, nil
	}
	//ch := make(chan struct{})
	//ch<-struct {}{}
	res,err:= l.Lock(nil)
	return res==nil,err
}

func(lock *ConsulLock)LockFunc(fun func())(bool,error){
	m:=sync.RWMutex{}
	m.Lock()
	fun()
	m.Unlock()
	return true, nil
}

func(lock *ConsulLock)UnLock(key string)(bool,error){
	l,err:=lock.Client.LockKey(key)
	if err!=nil{
		return false, err
	}
	if l==nil{
		return false, nil
	}
	err= l.Unlock()
	return err==nil,err
}

func(lock *ConsulLock)LockDate(key string,time1 time.Duration)(bool,error){
	lock.LockDates[key]=time.Now().Add(time1)
	return lock.Lock(key)
}