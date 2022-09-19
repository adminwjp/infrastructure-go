package configs

import (
	"fmt"
	"github.com/adminwjp/infrastructure-go/datas"
)


func GetConnectionString(dialect datas.DbFalg,db string,ip string,port int,user string,pwd string)string  {
	switch dialect {
	case datas.DbSqlite:
		return db
	case datas.DbMysql:
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		//?charset=utf8mb4&parseTime=True&loc=Local
		//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",user,pwd,ip,port,db)
	case datas.DbTidb:
	case datas.DbOracle:
		//https://www.jb51.net/article/251496.htm
		return fmt.Sprintf("%s/%s@%s:%d/XE",user,pwd,ip,port)
	case datas.DbSqlserver:
		// github.com/denisenkom/go-mssqldb
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",user,pwd,ip,port,db)
	case datas.DbPostgre:
		//"host=myhost port=myport user=gorm dbname=gorm password=mypassword"
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",ip,user,pwd,db,port)
	default:
		break

	}
	return ""
}

type EmptyConfig struct {


}
func (config *EmptyConfig) UseMq()bool{
	return false
}
func (config *EmptyConfig)UseEs()bool{
	return false
}
func (config *EmptyConfig)UseMong()bool{
	return false
}
func (config *EmptyConfig)UseDb()bool{
	return false
}

type Config interface {
	UseMq()bool
	UseEs()bool
	UseMong()bool
	UseDb()bool
	UseLog()bool
}
