package utils

import (
	"encoding/json"
	"encoding/xml"
)

type _seriablizeUtil struct {

}
func (_seriablizeUtil) ObjectSeriablizeXml(obj interface{})([]byte,error)  {
	return  xml.MarshalIndent(obj, "#", " ")
}
func (_seriablizeUtil) XmlDesriablizeObject(obj interface{},data []byte)error  {
	return  xml.Unmarshal(data,obj)
}
func (_seriablizeUtil) SeriablizeXml(obj interface{})([]byte)  {
	bu,_:=  xml.MarshalIndent(obj, "#", " ")
	return  bu
}
func (_seriablizeUtil) SeriablizeXmlString(obj interface{})string {
	bu,_:=  xml.MarshalIndent(obj, "#", " ")
	if bu==nil{return ""}
	return string(bu)
}
func (_seriablizeUtil) XmlDesriablize(obj interface{},data []byte)  {
	  xml.Unmarshal(data,obj)
}

func (_seriablizeUtil) ObjectSeriablizeJson(obj interface{})([]byte,error)  {
	return  json.Marshal(obj)
}
func (_seriablizeUtil) JsonDesriablizeObject(obj interface{},data []byte)error  {
	return  json.Unmarshal(data,obj)
}
func (_seriablizeUtil) SeriablizeJson(obj interface{})([]byte)  {
	bu,_:=   json.Marshal(obj)
	return  bu
}
func (_seriablizeUtil) SeriablizeJsonString(obj interface{})string  {
	bu,_:=   json.Marshal(obj)
	if bu==nil{return ""}
	return string(bu)
}
