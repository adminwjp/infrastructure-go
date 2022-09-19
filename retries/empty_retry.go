package retries

import "time"

type EmptyRetry struct {
	
	
}
func (retry *EmptyRetry)OnRetry(fun func(retryNum int,retryTime time.Duration)bool,sleep time.Duration,retryNum int)bool{
	return true
}