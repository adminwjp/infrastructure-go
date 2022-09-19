package grpcs

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)
var GrpcInstance=GrpcHelper{}
type GrpcHelper struct {
	Client *grpc.ClientConn
	ClientSalves []*grpc.ClientConn
	ClientUsedes []bool

	Server *grpc.Server
	ServerSalves []*grpc.ClientConn


	ClientStart bool
	ServerStart bool
	ClientStartDate int64
	ClientIdeaDate int64
	IdeaDate int64
	Used bool
	Tatget string
}
func (rpc *GrpcHelper)GetRpcClient()*grpc.ClientConn{
	return rpc.Client
}
func (rpc *GrpcHelper)RpcClientCheck()  {
	go func() {
		switch rpc.Client.GetState() {
		case connectivity.Connecting:
			break
		case connectivity.Idle:
			rpc.ClientIdeaDate=time.Now().Unix()
			break
		case connectivity.Ready:
			rpc.ClientStartDate=time.Now().Unix()
			break
		case connectivity.TransientFailure:
			break
		case connectivity.Shutdown:
			if rpc.Used{
				t:=rpc.Client.Target()
				rpc.Client.Close()
				rpc.Client,_=rpc.StartGrpcClient(t)

			}
			break
		default:
			break

		}
		if !rpc.Used&&time.Now().Unix()-rpc.ClientIdeaDate>rpc.IdeaDate-300{
			rpc.Client.Close()
		}
		time.Sleep(100)
	}()
}
func (rpc *GrpcHelper) StartGrpcClient(port string)(*grpc.ClientConn,error){
	//grpc.Dial("")
	//conn:=&grpc.ClientConn{}
	rpc.Tatget=port
	conn,err:=grpc.Dial(port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	return conn,err
}
func (rpc *GrpcHelper) StartGrpcServer(port string,register func(server *grpc.Server))  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("grpc server failed to listen: %s", err)
		return
	}
	log.Println("grpc server suc to listen: %s", port)
	s := grpc.NewServer()
	register(s)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("grpc server  listening ,Serve ..... ")
	if err := s.Serve(lis); err != nil {
		log.Println("grpc server failed to serve: %s", err)
		return
	}
	log.Println("grpc server suc to serve")
}
