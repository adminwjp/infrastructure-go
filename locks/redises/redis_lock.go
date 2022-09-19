package lock_redises

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
)

type RedisLock struct {
	Client *redis.Client
}

func(lock *RedisLock)Lock(key string)(bool,error){
	cmd:=lock.Client.SetNX(key,nil,time.Second*30)
	if cmd.Err()==nil{
		v,_:=cmd.Result()
		if v{
			return true, nil
		}
		return false, nil
	}
	return false, cmd.Err()
}

func(lock *RedisLock)LockFunc(fun func())(bool,error){
	key:=uuid.New().String()
	v,err:=lock.LockDate(key,time.Second*30)
	if !v{
		return v,err
	}
	defer func() {
		lock.UnLock(key)
	}()
	fun()
	return true, nil
}

func(lock *RedisLock)UnLock(key string)(bool,error){
	cmd:=lock.Client.Del(key)
	if cmd.Err()==nil{
		v,_:=cmd.Result()
		if v>0{
			return true, nil
		}
		return false, nil
	}
	return false, cmd.Err()
}

func(lock *RedisLock)LockDate(key string,time time.Duration)(bool,error){
	cmd:=lock.Client.SetNX(key,nil,time)
	if cmd.Err()==nil{
		v,_:=cmd.Result()
		if v{
			return true, nil
		}
		return false, nil
	}
	return false, cmd.Err()
}