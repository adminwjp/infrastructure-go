package kafkas

import (
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"
)

type  KafkaMq struct {
	Consumer sarama.Consumer
	Producer sarama.SyncProducer
	IsClose bool
	wg sync.WaitGroup
}

func(mq *KafkaMq) CreateProducer(addresses []string)sarama.SyncProducer{
	if addresses==nil{
		addresses=[]string{"127.0.0.1:9092"}
	}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	//producer closed, err: kafka: invalid configuration (Producer.Return.Errors must
	//be true to be used in a SyncProducer)
	//https://blog.csdn.net/yizhiniu_xuyw/article/details/108882988
	//config.Producer.Return.Successes=false
	//config.Producer.Return.Errors=false
	// 连接kafka
	producer, err := sarama.NewSyncProducer(addresses, config)
	if err != nil {
		log.Println("producer closed, err:", err)
		return nil
	}
	return  producer
	//defer client.Close()
}
// 基于sarama第三方库开发的kafka client

func(mq *KafkaMq) PublishMsg(producer sarama.SyncProducer,topic string,message string) (bool,error) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic

	msg.Value = sarama.StringEncoder(message)
	// 发送消息
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("send %s msg failed, err: %s\n",topic, err.Error())
		return false,err
	}
	log.Printf("send %s msg pid:%d offset:%d\n", topic,pid, offset)
	return  true,err
}
func (mq *KafkaMq)CreateConsumer(addresses []string)sarama.Consumer{
	if addresses==nil{
		addresses=[]string{"127.0.0.1:9092"}
	}
	config := sarama.NewConfig()
	//kafka_2.11-2.4.0-site-docs
	config.Version=sarama.V2_4_0_0
	/*	config.Consumer.MaxWaitTime = time.Second*2
	config.Consumer.Offsets.AutoCommit.Enable =true
	config.Consumer.Offsets.AutoCommit.Interval =time.Millisecond*10
	config.Consumer.Retry.Backoff=time.Second*2
	config.Consumer.Return.Errors=true*/
	// 成功交付的消息将在success channel返回
	consumer, err := sarama.NewConsumer(addresses, config)
	if err != nil {
		log.Printf("create consumer   fail , err:%s\n",err.Error())
		return nil
	}
	return consumer
}
func (mq *KafkaMq)ConsumerMsg(consumer sarama.Consumer,topic string,handler func(msg string))(bool,error) {
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		log.Printf("consumer %s msg fail to get list of partition:err%s\n",
			topic, err.Error())
		return false,err
	}
	//log.Println(partitionList)
	log.Printf("consumer %s msg ", topic)
	//consumer.ConsumePartition()
	//empty o
	res:=consumer.HighWaterMarks()
	for n, v := range res {
		if n==topic{
			for p, o := range v {
				pc, err := consumer.ConsumePartition(n, p, o)
				if err != nil {
					log.Printf("consumer %s msg failed to start consumer for partition %d,err:%s\n",
						topic,  p, err.Error())
					return false,err
				}
				defer pc.AsyncClose()
				mq.wg.Add(1)
				go func(  sarama.PartitionConsumer) {
					//一直阻塞
					mq.wg.Done()
					for msg := range pc.Messages() {
						log.Printf("consumer %s msg suc",msg.Topic)
						handler(string(msg.Value))
						log.Printf("consumer %s msg Partition:%d Offset:%d Key:%v Value:%v\n",
							msg.Topic,msg.Partition, msg.Offset, msg.Key, msg.Value)
					}
					for msg := range pc.Errors() {
						log.Printf("consumer %s msg err",msg.Topic)
						handler(string(msg.Error()))
					}
				}(pc)
				mq.wg.Wait()
			}
		}

	}
	//return  true,err
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Printf("consumer %s msg failed to start consumer for partition %d,err:%s\n",
				topic,  partition, err.Error())
			return false,err
		}
		if mq.IsClose{

		}

		// 异步从每个分区消费信息
		//没日志执行
		//pc1:=pc
		//defer pc.AsyncClose()
		mq.wg.Add(1)
		go func(  sarama.PartitionConsumer) {
			mq.wg.Done()
			log.Printf("consumer  kafka  %s  msg", topic)
			log.Printf("consumer kafka %s  msg ....,%d", topic,pc.HighWaterMarkOffset())
			//一直阻塞

		/*	for msg := range pc.Errors() {
				log.Printf("consumer %s msg err",msg.Topic)
				handler(string(msg.Error()))
			}
			log.Println("err end")*/
			//https://github.com/Shopify/sarama/blob/main/examples/consumergroup/main.go
			//pass
			for  {

				select {
				case msg:=<-pc.Messages():


				//for msg := range pc.Messages() {
					log.Printf("consumer %s msg suc",msg.Topic)
					handler(string(msg.Value))
					log.Printf("consumer %s msg Partition:%d Offset:%d Key:%v Value:%v\n",
						msg.Topic,msg.Partition, msg.Offset, msg.Key, msg.Value)
				case msg:=<-pc.Errors():
					log.Printf("consumer %s msg err",msg.Topic)
					handler(string(msg.Error()))
				default:
					time.Sleep(time.Millisecond*20)
				}
				time.Sleep(time.Millisecond*10)
			}
			log.Println("msg end")

		}(pc)
		mq.wg.Wait()
	}
	return  true,err
}

func (mq *KafkaMq)Publish(key string,val []byte)(bool,error){
	return mq.PublishMsg(mq.Producer,key,string(val))
}

func (mq *KafkaMq)Subscript(key string,handVal func([]byte)bool,err error)(bool,error){
	return mq.ConsumerMsg(mq.Consumer,key, func(msg string) {
		handVal([]byte(msg))
	})
}

func (mq *KafkaMq)PublishString(key string,val string)(bool,error){
	return mq.PublishMsg(mq.Producer,key,val)
}

func (mq *KafkaMq)SubscriptString(key string,handVal func(string)bool,err error)(bool,error){
	return mq.ConsumerMsg(mq.Consumer,key, func(msg string) {
		handVal(msg)
	})
}
