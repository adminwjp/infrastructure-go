package rabbitmqs

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}



//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.Conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Println("%s:%s", message, err)
		//panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(url string) *RabbitMQ {
	if url==""{
		url= MQURL
	}
	//创建RabbitMQ实例
	rabbitmq := &RabbitMQ{}
	var err error
	//获取connection
	rabbitmq.Conn, err = amqp.Dial(url)
	rabbitmq.failOnErr(err, "failed to connect rabb"+
		"itmq!")
	//获取channel
	rabbitmq.Channel, err = rabbitmq.Conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(queueName string,exchange string,message string) (bool,error){
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.Channel.QueueDeclare(
		queueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//调用channel 发送消息到队列中
	err=r.Channel.Publish(
		exchange,
		queueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err!=nil{
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(queueName string,handler func([]byte))(bool,error) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.Channel.QueueDeclare(
		queueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//接收消息
	msgs, err := r.Channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err!=nil{
		fmt.Println(err)
		return false, err
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			handler(d.Body)
			log.Printf("Received a message: %s", d.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return true, nil
}
func (mq *RabbitMQ)Publish(key string,val []byte)(bool,error){
	return mq.PublishSimple(key,"",string(val))
}

func (mq *RabbitMQ)Subscript(key string,handVal func([]byte)bool,err error)(bool,error){
	return mq.ConsumeSimple(key, func(msg []byte) {
		handVal(msg)
	})
}

func (mq *RabbitMQ)PublishString(key string,val string)(bool,error){
	return mq.PublishSimple(key,"",val)
}

func (mq *RabbitMQ)SubscriptString(key string,handVal func(string)bool,err error)(bool,error){
	return mq.ConsumeSimple(key, func(msg []byte) {
		handVal(string(msg))
	})
}