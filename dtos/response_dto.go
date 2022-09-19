package dtos

import "encoding/xml"

type ResponseDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Status bool `json:"status" xml:"status"`
	Code int `json:"code" xml:"code"`
	Msg string `json:"msg" xml:"msg"`
}
type ResponseTokenDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Status bool `json:"status" xml:"status"`
	Code int `json:"code" xml:"code"`
	Msg string `json:"msg" xml:"msg"`
	Token string `json:"token" xml:"token"`
	Expired int64 `json:"expired" xml:"expired"`
}
type ResponseTokenAndRefreshTokenDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Status bool `json:"status" xml:"status"`
	Code int `json:"code" xml:"code"`
	Msg string `json:"msg" xml:"msg"`
	Token string `json:"token" xml:"token"`
	Expired int64 `json:"expired" xml:"expired"`
	RefreshToken string `json:"refresh_token" xml:"refresh_token"`
	RefreshExpired int64 `json:"refresh_expired" xml:"refresh_expired"`
}
type ResponseDataDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Status bool `json:"status" xml:"status"`
	Code int `json:"code" xml:"code"`
	Msg string `json:"msg" xml:"msg"`
	Data interface{} `json:"data" xml:"data"`
}
type ResponsePageDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Status bool `json:"status" xml:"status"`
	Code int `json:"code" xml:"code"`
	Msg string `json:"msg" xml:"msg"`
	Data interface{} `json:"data" xml:"data"`
	Page int `json:"page" xml:"page"`
	Size int `json:"size" xml:"size"`
	Total int `json:"total" xml:"total"`
	Records int64 `json:"records" xml:"records"`
}
type ResponsePageListDto struct {
	XMLName xml.Name ` json:"-" form:"-"  xml:"response" `
	Status bool `json:"status" xml:"status"`
	Code int `json:"code" xml:"code"`
	Msg string `json:"msg" xml:"msg"`
	List interface{} `json:"list" xml:"list"`
	Page int `json:"page" xml:"page"`
	Size int `json:"size" xml:"size"`
	Total int `json:"total" xml:"total"`
	Records int64 `json:"records" xml:"records"`
}
var ResponseEmpty= ResponseDto{}
var ResponseFail= ResponseDto{Status: false,Code:400,Msg:"fail"}
var ResponseSuccess= ResponseDto{Status: true,Code:200,Msg:"success"}
var ResponseAccountError= ResponseDto{Status: false,Code:402,Msg:"account error"}