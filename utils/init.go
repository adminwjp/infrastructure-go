package utils

import (
	"sync"
)


var StringUtil  _stringUtil
var NetUtil _netUtil
var HttpInstance httpUtil
var SeurityInstance seurityUtl
var SeriablizeUtil _seriablizeUtil
var TypeUtil _typeUtil
var RegexUtil _regexUtil
var FileUtil fileUtil
var DateUtil _dateUtil
// 饿汉式 自带 初始化？ init
func init() {
	StringUtil=_stringUtil{}
	NetUtil=_netUtil{}
	SeurityInstance = seurityUtl{BASE64Table: BASE64Table, DesKey: []byte(DesKey), AesKey: AesKey}

	SeriablizeUtil=_seriablizeUtil{}
	//TypeUtil=newTypeUtil()
	RegexUtil=_regexUtil{}
	FileUtil=fileUtil{}
	DateUtil=_dateUtil{}
}



//空方法 加载 改包 不然 用不了 自定义包
/*func Empty() {

}
*/
//懒汉式  自带 初始化？ getInstane
var oSing sync.Once

func getInstane() httpUtil {
	if HttpInstance == (httpUtil{}) {
		//Do函数里面的函数只有在第一次才会被调用
		oSing.Do(func() {
			HttpInstance = httpUtil{}
		})
	}
	return HttpInstance
}

//懒汉式  自带 初始化？ getInstane
/*
var oSing sync.Once
func  getInstane() *seurityUtl  {
	if SeurityInstance==nil{
		//Do函数里面的函数只有在第一次才会被调用
		oSing.Do(func() {
			SeurityInstance=&seurityUtl{BASE64Table:BASE64Table,DesKey:[]byte(DesKey),AesKey:AesKey}
		})
	}
	return  SeurityInstance
}
*/


