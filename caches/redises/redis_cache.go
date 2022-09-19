package redises

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	Addr     string
	Password string
	DB       int
	// 作用域 容易 消失 指针 为 nil
	//client *redis.Client
	Client *redis.Client
}

func RedisCacheInstnce(addr string, password string, db int) *RedisCache {
	cache:= &RedisCache{ }
	cache.Client=cache.GetClient(addr,password,db)
	return cache
}

func NewRedisCache() *RedisCache {
	return RedisCacheInstnce("127.0.0.1:6379", "", 0)
}

func (cache *RedisCache) GetClient(addr string, password string, db int) *redis.Client {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		//PoolSize: 5, //连接池 默认情况下,连接池大小是10
	})
	pingResult, err := redisDb.Ping().Result()
	if err != nil {
		fmt.Println(pingResult, err)
		return nil
	} else {
		fmt.Println("connection redis success")
	}
	//defer  redisDb.Close() //延迟 关闭 操作失败
	return redisDb
}

func (cache *RedisCache) Connection(addr string, password string, db int) bool {
	if cache.Client != nil {
		pingResult, err := cache.Client.Ping().Result()
		if err != nil {
			fmt.Println(pingResult, err)
			//cache.Client.Close()
			cache.Client=cache.GetClient(addr,password,db)
			if cache.Client != nil {
				fmt.Println("reconnection redis success")
				return true
			}
			return false

		} else {
			return true
		}
	}
	redisDb := cache.GetClient(addr,password,db)
	if redisDb == nil {
		return false
	} else {
		cache.Client = redisDb
		return true
	}
}



func (cache *RedisCache) Close() bool {
	if cache.Client != nil {
		err := cache.Client.Close()
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func (cache *RedisCache) Keys( pattern string) ([]string,error){
	stringSliceCmd := cache.Client.Keys(pattern)
	if stringSliceCmd.Err() != nil {
		return nil,stringSliceCmd.Err()
	}
	return stringSliceCmd.Val(),nil
}

/**
   String 操作
　　Set(key, value)：给数据库中名称为key的string赋予值valueget(key)：返回数据库中名称为key的string的value
　　GetSet(key, value)：给名称为key的string赋予上一次的value
　　MGet(key1, key2,…, key N)：返回库中多个string的value
　　SetNX(key, value)：添加string，名称为key，值为value
　　SetXX(key, time, value)：向库中添加string，设定过期时间time
　　MSet(key N, value N)：批量设置多个string的值
　　MSetNX(key N, value N)：如果所有名称为key i的string都不存在
　　Incr(key)：名称为key的string增1操作
　　Incrby(key, integer)：名称为key的string增加integer
　　Decr(key)：名称为key的string减1操作
　　Decrby(key, integer)：名称为key的string减少integer
　　Append(key, value)：名称为key的string的值附加valuesubstr(key, start, end)：返回名称为key的string的value的子串
*/

func (cache *RedisCache)  Set(key string, value interface{}, expiration time.Duration) (bool,error) {
	statusCmd := cache.Client.Set(key, value, expiration)
	return statusCmd.Err()==nil,statusCmd.Err()
}

func (cache *RedisCache) GetSet( key string, value interface{}) (string,error) {
	stringCmd := cache.Client.GetSet(key, value)
	return stringCmd.Val(),stringCmd.Err()
}
func (cache *RedisCache) Get(key string) (string,error) {
	stringCmd := cache.Client.Get(key)
	return stringCmd.Val(),stringCmd.Err()
}

func (cache *RedisCache) MGet( keys ...string) ([]interface{},error) {
	sliceCmd := cache.Client.MGet(keys...)
	return sliceCmd.Val(),sliceCmd.Err()
}

func (cache *RedisCache) SetNX(key string, value interface{}, expiration time.Duration) (bool,error) {
	boolCmd := cache.Client.SetNX(key, value, expiration)
	return boolCmd.Val(),boolCmd.Err()
}

func (cache *RedisCache) SetXX(key string, value interface{}, expiration time.Duration) (bool,error) {
	boolCmd := cache.Client.SetXX(key, value, expiration)
	return boolCmd.Val(),boolCmd.Err()
}

func (cache *RedisCache) MSet( key string, value interface{}, expiration time.Duration) (bool,error) {
	statusCmd := cache.Client.MSet(key, value, expiration)
	return statusCmd.Err()!=nil,statusCmd.Err()
}

func (cache *RedisCache)  MSetNX(pairs []interface{}) (bool,error) {
	boolCmd :=  cache.Client.MSetNX(pairs)
	return boolCmd.Val(),boolCmd.Err()
}

func  (cache *RedisCache) Incr( key string) (int64,error) {
	intCmd := cache.Client.Incr(key)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache) IncrBy(key string, value int64) (int64,error) {
	intCmd := cache.Client.IncrBy(key, value)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache) Decr(key string) (int64,error) {
	intCmd := cache.Client.Decr(key)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache)  DecrBy(key string, value int64) (int64,error) {
	intCmd := cache.Client.DecrBy(key, value)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache)  Append(key string, value string) (int64,error) {
	intCmd := cache.Client.Append(key, value)
	return intCmd.Val(),intCmd.Err()
}

/**
   List 操作
　　RPush(key, value)：在名称为key的list尾添加一个值为value的元素
　　LPush(key, value)：在名称为key的list头添加一个值为value的 元素
　　LLen(key)：返回名称为key的list的长度
　　LRange(key, start, end)：返回名称为key的list中start至end之间的元素
　　LTrim(key, start, end)：截取名称为key的list
　　LIndex(key, index)：返回名称为key的list中index位置的元素
　　LSet(key, index, value)：给名称为key的list中index位置的元素赋值
　　LRem(key, count, value)：删除count个key的list中值为value的元素
　　LPop(key)：返回并删除名称为key的list中的首元素
　　RPop(key)：返回并删除名称为key的list中的尾元素
　　BLPop(key1, key2,… key N, timeout)：lpop命令的block版本。
　　BRPop(key1, key2,… key N, timeout)：rpop的block版本。
　　RPopLPush(srckey, dstkey)：返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
*/

func (cache *RedisCache) RPush( key string, values ...interface{}) (int64,error) {
	intCmd := cache.Client.RPush(key, values...)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache)LPush(key string, values ...interface{}) (int64,error) {
	intCmd := cache.Client.LPush(key, values...)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache) LLen( key string) (int64,error) {
	intCmd := cache.Client.LLen(key)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache) LRange(key string, start int64, end int64) ([]string,error) {
	stringSliceCmd := cache.Client.LRange(key, start, end)
	return stringSliceCmd.Val(),stringSliceCmd.Err()
}

func (cache *RedisCache) LTrim( key string, start int64, end int64) (bool,error) {
	statusCmd := cache.Client.LTrim(key, start, end)
	return statusCmd.Err()==nil,statusCmd.Err()
}

func (cache *RedisCache) LIndex(key string, index int64) (string,error) {
	stringCmd := cache.Client.LIndex(key, index)
	return stringCmd.Val(),stringCmd.Err()
}

func (cache *RedisCache) LSet( key string, index int64, value interface{}) (bool,error) {
	statusCmd := cache.Client.LSet(key, index, value)
	return statusCmd.Err()==nil,statusCmd.Err()
}

func (cache *RedisCache) LRem( key string, index int64, value interface{}) (bool,error) {
	intCmd := cache.Client.LRem(key, index, value)
	return intCmd.Val()>0,intCmd.Err()
}

func (cache *RedisCache) LPop( key string) (string,error) {
	stringCmd := cache.Client.LPop(key)
	return stringCmd.Val(),stringCmd.Err()
}

func (cache *RedisCache) RPop( key string) (string,error) {
	stringCmd := cache.Client.RPop(key)
	return stringCmd.Val(),stringCmd.Err()
}

func (cache *RedisCache) BLPop( timeout time.Duration, keys ...string) ([]string,error) {
	stringSliceCmd := cache.Client.BLPop(timeout, keys...)
	return stringSliceCmd.Val(),stringSliceCmd.Err()
}

func (cache *RedisCache) BRPop(timeout time.Duration, keys ...string) ([]string,error) {
	stringSliceCmd := cache.Client.BRPop(timeout, keys...)
	return stringSliceCmd.Val(),stringSliceCmd.Err()
}

func (cache *RedisCache) RPopLPush( source string, destination string) (bool,error) {
	stringCmd := cache.Client.RPopLPush(source, destination)
	return stringCmd.Err()!=nil,stringCmd.Err()
}


/**
   Hash 操作
　　HSet(key, field, value)：向名称为key的hash中添加元素field
　　HGet(key, field)：返回名称为key的hash中field对应的value
　　HMget(key, (fields))：返回名称为key的hash中field i对应的value
　　HMset(key, (fields))：向名称为key的hash中添加元素field
　　HIncrby(key, field, integer)：将名称为key的hash中field的value增加integer
　　HExists(key, field)：名称为key的hash中是否存在键为field的域
　　HDel(key, field)：删除名称为key的hash中键为field的域
　　HLen(key)：返回名称为key的hash中元素个数
　　HKeys(key)：返回名称为key的hash中所有键
　　HVals(key)：返回名称为key的hash中所有键对应的value
　　HGetall(key)：返回名称为key的hash中所有的键（field）及其对应的value
*/

func (cache *RedisCache) HSet( key string, field string, value interface{}) (bool,error) {
	boolCmd := cache.Client.HSet(key, field, value)
	return boolCmd.Val(),boolCmd.Err()
}

func (cache *RedisCache) HGet(key string, field string) (string,error) {
	stringCmd := cache.Client.HGet(key, field)
	return stringCmd.Val(),stringCmd.Err()
}

func (cache *RedisCache) HMGet(key string, fields ...string) ([]interface{},error) {
	stringCmd := cache.Client.HMGet(key, fields...)
	return stringCmd.Val(),stringCmd.Err()
}

func (cache *RedisCache) HMSet( key string, fields map[string]interface{}) (bool,error) {
	statusCmd := cache.Client.HMSet(key, fields)
	return statusCmd.Err()==nil,statusCmd.Err()
}

func (cache *RedisCache) HIncrBy( key string, field string, incr int64) (int64,error) {
	intCmd := cache.Client.HIncrBy(key, field, incr)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache) HExists(key string, field string) (bool,error) {
	boolCmd := cache.Client.HExists(key, field)
	return boolCmd.Val(),boolCmd.Err()
}

func (cache *RedisCache) HDel( key string, fields ...string) (bool,error) {
	intCmd := cache.Client.HDel(key, fields...)
	return intCmd.Val()>0,intCmd.Err()
}

func (cache *RedisCache) HLen( key string) (int64,error) {
	intCmd := cache.Client.HLen(key)
	return intCmd.Val(),intCmd.Err()
}

func (cache *RedisCache) HKeys( key string) ([]string,error) {
	stringSliceCmd := cache.Client.HKeys(key)
	return stringSliceCmd.Val(),stringSliceCmd.Err()
}

func (cache *RedisCache) HVals( key string) ([]string,error) {
	stringSliceCmd := cache.Client.HVals(key)
	return stringSliceCmd.Val(),stringSliceCmd.Err()
}

func (cache *RedisCache) HGetAll( key string) (map[string]string,error) {
	stringStringMapCmd := cache.Client.HGetAll(key)
	return stringStringMapCmd.Val(),stringStringMapCmd.Err()
}


func (cache *RedisCache)  GetBit(key string ,offset int64)(bool,error)  {
	val:=cache.Client.GetBit(key,offset)
	return val.Val()>0,val.Err()
}

func (cache *RedisCache)  SetBit(key string ,offset int64,val int)(bool,error)  {
	val1:=cache.Client.SetBit(key,offset,val)
	return val1.Val()>0,val1.Err()
}
