package lock_zookeepers

import (
	"github.com/adminwjp/infrastructure-go/retries"
	"github.com/google/uuid"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

type ZookeeperLock struct {
	Client *zk.Conn
	Retry retries.Retry
	LockDates map[string]time.Time
}
func (lock *ZookeeperLock)StartCleanTimeoutThread(){
	go func() {
		for  {
			if lock.Client==nil{
				time.Sleep(time.Millisecond*100)
				continue
			}
			lock.CleanTimeoutPath()
			time.Sleep(time.Millisecond*20)
		}
	}()
}
func (lock *ZookeeperLock)CleanTimeoutPath()  {
	for k, v := range lock.LockDates {
		if v.Unix()<time.Now().Unix()-5*1000*1000{
			suc,_:=lock.Delete(k,1)
			if suc{
				delete(lock.LockDates,k)
			}
		}
	}
}
func (lock *ZookeeperLock)Conn(connectionStrings []string,val string){
	if connectionStrings==nil{
		connectionStrings= []string{"localhost:2181"}
	}
	//zk.Connect(connectionStrings, 10*time.Second)
	conn,_,err:= zk.Connect(connectionStrings, 10*time.Second)
	if err != nil {
		log.Println(err)
		return
	}
	lock.Client=conn
}
func (lock *ZookeeperLock)Set(key string,val []byte ,version int32)(bool,error){
	state,err:=lock.Client.Set(key,val,version)
	if err!=nil{
		return false,err
	}
	if state.Version==version{

	}
	return true, err
}
func (lock *ZookeeperLock)Delete(key string,version int32)(bool,error){
	err:=lock.Client.Delete(key,version)
	if err!=nil{
		return false,err
	}
	return true, err
}
func (lock *ZookeeperLock)Get(key string)(string,error){
	v,state,err:=lock.Client.Get(key)
	if err!=nil{
		return "",err
	}
	if state.Version>0{

	}
	return string(v),nil
}

func(lock *ZookeeperLock)Lock(key string)(bool,error){
	return lock.Set(key,nil,1)
}

func(lock *ZookeeperLock)LockFunc(fun func())(bool,error){
	path:="zk/"+uuid.New().String()
	suc,err:= lock.Set(path,nil,1)
	if err!=nil{
		return suc,err
	}
	return lock.Delete(path,1)
}

func(lock *ZookeeperLock)UnLock(key string)(bool,error){
	return lock.Delete(key,1)
}

func(lock *ZookeeperLock)LockDate(key string,time1 time.Duration)(bool,error){
	lock.LockDates[key]=time.Now().Add(time1)
	return lock.Set(key,nil,1)
}