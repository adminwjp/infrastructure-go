package utils

import "strconv"

type _typeUtil struct {
	GrpcTypes map[string]string
	ThriftTypes map[string] string
}

func newTypeUtil()_typeUtil  {
	var typeUtil=  _typeUtil{GrpcTypes: make(map[string]string,100),
		ThriftTypes: make(map[string]string,100)}
	/*typeUtil.GrpcTypes[reflect.Type(byte(0)).Name()]="int32"
	typeUtil.GrpcTypes[reflect.Type(int(0)).Name()]="int32"
	typeUtil.GrpcTypes[reflect.Type(int8(0)).Name()]="int32"
	typeUtil.GrpcTypes[reflect.Type(int32(0)).Name()]="int32"
	typeUtil.GrpcTypes[reflect.Type(int64(0)).Name()]="int64"
	typeUtil.GrpcTypes[reflect.Type(float32(0)).Name()]="float"
	typeUtil.GrpcTypes[reflect.Type(float64(0)).Name()]="double"
	typeUtil.GrpcTypes["bool"]="bool"
	typeUtil.GrpcTypes["string"]="string"
	typeUtil.GrpcTypes[reflect.Type(time.Now()).Name()]="bytes"

	typeUtil.ThriftTypes[reflect.Type(byte(0)).Name()]="i8"
	typeUtil.ThriftTypes[reflect.Type(int(0)).Name()]="i32"
	typeUtil.ThriftTypes[reflect.Type(int8(0)).Name()]="i8"
	typeUtil.ThriftTypes[reflect.Type(int32(0)).Name()]="i32"
	typeUtil.ThriftTypes[reflect.Type(int64(0)).Name()]="i64"
	typeUtil.ThriftTypes[reflect.Type(float32(0)).Name()]="double"
	typeUtil.ThriftTypes[reflect.Type(float64(0)).Name()]="double"
	typeUtil.ThriftTypes["bool"]="bool"
	typeUtil.ThriftTypes["string"]="string"
	typeUtil.ThriftTypes[reflect.Type(time.Now()).Name()]="binary"
	*/
	return  typeUtil
}
func (typeUtil *_typeUtil) GetGrpc(type1 string) string  {
	var t=typeUtil.GrpcTypes[type1]
	return t
}
func (typeUtil *_typeUtil) ToInt64(num string) int64  {
	n,_:=strconv.ParseInt(num,10,64)
	return  n
}
func (typeUtil *_typeUtil) ToInt32(num string) int32  {
	n,_:=strconv.Atoi(num)
	return  int32(n)
}
func (typeUtil *_typeUtil) ToBool(bool1 string) bool  {
	b,_:=strconv.ParseBool(bool1)
	return  b
}
func (typeUtil *_typeUtil) ToFloat(num string) float64  {
	n,_:=strconv.ParseFloat(num,10)
	return  n
}