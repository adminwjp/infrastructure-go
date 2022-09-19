package https

import (
	"encoding/json"
	"encoding/xml"
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpWeb struct {
	Writer http.ResponseWriter
	Request *http.Request
	ContentType1 webs.ContentType
	Accept1 webs.ContentType
}

func NewHttpWeb(writer http.ResponseWriter,request *http.Request)*HttpWeb  {
	return &HttpWeb{Writer: writer,Request: request}
}
func (web *HttpWeb)ContentType() string{
	c:= ContentType(web.Request)
	web.ContentType1=webs.ParseContentType(c)
	return  c
}

func (web *HttpWeb)UserAgent() string{
	return UserAgent(web.Request)
}
func (web *HttpWeb)Accept()string{
	a:= Accept(web.Request)
	web.Accept1=webs.ParseAccept(a)
	return  a
}
func (web *HttpWeb)ClientIP() string{
	return  web.Request.RemoteAddr
}

func (web *HttpWeb)GetHeader(key string)string{
	return  web.Request.Header.Get(key)
}

func (web *HttpWeb)GetQueryString(key string)string{
	return  web.Request.URL.Query().Get(key)
}

// /v1/1 v1/:id id = > 1
func (web *HttpWeb)GetPathString(key string)string{
	return ""
}
func (web *HttpWeb)GetPathInt64(key string)int64{
	return  0
}
func ContentType(request *http.Request) string{
	return request.Header.Get("Content-Type")
}
func UserAgent(request *http.Request) string{
	return request.UserAgent()
}
func Accept(request *http.Request)string{
	return request.Header.Get("Accept")
}
func Get(request *http.Request,key string)string{
	return ""
}
func Method(request *http.Request)string{
	return strings.ToLower(request.Method)
}
func (web *HttpWeb)Get(key string)string{
	return Get(web.Request,key)
}
func (web *HttpWeb)ShouldBindIds(json1 *dtos.DeleteDto,xml1 *dtos.DeleteXmlDto)error{
	web.ContentType()
	switch web.ContentType1 {
	case webs.Json:return web.ShouldBindJSON(json1)
	case webs.Form:return web.ShouldBindForm(json1)
	case webs.Xml:return web.ShouldBindXML(xml1)
	default:
		return web.ShouldBindJSON(json1)
	}
}
func (web *HttpWeb)ShouldBind(bind interface{})error{
	web.ContentType()
	switch web.ContentType1 {
	case webs.Json:return web.ShouldBindJSON(bind)
	case webs.Form:return web.ShouldBindForm(bind)
	case webs.Xml:return web.ShouldBindXML(bind)
	default:
		return web.ShouldBindJSON(bind)
	}
}
func (web *HttpWeb)ShouldBindUri(bind interface{})error{
	return ShouldBindUri(web.Request,bind)
}
func (web *HttpWeb)ShouldBindJSON(bind interface{})error{
	return ShouldBindJSON(web.Request,bind)
}
func (web *HttpWeb)ShouldBindForm(bind interface{})error{
	return ShouldBindForm(web.Request,bind)
}
func (web *HttpWeb)ShouldBindXML(bind interface{})error{
	return ShouldBindXML(web.Request,bind)
}
func (web *HttpWeb)XML(code int,data interface{}){
	XML(web.Writer,code,data)
}
func (web *HttpWeb)Json(code int,data interface{}){
	Json(web.Writer,code,data)
}
func (web *HttpWeb)Html(code int,name string,data interface{}){
	Html(web.Writer,code,name,data)
}
func (web *HttpWeb)Jsonp(code int,data interface{}){
	Jsonp(web.Writer,code,data)
}
func (web *HttpWeb)Response(code int,data interface{}){
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
func (web *HttpWeb)BindModel(model interface{})(webs.ContentType,error){
	return BindModel(web.Request,model)
}
func (web *HttpWeb)OutputXmlOrJson(c1 webs.ContentType,res interface{})  {
	 OutputXmlOrJson(web.Writer,c1,res)
}



func ShouldBindUri(request *http.Request,bind interface{})error{
	return nil
}
func ShouldBindJSON(request *http.Request,bind interface{})error{
	bs,err:=ioutil.ReadAll(request.Body)
	if err!=nil{
		return err
	}
	return json.Unmarshal(bs,&bind)
}
func ShouldBindForm(request *http.Request,bind interface{})error{
	return ParseForm(request.Form,bind)
}
func ShouldBindXML(request *http.Request,bind interface{})error{
	bs,err:=ioutil.ReadAll(request.Body)
	if err!=nil{
		return err
	}
	return xml.Unmarshal(bs,&bind)
}
func XML(writer http.ResponseWriter,code int,data interface{}){
	JsonOrXml(writer,code,webs.Xml,data)
}
func Json(writer http.ResponseWriter,code int,data interface{}){
	JsonOrXml(writer,code,webs.Json,data)
}
func JsonOrXml(writer http.ResponseWriter,code int,c1 webs.ContentType,data interface{}){
	writer.WriteHeader(code)
	OutputXmlOrJson(writer,c1,data)
}
func Html(writer http.ResponseWriter,code int,name string,data interface{}){

}
func Jsonp(writer http.ResponseWriter,code int,data interface{}){

}
func BindModel(request *http.Request,model interface{})(webs.ContentType,error){
	c1 := webs.Json
	var err error
	contentType := ContentType(request)
	if strings.Contains(contentType, "json") {
		err = ShouldBindJSON(request,&model)
	} else if strings.Contains(contentType, "form") {
		err = ShouldBindForm(request,&model)
		c1 = webs.Form
	} else if strings.Contains(contentType, "xml") {
		c1 = webs.Xml
		err = ShouldBindXML(request,&model)
	}
	return  c1,err
}
func OutputXmlOrJson(writer http.ResponseWriter,c1 webs.ContentType,res interface{})  {
	switch c1 {
		case webs.Xml:
			bs,err:=xml.Marshal(res)
			if err!=nil{
				writer.Write([]byte("<response><status>false</status><msg>system error</msg><response>"))
			}else{
				writer.Write(bs)
			}
			break
		case webs.Form,webs.Json:
			bs,err:=json.Marshal(res)
			if err!=nil{
				writer.Write([]byte("{\"status\":false,\"msg\":\"system error\"}"))
			}else{
				writer.Write(bs)
			}
			break
	}
	return
}