package thrifts

import (
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"time"
)
var ThiriftInstance=ThiriftHelper{}
type ThiriftHelper struct {
	ClientTransport thrift.TTransport
	ClientProtocol thrift.TProtocolFactory

	ClientTransportSalves []thrift.TTransport
	ClientProtocolSalves []thrift.TProtocolFactory

	ServerTransportSalves []thrift.TTransport
	ServerProtocolSalves []thrift.TProtocolFactory

	Server *thrift.TSimpleServer
	Client *thrift.TSocket
}
func (rpc *ThiriftHelper)StartThriftClient(listenAdd string)(thrift.TTransport,
	thrift.TProtocolFactory,error){
	read:=true
	write:=true
	conf:=&thrift.TConfiguration{
		TBinaryStrictWrite: &write,
		TBinaryStrictRead: &read,
	}
	socket1:=thrift.NewTSocketConf(listenAdd,conf)
	socket1.Open()
	rpc.Client=socket1
	protocolFactory:=thrift.NewTBinaryProtocolFactoryConf(conf)
	transportFactory:=thrift.NewTTransportFactory()
	transportFactory=thrift.NewTFramedTransportFactoryConf(transportFactory, conf)
	/*transportFactory:=thrift.NewTMemoryBufferTransportFactory(1024)
	 */
	transport,err:=transportFactory.GetTransport(socket1)
	if err!=nil{
		log.Println("thrift client transport fail,err:%s")
		return transport,protocolFactory,err
	}
	return transport,protocolFactory,nil
}
func(rpc *ThiriftHelper) StartThriftServer(listenAdd string,multiplexedProcessor1 func(*thrift.TMultiplexedProcessor))  {
	serverSocket,err:=thrift.NewTServerSocketTimeout(listenAdd,time.Second*10)
	if err!=nil{
		log.Println("thrift server serverSocket fail,err:%s",err.Error())
		return
	}
	log.Println("thrift server serverSocket suc,:%s",listenAdd)
	read:=true
	write:=true
	conf:=&thrift.TConfiguration{
		TBinaryStrictWrite: &write,
		TBinaryStrictRead: &read,
		MaxFrameSize: 16384000,
		MaxMessageSize: 104857600,
		ConnectTimeout:time.Second*10,
		SocketTimeout:time.Second*5 ,
	}
	protocolFactory:=thrift.NewTBinaryProtocolFactoryConf(conf)
	multiplexedProcessor:=thrift.NewTMultiplexedProcessor()
	//multiplexedProcessor.RegisterDefault(multiplexedProcessor)
	multiplexedProcessor1(multiplexedProcessor)
	//multiplexedProcessor.DefaultProcessor=multiplexedProcessor;
	transportFactory:=thrift.NewTTransportFactory()
	//transportFactory:=thrift.NewTMemoryBufferTransportFactory(1024*1024*10)
	//transportFactory:=thrift.NewTBufferedTransportFactory(1024*1024*10)
	//thrift.NewTMultiplexedProtocol(nil,"default")
	factory:=thrift.NewTFramedTransportFactoryConf(transportFactory, conf)
	/*	thriftServer:=thrift.NewTSimpleServerFactory6(thrift.NewTProcessorFactory(multiplexedProcessor),
		serverSocket,transportFactory,transportFactory,protocolFactory,protocolFactory)
	*/
	thriftServer:=thrift.NewTSimpleServerFactory4(thrift.NewTProcessorFactory(multiplexedProcessor),
		serverSocket,factory,protocolFactory)
	//thriftServer:=thrift.NewTSimpleServerFactory2(thrift.NewTProcessorFactory(multiplexedProcessor),
	//	serverSocket)
	log.Println("thrift server  listening ,Serve ..... ")
	rpc.Server=thriftServer
	thriftServer.SetLogger(func(msg string) {
		log.Println("thrift server log:>"+msg)
	})

	err=thriftServer.Serve()
	if err!=nil{
		log.Println("thrift server  fail,Serve err:%s",err.Error())
		return
	}

}