package gins

import (
	"encoding/json"
	 dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"github.com/adminwjp/infrastructure-go/webs"
)

var GinServerInstance=&WebGinServer{}

type WebGinServer struct {

}
// @title gin swageer
// @version 1.0
// @description gin swageer  description
// @termsOfService [http://swagger.io/terms/](http://swagger.io/terms/)`
// @contact.name swagger
// @contact.url [http://www.swagger.io/support](http://www.swagger.io/support)`
// @contact.email support@swagger.io`
// @license.name Apache 2.0`
// @license.url [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`
// @host  http://192.168.1.3:4900
// @BasePath ""
func  init()  {

}
func (server *WebGinServer) Register(r *gin.Engine)  {
	//r.Use(server.RegisterCors())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": false,"code":404,"msg":"resource not found "})
		c.Abort()
	})

	r.NoMethod(func(c *gin.Context) {
		defer func() {
			if r:=recover();r!=nil{
				c.JSON(http.StatusOK, gin.H{"status": false,"code":500,"msg":"server inner error "})
			}

		}()
		c.Next()
	})
	r.GET("/test", func(c *gin.Context){
		c.JSON(200,gin.H{
			"status": true, "code": 200, "msg": "success",
		})
	})
}
func (server *WebGinServer)RegisterCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		defer func() {
			if err:=recover();err!=nil{
				log.Println(string(debug.Stack()))
				log.Println(err)
				c.JSON(http.StatusOK, gin.H{"status": false,"code":500,"msg":"system error"})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
type GinFilter struct {
	SourceType webs.SourceType
	ReqParams map[string]string
	TOKEN_NAME string
	Sign string
	AllowAnonymouUrls []string
}

func NewGinFilter() *GinFilter {
	return &GinFilter{TOKEN_NAME: "token",Sign:"sign"}
}

func (filter *GinFilter) OnActionExecution(c *gin.Context) bool{
	userAgent:=c.Request.UserAgent()
	if strings.Contains(userAgent,"MicroMessenger"){
		filter.SourceType= webs.WeChatApp
	}else if strings.Contains(userAgent,"iPhone") || strings.Contains(userAgent,"iPod")||strings.Contains(userAgent,"iPad"){
		filter.SourceType= webs.IOS
	}else if strings.Contains(userAgent,"Android"){
		filter.SourceType= webs.Android
	}else{
		filter.SourceType= webs.Web
	}
	return false
	vals:=c.Request.URL.Query()
	for k,v:=range vals{
		filter.ReqParams[k]=v[0]
	}
	if c.Request.Form!=nil{
		for k,v:=range c.Request.Form{
			filter.ReqParams[k]=v[0]
		}
	}
	reader,err:=c.Request.GetBody()
	if err==nil{
		buffer,err:=ioutil.ReadAll(reader)
		if err==nil{
			var post	map[string]string
			err=json.Unmarshal(buffer,&post)
			if err==nil{
				for k,v:=range post{
					filter.ReqParams[k]=v
				}
			}
		}
	}

	if filter.SourceType== webs.Unknown {
		c.JSON(int(webs.Unauthorized),"请设置User-Agent请求头: 如:iPhone 或者 Android 或则web")
		return  true
	}else{
		token:=""
		sign:=""
		str,ok:=filter.ReqParams[filter.TOKEN_NAME]
		if ok{
			token=str
		}
		str,ok=filter.ReqParams[filter.Sign]
		if ok{
			sign=str
		}
		if token==""{
			c.JSON(int(webs.Unauthorized),"token is empty you are error！")
			return  true
		}else if sign==""{
			c.JSON(int(webs.Unauthorized),"sign is empty you are error！")
			return  true
		}else{
			allow:=false
			for i:=range filter.AllowAnonymouUrls{
				if strings.Contains(c.Request.URL.Path,filter.AllowAnonymouUrls[i]){
					allow=true
				}
			}
			if allow{
				return  false
			}else{

			}
		}
		return  false
	}
}

func (filter *GinFilter) OnActionExecuted(c *gin.Context) bool{
	return  false
}

func (server *WebGinServer)GetPage(c *gin.Context)(int,int)  {
	page:=c.GetInt(":page")
	size:=c.GetInt(":size")
	if page<1{
		page=1
	}else if page>1000{
		page=1000
	}
	if size<1{
		size=10
	}else if size>500{
		size=500
	}
	return page,size
}

func  (server *WebGinServer)BindModel(c *gin.Context,model interface{},equalModel interface{})(interface{},string,bool) {
	/*if userCtl.OnActionExecution(c){
		return
	}*/
	c1,err:=NewGinHttpWeb(c).BindModel(&model)
	if err != nil || reflect.DeepEqual(model, equalModel) {
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail"}
		server.OutputXmlOrJson(c,c1,res)
		return model, c1, false
	}
	return model, c1, true

}
func (server *WebGinServer) BindInt64Id(c *gin.Context)(string,int64,bool)  {
		id:=c.GetInt64(":id")
		//p:= c.Request.URL.Path
		//p=strings.ToLower(p)
		//m:=regexp.MustCompile(p)
		//bs:=m.Find([]byte("/role/remove/(\\d+)"))
		//id,_:=strconv.ParseInt(string(bs),10,64)
		str:="json"
		a:=c.GetHeader("accept")
		if strings.Contains(a,"xml"){
			str="xml"
		}else if strings.Contains(a,"json"){
			str="json"
		}
		if id<1{
			res:=dto.ResponseDto{Status: false, Code: 400, Msg: "id error"}
			server.OutputXmlOrJson(c,str,res)
			return str,id, false
		}
		return str, id,true
}
func (server *WebGinServer) BindInt64Ids(c *gin.Context)(dto.DeleteDto,dto.DeleteXmlDto,string,bool){
	var m dto.DeleteDto
	var m1 dto.DeleteXmlDto
	mm,mm1,str,su:=server.BindIds(c,m,dto.DeleteDto{},m1,dto.DeleteXmlDto{})
	return mm.(dto.DeleteDto),mm1.(dto.DeleteXmlDto),str,su

}
func (server *WebGinServer) BindIds(c *gin.Context,
	model interface{},equalModel interface{},
	xmlModel interface{},xmlEqualModel interface{})(interface{},interface{},
	string,bool)  {
	c1 := "json"
	var err error

	contentType := c.ContentType()
	if strings.Contains(contentType, "json") {
		err = c.ShouldBindJSON(&model)
	} else if strings.Contains(contentType, "form") {
		err = c.ShouldBind(&model)
		c1 = "form"
	} else if strings.Contains(contentType, "xml") {
		c1 = "xml"
		err = c.ShouldBindXML(&xmlModel)
	}
	if  err != nil ||
		( c1=="json"&&reflect.DeepEqual(model, equalModel)||
			( c1=="xml"&&   reflect.DeepEqual(xmlModel, xmlEqualModel) )){
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail"}
		server.OutputXmlOrJson(c,c1,res)
		return model,xmlModel, c1, false
	}
	return model,xmlModel, c1, true
}
func (server *WebGinServer)OutputXmlOrJson(c *gin.Context,c1 string,res interface{})  {
	if c1=="xml"{
		c.XML(http.StatusOK,res)
	}else{
		c.JSON(http.StatusOK,res)
	}

	/*switch c1 {
		case "xml":
			c.XML(http.StatusOK, res)
			break
		case "form":
		case "json":
		default:
			c.JSON(http.StatusOK, res)
			break
	}
	return*/
}