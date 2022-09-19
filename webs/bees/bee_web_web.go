package bees

import (
	"github.com/adminwjp/infrastructure-go/webs/https"
	"github.com/beego/beego/v2/server/web"
)



type BeeWebHttpWeb struct {
	web.Controller
	BeeHttpWeb
}

func (web *BeeWebHttpWeb)Update()  {
	web.C=web.Ctx
}

func (web *BeeWebHttpWeb)Get(key string)string{
	return https.Get(web.C.Request,key)
}

func (web *BeeWebHttpWeb)ShouldBindForm(bind interface{})error{
	return web.ParseForm(bind)
}

func (web *BeeWebHttpWeb)XML(code int,data interface{}){
	web.Ctx.ResponseWriter.Status=code
	web.SetData(data)
	web.ServeXML()
}
func (web *BeeWebHttpWeb)Json(code int,data interface{}){
	web.Ctx.ResponseWriter.Status=code
	web.SetData(data)
	web.ServeJSON()
}
func (web *BeeWebHttpWeb)Html(code int,name string,data interface{}){
	web.Ctx.ResponseWriter.Status=code

}
func (web *BeeWebHttpWeb)Jsonp(code int,data interface{}){
	web.Ctx.ResponseWriter.Status=code
	web.SetData(data)
	web.ServeJSONP()
}

