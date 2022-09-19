package utils

//package utility
//go run: cannot run non-main package

import (
	"bufio"
	"fmt"
	//"io"
	"log"
	"net"
	//"net/http"
	"time"
)

var (
	serverConnect    = make(chan chan<- string)
	serverDisconnect = make(chan chan<- string)
	globalMsg        = make(chan string)
)

type SocketUtil struct {
	Server net.Listener
}
func (s *SocketUtil)ServerStart(serverAddr string) {
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("服务器启动成功,", serverAddr)
	s.Server=listener
}
func (s *SocketUtil)ServerHandler(brodcaseMsg string, handlerMsg func(msg string) string) {

	go Broadcase()
	for {
		conn, err := s.Server.Accept()
		if err != nil {
			log.Println("server listener client fail,error:" + err.Error())
			continue
		}
		who := conn.RemoteAddr().String()
		go s.ServerHandleConn(conn, brodcaseMsg+who , handlerMsg)
	}
	return
}
func (s *SocketUtil)ServerHandleConn(conn net.Conn, brodcaseMsg string, reciverMsg func(msg string) string) {
	// 一个通道
	ch := make(chan string)
	// 从通道里面读取信息,读到后就发送给客户端
	go SendMsg(conn, ch)
	who := conn.RemoteAddr().String()
	// 给客户端发送信息
	//ch <- msg

	// 给广播
	globalMsg <- brodcaseMsg
	log.Println("brodcase msg,who:" + who + ",msg:" + brodcaseMsg)
	serverConnect <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		//globalMsg<-"get msg,who:"+who+",msg:"+input.Text()
		log.Println("server get  client msg,who:" + who + ",msg:" + input.Text())
		msg := reciverMsg(input.Text())
		// 给客户端发送信息
		ch <- msg
		//不给 cpu 爆炸
		time.Sleep(1)
	}
	// 有通道断开连接了
	serverConnect <- ch
	//globalMsg<-"client left,who:"+who
	log.Println("client left,who:" + who)
	conn.Close()
}
func Broadcase() {
	clients := make(map[chan<- string]bool)
	for {
		select {
		case client := <-serverConnect:
			clients[client] = true
		case msg := <-globalMsg:
			for client := range clients {
				client <- msg
			}
		case clent := <-serverDisconnect:
			delete(clients, clent)
			close(clent)
		}
	}
}
func TestHandleConn(conn net.Conn, reciverMsg func(msg string) string) {
	who := conn.RemoteAddr().String()
	HandleConn(conn, "You are "+who, "brodcase msg,who:"+who, reciverMsg)

}

func HandleConn(conn net.Conn, msg string, brodcaseMsg string, reciverMsg func(msg string) string) {
	// 一个通道
	ch := make(chan string)
	// 从通道里面读取信息,读到后就发送给客户端
	go SendMsg(conn, ch)
	who := conn.RemoteAddr().String()
	// 给客户端发送信息
	ch <- msg
	log.Println("server send  to client msg,who:" + who + ",msg:" + msg)

	// 给广播
	globalMsg <- brodcaseMsg
	log.Println("brodcase msg,who:" + who + ",msg:" + brodcaseMsg)
	serverConnect <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		//globalMsg<-"get msg,who:"+who+",msg:"+input.Text()
		log.Println("server get  client msg,who:" + who + ",msg:" + input.Text())
		msg := reciverMsg(input.Text())
		// 给客户端发送信息
		ch <- msg
		//不给 cpu 爆炸
		time.Sleep(1)
	}
	// 有通道断开连接了
	serverConnect <- ch
	//globalMsg<-"client left,who:"+who
	log.Println("client left,who:" + who)
	conn.Close()
}

//发送消息给客户端
func SendMsg(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func TestServerStart() {
	ServerStart("127.0.0.1:5010","You are ", "brodcase msg,who:", func(msg string) string {
		if msg == "login" {
			return "1"
		}
		return "test"
	})
}

//https://blog.csdn.net/themagickeyjianan/article/details/106953939
//127.0.0.1:8000
func ServerStart1(serverAddr string) (net.Listener,error){
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("服务器启动成功,", serverAddr)
	return listener, err
}
func ServerStart(serverAddr string, msg string, brodcaseMsg string, handlerMsg func(msg string) string) {
	listener, _ := ServerStart1( serverAddr)
	go Broadcase()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("server listener client fail,error:" + err.Error())
			continue
		}
		who := conn.RemoteAddr().String()
		go HandleConn(conn, msg +who, brodcaseMsg+who , handlerMsg)
	}
	return
}
