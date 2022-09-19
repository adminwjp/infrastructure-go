package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type  httpUtil struct {

}
type HttpEntity struct {
	Url string
	Method string
	Referer string
	Param string
	ParamByte []byte
	UserAgent string
	AuthUser string
	AuthPwd string
	ContentType string
	Timeout int64 //mill
	Host string
	Accept string
	Cookie string //a=b;
	Cookies []*http.Cookie
}

const  (
	HttpGet="GET"
	HttpPost="POST"
	HttpPut="PUT"
	HttpDelete="DELETE"
	HttpHead="HEAD"
	HttpTrance="TRANCE"
)
const  (
	ContentTypeForm="application/x-www-form-urlencoded"
	ContentTypeJson="application/json"
	ContentTypeFormData="application/x-www-form-data"
)

func NewHttpEntity() *HttpEntity {
	return &HttpEntity{Referer: "",UserAgent: "",AuthUser: "",AuthPwd: "",ContentType: "",Timeout: 5000}
}



func (httpUtil)   GetString(url string) string  {
	res,err :=http.Get(url)
	if err!=nil{
		fmt.Print("get  ex : %s",err.Error())
		return ""
	}
	defer res.Body.Close()
	body,err :=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Print("get  buffer : %s",err.Error())
		return ""
	}
	return string(body)
}

func (httpUtil) Http(httEntity HttpEntity) []byte  {
	return Http(httEntity)
}

func (httpUtil) PostForm(url string,referer string,param string) string  {
	return PostString(url,referer,param,ContentTypeForm)
}

func (httpUtil) PostJson(url string,referer string,json string) string  {
	return PostString(url,referer,json,ContentTypeJson)
}

func (httpUtil) PostFormData(url string,referer string,data string) string  {
	return PostString(url,referer,data,ContentTypeFormData)
}

func GetString(url string,referer string) string  {
	httpEntity :=HttpEntity{Url: url,Referer: referer,Method: HttpGet}
	buffer:= Http(httpEntity)
	if buffer!=nil{
		return string(buffer)
	}
	return ""
}

func PostForm(url string,referer string,param string) string  {
	return PostString(url,referer,param,ContentTypeForm)
}

func PostJson(url string,referer string,json string) string  {
	return PostString(url,referer,json,ContentTypeJson)
}

func PostFormData(url string,referer string,data string) string  {
	return PostString(url,referer,data,ContentTypeFormData)
}

func PostString(url string,referer string,param string, contentType string) string  {
	httpEntity :=HttpEntity{Url: url,Referer: referer,Method: HttpPost,ContentType: contentType}
	buffer:= Http(httpEntity)
	if buffer!=nil{
		return string(buffer)
	}
	return ""
}

func  Http(httpEntity HttpEntity) []byte  {
	client :=http.Client{}
	if httpEntity.Timeout==0{
		httpEntity.Timeout=3000
	}
	/*s:=httpEntity.Timeout/1000
	p,err :=time.ParseDuration(string(s)+"s")
	if err!=nil{
		println("Duration time set fial ",err)
		return nil
	}
	client.Timeout=p*/
	var req *http.Request=&http.Request{}
	req.Method=httpEntity.Method
	u, err := url.Parse(httpEntity.Url)
	if err!=nil {
		println("url parse fial ",err)
		return nil
	}
	req.URL=u
	if httpEntity.Referer!=""{
		req.Header.Set("Referer",httpEntity.Referer)
	}
	if httpEntity.Referer!=""{
		req.Header.Set("Accept",httpEntity.Accept)
	}
	if httpEntity.Referer!=""{
		req.Header.Set("User-Agent",httpEntity.UserAgent)
	}
	if httpEntity.Cookie!=""{
		req.Header.Set("Cookie",httpEntity.Cookie)
	}
	if httpEntity.Referer!=""{
		req.Header.Set("Content-Type",httpEntity.ContentType)
	}
	if httpEntity.AuthUser!=""&&httpEntity.AuthPwd!=""{
		req.SetBasicAuth(httpEntity.AuthUser,httpEntity.AuthPwd)
	}
	if httpEntity.Referer!=""{
		req.Host=httpEntity.Host
	}
	if httpEntity.ParamByte!=nil{
		req.Body=io.NopCloser(bytes.NewReader(httpEntity.ParamByte))
	}
	if httpEntity.Param!=""{
		req.Body=io.NopCloser(strings.NewReader(httpEntity.Param))
	}
	res,err :=client.Do(req)
	if err!=nil{
		fmt.Print("get  ex : %s",err.Error())
		return nil
	}
	cookies :=res.Cookies()
	httpEntity.Cookies=cookies
	/*
	if cookies!=nil&&len(cookies)>0{
		for i := range cookies{
			key:=cookies[i].Name
		}
	}*/
	defer res.Body.Close()
	body,err :=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Print("get  buffer : %s",err.Error())
		return nil
	}
	return body
}


