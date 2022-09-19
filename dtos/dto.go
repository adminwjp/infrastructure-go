package dtos

import "encoding/xml"

var Response ResponseDto = ResponseDto{}
var Result ApiResult = ApiResult{}
//AccounType
type  AccounType int
const  (
	AccounTypeByPhone AccounType=1
	AccounTypeByEamil AccounType=2
	AccounTypeByUsername AccounType=3
)
type UserLogWay int
const  (
	UserLogByNone UserLogWay=iota
	UserLogByDb
	UserLogByMongo
	UserLogByMq
	UserLogByInfluxdb
	UserLogByEs
	UserLogByHbase
)
//ReturnCode
type ReturnCode int
const  (
	ReturnCodeNone ReturnCode=0
	ReturnCodeExists ReturnCode=202
	ReturnCodeNotExists ReturnCode=404
	ReturnCodeSuccess ReturnCode=1
	ReturnCodeFail ReturnCode=-1
	ReturnCodeError ReturnCode=5
	ReturnCodeUnkow ReturnCode=6
	ReturnCodeJsonFail ReturnCode=-2
	ReturnCodeXmlFail ReturnCode=-4
	ReturnCodeCacheFail ReturnCode=-5
	ReturnCodeLoginFail ReturnCode=-6

)

type ConfigDto struct {
	Rpc string `json:"rpc"`
	EnableRpc bool `json:"enable_rpc"`
	GRpcAddress string `json:"grpc_address"`
	ThriftAddress string `json:"thrift_address"`

	Mq string `json:"mq"`
	EnableMq bool `json:"enable_mq"`
	MqMaster []string `json:"mq_master"`
	MqSlave []string `json:"mq_slave"`
	RabbitMqMaster []string `json:"-"`
	RabbitMqSlave []string `json:"-"`
	KafkaMaster []string `json:"-"`
	KafkaSlave []string `json:"-"`

	Lock string `json:"lock"`
	LockMaster []string `json:"lock_master"`
	LockSlave []string `json:"lock_slave"`
	RedisMaster []string `json:"-"`
	RedisSlave []string `json:"-"`
	ZookeeperMaster []string `json:"-"`
	ZookeeperSlave []string `json:"-"`

	Driver string `json:"driver"`
	DriverMaster []string `json:"driver_master"`
	DriverSlave []string `json:"driver_slave"`
}
type DeleteDto struct {

	Ids []int64 `json:"ids" form:"ids" `
}
type DeleteXmlDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"request" `
	Ids []DeleteIdXmlDto `xml:"ids"`
}
type DeleteIdXmlDto struct {
	Id int64 `xml:"id"`
}
type CountDto struct {
  Total	int64
}
type PageDto struct {
	Page int `uri:"page"`
	Size int `uri:"size"`
}
type ResponseApiDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func CreateResponseApiSuccess() ResponseApiDto {
	return ResponseApiDto{Success: true, Msg: "success", Code: 200}
}

func  CreateResponseApiFail() ResponseApiDto {
	return ResponseApiDto{Success: false, Msg: "fail", Code: 400}
}

func  CreateResponseApiError() ResponseApiDto {
	return ResponseApiDto{Success: false, Msg: "error", Code: 500}
}

type ApiResult struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

func   CreateApiResultSuccess() ApiResult {
	return ApiResult{Status: true, Msg: "success", Code: 200}
}

func  CreateApiResultFail() ApiResult {
	return ApiResult{Status: false, Msg: "fail", Code: 400}
}

func  CreateApiResultError() ApiResult {
	return ApiResult{Status: false, Msg: "error", Code: 500}
}

func Success() map[string]interface{} {
	var res map[string]interface{}= map[string]interface{}{
		"code":200,"status":true,"msg":"success",
	}
	return res
}
func Fail() map[string]interface{} {
	var res map[string]interface{}= map[string]interface{}{
		"code":400,"status":false,"msg":"fail",
	}
	return res
}
func Error() map[string]interface{} {
	var res map[string]interface{}= map[string]interface{}{
		"code":500,"status":false,"msg":"error",
	}
	return res
}