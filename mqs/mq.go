package mqs

type Mq interface {
	Publish(key string,val []byte)(bool,error)
	Subscript(key string,handVal func([]byte)bool,err error)(bool,error)
	PublishString(key string,val string)(bool,error)
	SubscriptString(key string,handVal func(string)bool,err error)(bool,error)
}
