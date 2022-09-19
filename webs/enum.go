package webs

type ContentType int
const (
	Json ContentType=iota
	Form
	Xml
)


const  (
	ContentTypeJson="application/json"
	ContentTypeForm="application/x-www-form-urlencoded"
	ContentTypeXml="application/xml"
)


type AccountStatus int
//账号状态
const  (
	Disabled AccountStatus = 0 //禁用
	Normal AccountStatus = 1 //正常/注册
	Blacklist //黑名单
)

//账号类型
type AccountType int
const  (
	AdminRole AccountType= 1 //超级管理员
	JustSoSoAdminRole AccountType= 2 //普通管理员
)

//认证方式
type AuthType int
const  (
	AutoApprove AuthType = 1 //自动认证
	PersonApprove  AuthType= 2 //人工认证
)


//错误码
type ErrorCode int
const  (
	Unauthorized ErrorCode= 403 //未授权
	SystemError ErrorCode= 503 //系统错误
	ReLogin ErrorCode= 10001 //请重新登录
	InvalidToken ErrorCode= 10002 //非法token
	ErrorSign ErrorCode= 10003 //sign 签名非法
	ErrorUserNameOrPass  //用户名或密码有误
	NotFound  //不存在
	Forbidden  //禁止
	InvalidPassword  //无效密码
	AccountDisabled  //账户禁用
	InvalidData  //非法数据
	HasValued ErrorCode = 20001 //数据已存在
)

//上传文件类别枚举
type FileCatagory int
const  (
	Head FileCatagory= 1 //头像
	IdCardFace  //身份证正面照片
	IdCardBack  //身份证反面照片
	Feedbacks  //意见反馈
	Store  //店铺
	Message  //消息
	Description  //其他
)

//来源类型
type SourceType int
const  (
	Unknown SourceType= 0 //未知
	Web  //网站
	WeChat  //微信
	Android  //Android
	IOS  //iOS
	WeChatApp  //小程序
)