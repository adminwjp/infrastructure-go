package gins

import (
	"github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type GinHttpWeb struct {
  C	*gin.Context
  ContentType1 webs.ContentType
  Accept1 webs.ContentType
}

func NewGinHttpWeb(c *gin.Context)*GinHttpWeb  {
	return &GinHttpWeb{C: c}
}
func (web *GinHttpWeb)ContentType() string{
	c:= web.C.ContentType()
	web.ContentType1=webs.ParseContentType(c)
	return c
}
func (web *GinHttpWeb)UserAgent() string{
	return web.C.Request.UserAgent()
}
func (web *GinHttpWeb)Accept()string{
	a:= web.C.GetHeader("Accept")
	web.Accept1=webs.ParseAccept(a)
	return  a
}
func (web *GinHttpWeb)ClientIP() string{
	return  web.C.ClientIP()
}
func (web *GinHttpWeb)Get(key string)(interface{},bool){
	return web.C.Get(key)
}

func (web *GinHttpWeb)GetHeader(key string)string{
	return web.C.GetHeader(key)
}

func (web *GinHttpWeb)GetQueryString(key string)string{
	s,_:= web.C.GetQuery(key)
	return  s
}
// /v1/1 v1/:id id = > 1
func (web *GinHttpWeb)GetPathString(key string)string{
	return web.C.Param(strings.Replace(key,":","",-1))
}
func (web *GinHttpWeb)GetPathInt64(key string)int64{
	i,_:= strconv.ParseInt(web.C.Param(strings.Replace(key,":","",-1)),10,64)
	return  i
}

func (web *GinHttpWeb)ShouldBindUri(bind interface{})error{
	return web.C.ShouldBindUri(bind)
}
func (web *GinHttpWeb)ShouldBindIds(json1 *dtos.DeleteDto,xml1 *dtos.DeleteXmlDto)error{
	web.ContentType()
	switch web.ContentType1 {
	case webs.Json:return web.ShouldBindJSON(json1)
	case webs.Form:return web.ShouldBindForm(json1)
	case webs.Xml:return web.ShouldBindXML(xml1)
	default:
		return web.ShouldBindJSON(json1)
	}
}
func (web *GinHttpWeb)ShouldBind(bind interface{})error{
	web.ContentType()
	switch web.ContentType1 {
	case webs.Json:return web.ShouldBindJSON(bind)
	case webs.Form:return web.ShouldBindForm(bind)
	case webs.Xml:return web.ShouldBindXML(bind)
	default:
		return web.ShouldBindJSON(bind)
	}
}
func (web *GinHttpWeb)ShouldBindJSON(bind interface{})error{
	return web.C.ShouldBindJSON(bind)
}
func (web *GinHttpWeb)ShouldBindForm(bind interface{})error{
	return web.C.ShouldBind(bind)
}
func (web *GinHttpWeb)ShouldBindXML(bind interface{})error{
	return web.C.ShouldBindXML(bind)
}
func (web *GinHttpWeb)Response(code int,data interface{}){
	web.Accept()
	switch web.Accept1 {
	case webs.Json: web.Json(code,data)
		break
	case webs.Xml: web.XML(code,data)
		break
	default:
		 web.Json(code,data)
		break
	}
}
func (web *GinHttpWeb)XML(code int,data interface{}){
	 web.C.XML(code,data)
}
func (web *GinHttpWeb)Json(code int,data interface{}){
	web.C.JSON(code,data)
}
func (web *GinHttpWeb)Html(code int,name string,data interface{}){
	web.C.HTML(code,name,data)
}
func (web *GinHttpWeb)Jsonp(code int,data interface{}){
	web.C.JSONP(code,data)
}
func (web *GinHttpWeb)BindModel(model interface{})(string,error){
	c1 := "json"
	var err error
	contentType := web.ContentType()
	if strings.Contains(contentType, "json") {
		err = web.ShouldBindJSON(&model)
	} else if strings.Contains(contentType, "form") {
		err = web.ShouldBindForm(&model)
		c1 = "form"
	} else if strings.Contains(contentType, "xml") {
		c1 = "xml"
		err = web.ShouldBindXML(&model)
	}
	return  c1,err
}
