package retries

import "time"

type Retry interface {
	OnRetry(fun func(retryNum int,retryTime time.Duration)bool,sleep time.Duration,retryNum int)bool
}
