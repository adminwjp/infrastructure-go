package webs

import (
	"encoding/json"
	"encoding/xml"
	"github.com/adminwjp/infrastructure-go/dtos"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type HttpWeb interface {
	ContentType() string
	Accept() string
	ClientIP() string

	GetHeader(key string)string

	GetQueryString(key string)string

	// /v1/1 v1/:id id = > 1
	GetPathString(key string)string
	GetPathInt64(key string)int64

	ShouldBindUri(bind interface{})error

	ShouldBind(bind interface{})error

	ShouldBindIds(json1 *dtos.DeleteDto,xml1 *dtos.DeleteXmlDto)error

	ShouldBindJSON(bind interface{})error

	ShouldBindForm(bind interface{})error

	ShouldBindXML(bind interface{})error

	Response(code int,data interface{})

	XML(code int,data interface{})
	Json(code int,data interface{})
	Html(code int,name string,data interface{})
	Jsonp(code int,data interface{})
}
var Success= map[string]interface{}{
	"code":200,"status":true,"msg":"success",
}
var Fail map[string]interface{}= map[string]interface{}{
	"code":400,"status":false,"msg":"fail",
}
var Error= map[string]interface{}{
	"code":500,"status":false,"msg":"error",
}

func ParseAccept(accept string) ContentType {
	c:=Json
	if strings.Contains(accept,ContentTypeJson){
		c=Json
	}else if strings.Contains(accept,ContentTypeForm){
		c=Form
	}else if strings.Contains(accept,ContentTypeXml){
		c=Xml
	}
	return c
}
func   CheckFileIsExists(file string)bool  {
	if fi,err:=os.Stat(file);err==nil{
		fi.Size()
		return true
	}
	return  false
}
func  GetFile(request *http.Request,response http.ResponseWriter,fileName string){
	if !CheckFileIsExists(fileName) {return}
	file, err := os.Open(fileName)
	if err != nil {
		response.WriteHeader(404)
		//io.WriteString(response, "")
		return
	}
	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	fileStat, err := file.Stat()
	if err != nil {
		io.WriteString(response, "")
		return
	}
	request.Header.Set("Content-Disposition", "attachment; filename="+fileName)
	request.Header.Set("Content-Type", http.DetectContentType(fileHeader))
	request.Header.Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))
	file.Seek(0, 0)
	io.Copy(response, file)
}
func GetContentType(request *http.Request) string{
	return request.Header.Get("Content-Type")
}
func GetUserAgent(request *http.Request) string{
	return request.UserAgent()
}

func GetAccept(request *http.Request)string{
	return request.Header.Get("Accept")
}
func ParseContentType(contentType string)ContentType{
	c1 := Json
	if strings.Contains(contentType, ContentTypeJson) {

	} else if strings.Contains(contentType, ContentTypeForm) ||strings.Contains(contentType, "form") {
		c1 = Form
	} else if strings.Contains(contentType, ContentTypeXml) ||strings.Contains(contentType, "xml") {
		c1 = Xml
	}
	return  c1
}
func ShouldBindJSON(request *http.Request,bind interface{})error{
	bs,err:=ioutil.ReadAll(request.Body)
	if err!=nil{
		return err
	}
	return json.Unmarshal(bs,&bind)
}

func ShouldBindXML(request *http.Request,bind interface{})error{
	bs,err:=ioutil.ReadAll(request.Body)
	if err!=nil{
		return err
	}
	return xml.Unmarshal(bs,&bind)
}

