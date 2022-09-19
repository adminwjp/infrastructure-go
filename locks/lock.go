package locks

import "time"

type Lock interface {
	Lock(key string)(bool,error)

	LockFunc(fun func())(bool,error)

	UnLock(key string)(bool,error)

	LockDate(key string,time time.Duration)(bool,error)
}
