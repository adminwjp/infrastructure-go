package xorms

//https://gitea.com/xorm/xorm
import (
	"xorm.io/xorm"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-oci8"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var XOrmConfigInstance=&XOrmConfig{}
type XOrmConfig struct {
	Db *xorm.Engine
}

func (cfg *XOrmConfig) UpdateDb(dialect string,connectionString string,debug bool) {
	engine ,err := xorm.NewEngine(dialect, connectionString)
	if err != nil {
		log.Fatal(engine, err)
		return
	}
	if debug{
		engine.ShowSQL(true)
	}
	cfg.Db= engine
}
