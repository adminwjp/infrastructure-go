package locks

import "time"

var EmptyLockInstance=&EmptyLock{}
type EmptyLock struct {

}
func(lock *EmptyLock)Lock(key string)(bool,error){
	return true, nil
}

func(lock *EmptyLock)LockFunc(fun func())(bool,error){
	return true, nil
}

func(lock *EmptyLock)UnLock(key string)(bool,error){
	return true, nil
}

func(lock *EmptyLock)LockDate(key string,time time.Duration)(bool,error){
	return true, nil
}