package https

import (
	"encoding/json"
	"fmt"
	dto "github.com/adminwjp/infrastructure-go/dtos"
	"github.com/adminwjp/infrastructure-go/webs"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)
const (
	formatTime      = "15:04:05"
	formatDate      = "2006-01-02"
	formatDateTime  = "2006-01-02 15:04:05"
	formatDateTimeT = "2006-01-02T15:04:05"
)
var sliceOfInts = reflect.TypeOf([]int(nil))
var sliceOfStrings = reflect.TypeOf([]string(nil))
func (server *HttpServerHelper)GetPage(writer http.ResponseWriter, request *http.Request)(int,int)  {
	return  server.GetPageAndSize(request)
}

func  (server *HttpServerHelper)BindModel(writer http.ResponseWriter, request *http.Request,model interface{},equalModel interface{})(interface{},webs.ContentType,bool) {
	/*if userCtl.OnActionExecution(c){
		return
	}*/
	c1 ,err:= BindModel(request,&model)
	if err != nil || reflect.DeepEqual(model, equalModel) {
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail"}
		OutputXmlOrJson(writer,c1,res)
		return model, c1, false
	}
	return model, c1, true

}
func (server *HttpServerHelper) BindInt64Id(writer http.ResponseWriter, request *http.Request)(webs.ContentType,int64,bool)  {




	//p:= c.Request.URL.Path
	//p=strings.ToLower(p)
	//m:=regexp.MustCompile(p)
	//bs:=m.Find([]byte("/role/remove/(\\d+)"))
	//id,_:=strconv.ParseInt(string(bs),10,64)
	str:=webs.Json
	a:=Accept(request)
	if strings.Contains(a,"xml"){
		str=webs.Xml
	}else if strings.Contains(a,"json"){
		str=webs.Json
	}
	var id int64=0
	reg,err:=regexp.Compile(`/(\d+/{0,1})[/|?]{0,1}`)
	if err!=nil{
		log.Println(" find reg pattern faill")
		return str, 0, false
	}else if !reg.MatchString(request.URL.RawPath){
		log.Printf( "%s find reg match faill",request.URL.RawPath)
		return str, 0, false
	}
	strs:=reg.FindAllString(request.URL.RawPath,-1)
	if len(strs)>=1{
		id,_=strconv.ParseInt(strs[0],10,64)
	}
	if id<1{
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "id error"}
		OutputXmlOrJson(writer,str,res)
		return str,id, false
	}
	return str, id,true
}
func (server *HttpServerHelper) BindInt64Ids(writer http.ResponseWriter, request *http.Request)(dto.DeleteDto,dto.DeleteXmlDto,webs.ContentType,bool){
	var m dto.DeleteDto
	var m1 dto.DeleteXmlDto
	//contentType := request.Header.Get("Content-Type")
	/*if strings.Contains(contentType, "form") {
		t := reflect.TypeOf(m)
		v := reflect.ValueOf(m)
		//array ids[0] ids[1] ...
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fo := f.Tag.Get("form")
			va := request.Form.Get(fo)
			if va != "" {
				//type parse
				v.Field(i).SetString(va)
			}
		}
	}*/

		mm,mm1,str,su:=server.BindIds(writer,request,m,dto.DeleteDto{},m1,dto.DeleteXmlDto{})
	return mm.(dto.DeleteDto),mm1.(dto.DeleteXmlDto),str,su

}
func (server *HttpServerHelper) BindIds(writer http.ResponseWriter, request *http.Request,
	model interface{},equalModel interface{},
	xmlModel interface{},xmlEqualModel interface{})(interface{},interface{},
	webs.ContentType,bool)  {
	c1 := webs.Json
	var err error

	contentType := request.Header.Get("Content-Type")
	if strings.Contains(contentType, "json") {
		c1 =  webs.Json
		bs,err:=ioutil.ReadAll(request.Body)
		if err!=nil{
			res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail"}
			OutputXmlOrJson(writer,c1,res)
			return model,xmlModel, c1, false
		}
		err = json.Unmarshal(bs,&model)
	} else if strings.Contains(contentType, "form") {
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail,not support"}
		OutputXmlOrJson(writer,c1,res)
		return model,xmlModel, c1, false
		t:=reflect.TypeOf(model)
		v:=reflect.ValueOf(model)
		for i := 0; i < t.NumField(); i++ {
			f:=t.Field(i)
			fo:=f.Tag.Get("form")
			va:=request.Form.Get(fo)
			if va!=""{
				//type parse
				v.Field(i).SetString(va)
			}
		}
		c1 = webs.Form
	} else if strings.Contains(contentType, "xml") {
		c1 =  webs.Xml
		bs,err:=ioutil.ReadAll(request.Body)
		if err!=nil{
			res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail"}
			OutputXmlOrJson(writer,c1,res)
			return model,xmlModel, c1, false
		}
		err = json.Unmarshal(bs,&model)
	}
	if  err != nil ||
		( c1== webs.Json&&reflect.DeepEqual(model, equalModel)||
			( c1== webs.Xml&&   reflect.DeepEqual(xmlModel, xmlEqualModel) )){
		res:=dto.ResponseDto{Status: false, Code: 400, Msg: "bind fail"}
		OutputXmlOrJson(writer,c1,res)
		return model,xmlModel, c1, false
	}
	return model,xmlModel, c1, true
}

func (*HttpServerHelper) IsPost(writer http.ResponseWriter, request *http.Request) bool{
	if strings.ToLower(request.Method) != "post" {
		io.WriteString(writer, "http method not support,only support post")
		return  false
	}
	return true
}
func (*HttpServerHelper) ReadBuffer(writer http.ResponseWriter, request *http.Request,prefix string )([]byte,bool) {
	reader, err := request.GetBody()
	if err != nil {
		log.Println(prefix + " read body stream fail,error :" + err.Error())
		io.WriteString(writer, "0")
		return nil,false
	}
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(prefix + " read body buffer fail,error :" + err.Error())
		io.WriteString(writer, "0")
		return nil,false
	}
	return buffer,true
}

// ParseForm will parse form values to struct via tag.
// Support for anonymous struct.
func parseFormToStruct(form url.Values, objT reflect.Type, objV reflect.Value) error {
	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		fieldT := objT.Field(i)
		if fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			err := parseFormToStruct(form, fieldT.Type, fieldV)
			if err != nil {
				return err
			}
			continue
		}

		tags := strings.Split(fieldT.Tag.Get("form"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			continue
		} else {
			tag = tags[0]
		}

		formValues := form[tag]
		var value string
		if len(formValues) == 0 {
			defaultValue := fieldT.Tag.Get("default")
			if defaultValue != "" {
				value = defaultValue
			} else {
				continue
			}
		}
		if len(formValues) == 1 {
			value = formValues[0]
			if value == "" {
				continue
			}
		}

		switch fieldT.Type.Kind() {
		case reflect.Bool:
			if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
				fieldV.SetBool(true)
				continue
			}
			if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
				fieldV.SetBool(false)
				continue
			}
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			fieldV.SetString(value)
		case reflect.Struct:
			switch fieldT.Type.String() {
			case "time.Time":
				var (
					t   time.Time
					err error
				)
				if len(value) >= 25 {
					value = value[:25]
					t, err = time.ParseInLocation(time.RFC3339, value, time.Local)
				} else if strings.HasSuffix(strings.ToUpper(value), "Z") {
					t, err = time.ParseInLocation(time.RFC3339, value, time.Local)
				} else if len(value) >= 19 {
					if strings.Contains(value, "T") {
						value = value[:19]
						t, err = time.ParseInLocation(formatDateTimeT, value, time.Local)
					} else {
						value = value[:19]
						t, err = time.ParseInLocation(formatDateTime, value, time.Local)
					}
				} else if len(value) >= 10 {
					if len(value) > 10 {
						value = value[:10]
					}
					t, err = time.ParseInLocation(formatDate, value, time.Local)
				} else if len(value) >= 8 {
					if len(value) > 8 {
						value = value[:8]
					}
					t, err = time.ParseInLocation(formatTime, value, time.Local)
				}
				if err != nil {
					return err
				}
				fieldV.Set(reflect.ValueOf(t))
			}
		case reflect.Slice:
			if fieldT.Type == sliceOfInts {
				formVals := form[tag]
				fieldV.Set(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(int(1))), len(formVals), len(formVals)))
				for i := 0; i < len(formVals); i++ {
					val, err := strconv.Atoi(formVals[i])
					if err != nil {
						return err
					}
					fieldV.Index(i).SetInt(int64(val))
				}
			} else if fieldT.Type == sliceOfStrings {
				formVals := form[tag]
				fieldV.Set(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf("")), len(formVals), len(formVals)))
				for i := 0; i < len(formVals); i++ {
					fieldV.Index(i).SetString(formVals[i])
				}
			}
		}
	}
	return nil
}

// ParseForm will parse form values to struct via tag.
func ParseForm(form url.Values, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !IsStructPtr(objT) {
		return fmt.Errorf("%v must be  a struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()

	return parseFormToStruct(form, objT, objV)
}

func IsStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}