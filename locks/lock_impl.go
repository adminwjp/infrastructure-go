package locks

import (
	"sync"
	"time"
)
var LockInstance=&LockImpl{}
type LockImpl struct {
	Mutex *sync.RWMutex
	Locks map[string]sync.RWMutex
	LockDates map[string]time.Time
}

func NewLock()*LockImpl {
	return &LockImpl{}
}
func (lock *LockImpl)StartCleanTimeoutThread(){
	go func() {
		for  {
			lock.CleanTimeoutPath()
			time.Sleep(time.Millisecond*20)
		}
	}()
}
func (lock *LockImpl)CleanTimeoutPath()  {
	for k, v := range lock.LockDates {
		if v.Unix()<time.Now().Unix()-5*1000*1000{
			delete(lock.LockDates,k)
		}
	}
}
func(lock *LockImpl)Lock(key string)(bool,error){
	lock.Mutex.Lock()
	m,e:=lock.Locks[key]
	if !e{
		m:=sync.RWMutex{}
		lock.Locks[key]=m
	}
	m.Lock()
	return true,nil
}

func(lock *LockImpl)LockFunc(fun func())(bool,error){
	m:=sync.RWMutex{}
	m.Lock()
	fun()
	m.Unlock()
	return true, nil
}

func(lock *LockImpl)UnLock(key string)(bool,error){
	lock.Mutex.Unlock()
	m,e:=lock.Locks[key]
	if e{
		m.Unlock()
		return true, nil
	}
	return false, nil
}

func(lock *LockImpl)LockDate(key string,time1 time.Duration)(bool,error){
	lock.LockDates[key]=time.Now().Add(time1)
	return lock.Lock(key)
}