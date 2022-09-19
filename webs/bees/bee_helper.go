package bees

import (
	"github.com/adminwjp/infrastructure-go/register_services"
	consul_register_service "github.com/adminwjp/infrastructure-go/register_services/consuls"
	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/session"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
	"time"
)

var BeeServerInstance=&BeeServer{}
type BeeServer struct {

}
func (server *BeeServer)RegisterCor()  {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: 	[]string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders:	[]string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
}
func  (server *BeeServer) RegisterConsul() func() {
	port:=beego.AppConfig.DefaultInt("Port",8701)

	ip:=beego.AppConfig.DefaultString("Ip","192.168.1.4")

	consulIp:=beego.AppConfig.DefaultString("ConsulIp","192.168.1.4")

	serviceName:=beego.AppConfig.DefaultString("ServiceName","test.api")

	consulPort:=beego.AppConfig.DefaultInt("ConsulPort",8500)

	consulTag:=beego.AppConfig.DefaultString("ConsulTag","test.api,go, beego ")

	consul,err:=consul_register_service.NewConsulServiceRegistry(consulIp,consulPort,"")
	if err!=nil{
		log.Fatal("consul get instace fail,error:"+err.Error())
	}
	reg:=consul.Register(register_services.ServiceInfo{Id: uuid.New().String(),Host: ip,Ip:ip,Port:port,ServiceName:serviceName ,Tags: strings.Split(consulTag,",")})
	if !reg {
		log.Fatal("consul register fail")
	}
	return func() {
		consul.Deregister()
	}
}

func (server *BeeServer) RegisterDb()  {
	/*
		orm.DRMySQL
		orm.DRSqlite
		orm.DRPostgres
		orm.DRTiDB

		// < 1.6
		orm.DR_MySQL
		orm.DR_Sqlite`
		orm.DR_Postgres
	*/
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wjp930514.@tcp(192.168.1.3:3306)/samplesystem?charset=utf8")

	//orm.DR_Sqlite undefined
	//orm.RegisterDriver("sqlite3", orm.DRSqlite)
	//orm.RegisterDataBase("default", "sqlite3", "E:/work/db/sqlite/bee.sqlite3")

	//orm.RunSyncdb("default", false, true)
	orm.DefaultTimeLoc = time.UTC
}
var GlobalSessions *session.Manager
var bm cache.Cache
func (server *BeeServer) RegisterCache()  {
	//beego.ViewsPath = "configs/tpl"
	//beego.AutoRender = false
	logs.SetLogger("console")
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	var err error
	bm, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		log.Println("create beego cache fail,err :" + err.Error())
	} else {
		log.Println("create beego cache suc ")
	}
	GlobalSessions, err = session.NewManager("memory", sessionConfig)
	if err != nil {
		log.Println("create beego session fail,err :" + err.Error())
	} else {
		log.Println("create beego session suc ")
	}
	go GlobalSessions.GC()
}