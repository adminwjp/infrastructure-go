package bee_orms

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-oci8"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type BeeOrmConfig struct {
	DriverNames map[string]bool
	AliasNames map[string]bool
}

func (beeOrmConfig *BeeOrmConfig)Register()  {

}

func (beeOrmConfig *BeeOrmConfig) UpdateDb(dialect datas.DbFalg,connectionString ,aliasName string,debug bool) {

	if debug{
		//sql show
		orm.Debug=true
	}
	name:="sqlite3"
	switch dialect {
	case datas.DbSqlite:
		_,e:=beeOrmConfig.DriverNames["sqlite3"]
		if !e{
			orm.RegisterDriver("sqlite3", orm.DRSqlite)
			beeOrmConfig.DriverNames["sqlite3"]=true
		}
		break
	case datas.DbMysql:
		name="mysql"
		_,e:=beeOrmConfig.DriverNames["mysql"]
		if !e{
			orm.RegisterDriver("mysql", orm.DRMySQL)
			beeOrmConfig.DriverNames["mysql"]=true
		}
		break
	case datas.DbTidb:
		name="tidb"
		_,e:=beeOrmConfig.DriverNames["tidb"]
		if !e{
			orm.RegisterDriver("tidb", orm.DRTiDB)
			beeOrmConfig.DriverNames["tidb"]=true
		}
		break
	case datas.DbOracle:
		name="oracle"
		_,e:=beeOrmConfig.DriverNames["oracle"]
		if !e{
			orm.RegisterDriver("oracle", orm.DROracle)
			beeOrmConfig.DriverNames["oracle"]=true
		}
		break
	case datas.DbSqlserver:
		name="sqlserver"
		break
	case datas.DbPostgre:
		name="postgres"
		_,e:=beeOrmConfig.DriverNames["postgres"]
		if !e{
			orm.RegisterDriver("postgres", orm.DRPostgres)
			beeOrmConfig.DriverNames["postgres"]=true
		}
		break
	default:
		break

	}

	if aliasName==""{
		aliasName="default"
	}
	_,e:=beeOrmConfig.AliasNames[aliasName]
	if e{
		return
	}
	 err := orm.RegisterDataBase(aliasName, name, connectionString)
	if err!= nil{
		log.Printf("bee orm dialect => %s connectionString => %s, connection database fail",
			dialect,connectionString)
		panic(err)
	}
	beeOrmConfig.AliasNames[aliasName]=true
	//orm.RunSyncdb("default", false, true)
	//orm.DefaultTimeLoc = time.UTC
}