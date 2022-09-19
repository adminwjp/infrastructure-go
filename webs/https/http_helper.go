package https

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var HttpServerInstance =&HttpServerHelper{}

type HttpServerHelper struct {

}
func (*HttpServerHelper) Test(writer http.ResponseWriter,request *http.Request){
	io.WriteString(writer,"test")
}
func (*HttpServerHelper) HTTPHandlerFuncInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// TODO: 进行身份验证，比如校验cookie或token
			h(w, r)
		})
}

func (*HttpServerHelper) GetPageAndSize(request *http.Request) (int,int) {
	url :=request.URL.Path
	reg,err:=regexp.Compile(`/(\d+/{0,1})/(\d+/{0,1})[/|?]{0,1}`)
	if err!=nil{
		log.Println(" find reg pattern faill")
		return 0,0
	}else if !reg.MatchString(url){
		log.Printf( "%s find reg match faill",url)
		return 0,0
	}else{
		strs:=reg.FindAllString(url,-1)
		if len(strs)>=2{
			println(strs[0]+"="+strs[1])
		}else{
			log.Printf( "%s find reg match page or  size faill ,len %d %s",url,len(strs),strs[0])
		}
		page,err :=strconv.Atoi(strings.Replace(strs[0],"/","",-1))

		if err!=nil{
			log.Printf( "%s find reg match page faill",url)
			return 0,0
		}
		size,err :=strconv.Atoi(strings.Replace(strs[1],"/","",-1))
		if err!=nil{
			log.Printf( "%s find reg match size faill",url)
			return page,0
		}
		return page,size
	}
	return 0,0
}

type  handlerInterceptor struct {
	Handlers []http.Handler
}
func  (handlerInterceptor handlerInterceptor) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// TODO: 进行身份验证，比如校验cookie或token
	//w.Write([]byte("not login"))
	//return
	for i:=0;i<len(handlerInterceptor.Handlers);i++ {
		handlerInterceptor.Handlers[i].ServeHTTP(w,r)
	}
}
var Handler handlerInterceptor=handlerInterceptor{}
func HTTPInterceptor(h http.Handler) http.Handler {
	return Handler
}
func (server *HttpServerHelper) Start( router func(),port string,ports []string,statics []string)  {

	http.HandleFunc("/test",server.HTTPHandlerFuncInterceptor(server.Test))
	//http.Handle("/", http.FileServer(http.Dir("static")))
	Handler.Handlers=append(Handler.Handlers,http.FileServer(http.Dir("static")))
	if statics!=nil && len(statics)>0{
		for i := 0; i < len(statics); i++ {
			Handler.Handlers=append(Handler.Handlers,http.FileServer(http.Dir(statics[i])))
		}
	}
	http.Handle("/", Handler)
	if port!="" {
		server:=&http.Server{
			Addr: port,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		//server.Close()
		//err:=http.ListenAndServe("",nil)
		log.Println("http server starting")
		err:=server.ListenAndServe()
		if err!=nil{
			log.Fatal("http server start fail")
			panic(err)
		}
		log.Println("http server start success")
	}else{
		//浏览器访问 http://localhost:8080/api
		for i := 0; i < len(ports); i++ {
			mux := http.NewServeMux()
			go http.ListenAndServe(ports[i], mux)
		}
		log.Println("http server start success")
		//阻塞程序
		select {}
	}
}
