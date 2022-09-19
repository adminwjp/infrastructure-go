package mqs

type EmptyMq struct {


}

func (mq *EmptyMq)Publish(key string,val []byte)(bool,error){
	return false, nil
}

func (mq *EmptyMq)Subscript(key string,handVal func([]byte)bool,err error)(bool,error){
	return false, nil
}

func (mq *EmptyMq)PublishString(key string,val string)(bool,error){
	return false, nil
}

func (mq *EmptyMq)SubscriptString(key string,handVal func(string)bool,err error)(bool,error){
	return false, nil
}