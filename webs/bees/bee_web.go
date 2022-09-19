package bees

import (
	"encoding/json"
	"encoding/xml"
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	 "github.com/adminwjp/infrastructure-go/webs/https"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
)


type BeeHttpWeb struct {
	C *context.Context
	ContentType1 webs.ContentType
	Accept1 webs.ContentType
}

func NewBeeHttpWeb(c *context.Context)*BeeHttpWeb  {
	return &BeeHttpWeb{C: c}
}
func (web *BeeHttpWeb)ContentType() string{
	c:= https.ContentType(web.C.Request)
	web.ContentType1=webs.ParseContentType(c)
	return  c
}
func (web *BeeHttpWeb)UserAgent() string{
	return https.UserAgent(web.C.Request)
}
func (web *BeeHttpWeb)Accept()string{
	a:= https.Accept(web.C.Request)
	web.Accept1=webs.ParseAccept(a)
	return  a
}
func (web *BeeHttpWeb)Get(key string)string{
	return https.Get(web.C.Request,key)
}
func (web *BeeHttpWeb)ClientIP() string{
	return  web.C.Request.RemoteAddr
}

func (web *BeeHttpWeb)GetHeader(key string)string{
	return  web.C.Request.Header.Get(key)
}

func (web *BeeHttpWeb)GetQueryString(key string)string{
	return  web.C.Input.Query(key)
}

// /v1/1 v1/:id id = > 1
func (web *BeeHttpWeb)GetPathString(key string)string{
	return  web.C.Input.Query(key)
}
func (web *BeeHttpWeb)GetPathInt64(key string)int64{
	i,_:=strconv.ParseInt(web.C.Input.Query(key),10,64)
	return  i
}
func (web *BeeHttpWeb)ShouldBindUri(bind interface{})error{
	return https.ShouldBindUri(web.C.Request,bind)
}
func (web *BeeHttpWeb)ShouldBindJSON(bind interface{})error{
	//web.C.Input.Bind(bind,"")
	return https.ShouldBindJSON(web.C.Request,bind)
}
func (web *BeeHttpWeb)ShouldBindForm(bind interface{})error{
	return https.ShouldBindForm(web.C.Request,bind)
}
func (web *BeeHttpWeb)ShouldBindXML(bind interface{})error{
	return https.ShouldBindXML(web.C.Request,bind)
}
func (web *BeeHttpWeb)XML(code int,data interface{}){
	 web.C.ResponseWriter.Status=code
	 OutputXmlOrJson(web.C,"xml",data)
}
func (web *BeeHttpWeb)Json(code int,data interface{}){
	web.C.ResponseWriter.Status=code
	OutputXmlOrJson(web.C,"json",data)
}
func (web *BeeHttpWeb)Html(code int,name string,data interface{}){

}
func (web *BeeHttpWeb)Jsonp(code int,data interface{}){

}
func (web *BeeHttpWeb)Response(code int,data interface{}){
	account:=web.Accept()
	a:=webs.ParseAccept(account)
	switch a {
	case webs.Xml:
		web.XML(code,data)
		break
	default:
		web.Json(code,data)
		break
	}
}
func (web *BeeHttpWeb)ShouldBindIds(json1 *dtos.DeleteDto,xml1 *dtos.DeleteXmlDto)error{
	web.ContentType()
	switch web.ContentType1 {
	case webs.Json:return web.ShouldBindJSON(json1)
	case webs.Form:return web.ShouldBindForm(json1)
	case webs.Xml:return web.ShouldBindXML(xml1)
	default:
		return web.ShouldBindJSON(json1)
	}
}
func (web *BeeHttpWeb)ShouldBind(bind interface{})error{
	web.ContentType()
	switch web.ContentType1 {
	case webs.Json:return web.ShouldBindJSON(bind)
	case webs.Form:return web.ShouldBindForm(bind)
	case webs.Xml:return web.ShouldBindXML(bind)
	default:
		return web.ShouldBindJSON(bind)
	}
}
func (web *BeeHttpWeb)BindModel(model interface{})(webs.ContentType,error){
	return https.BindModel(web.C.Request,model)
}
func (web *BeeHttpWeb)OutputXmlOrJson(c1 string,res interface{})  {
	 OutputXmlOrJson(web.C,c1,res)
}

func OutputXmlOrJson(c *context.Context,c1 string,res interface{})  {
	if c1=="xml"{
		bs,err:=xml.Marshal(res)
		if err!=nil{
			c.WriteString("<response><status>false</status><msg>system error</msg><response>")
		}else{
			c.ResponseWriter.Write(bs)
		}
	}else{
		bs,err:=json.Marshal(res)
		if err!=nil{
			c.WriteString("{\"status\":false,\"msg\":\"system error\"}")
		}else{
			c.ResponseWriter.Write(bs)
		}
	}

	/*switch c1 {
		case "xml":
			bs,err:=xml.Marshal(res)
			if err!=nil{
				c.WriteString("<response><status>false</status><msg>system error</msg><response>")
			}else{
				c.ResponseWriter.Write(bs)
			}
			break
		case "form":
		case "json":
		default:
			bs,err:=json.Marshal(res)
			if err!=nil{
				c.WriteString("{\"status\":false,\"msg\":\"system error\"}")
			}else{
				c.ResponseWriter.Write(bs)
			}
			break
	}
	return*/
}