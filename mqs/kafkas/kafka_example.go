package kafkas

import "log"

var kafkaMq=&KafkaMq{}

func TestKafkaInit()  {
	kafkaMq.Producer=kafkaMq.CreateProducer(nil)
	kafkaMq.Consumer=kafkaMq.CreateConsumer(nil)

}
func TestKafkaPublish(){
	log.Printf("kafka publis test,msg:test")
	kafkaMq.Publish("test",[]byte("test"))
}
func TestKafkaSubscribe(){
	kafkaMq.Subscript("test", func(bytes []byte) bool {
		log.Printf("kafka subscribe test,msg:%s",string(bytes))
		return true
	},nil)
}