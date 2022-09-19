package jinzhus

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"

	//gorm.io/gorm
	//_ "github.com/mattn/go-oci8"
	//_ "github.com/CengSin/oracle"
)

type GormConfig struct {
	Db *gorm.DB
}

func (gormConfig *GormConfig)Register()  {

}
func (gormConfig *GormConfig) UpdateDb(dialect datas.DbFalg,connectionString string,debug bool ) *gorm.DB{
	var db *gorm.DB
	var err error
	switch dialect {
	case datas.DbSqlite:
		db, err = gorm.Open("sqlite3", connectionString)
		break
	case datas.DbMysql:
		db, err = gorm.Open("mysql", connectionString)
		break
	case datas.DbTidb:
		db, err = gorm.Open("tidb", connectionString)
		break
	case datas.DbOracle:
		db, err = gorm.Open("oracle", connectionString)
		break
	case datas.DbSqlserver:
		db, err = gorm.Open("mssql", connectionString)
		break
	case datas.DbPostgre:
		db, err = gorm.Open("postgres", connectionString)
		break
	default:
		return nil

	}
	if err!= nil{
		log.Printf(" gorm jinzhus dialect => %s connectionString => %s, connection database fail",
			dialect,connectionString)
		return nil
	}
	log.Printf(" gorm jinzhus dialect => %s connectionString => %s, connection database suc",
		dialect,connectionString)
	/*if strings.Contains(dialect,"sqlite"){
		//https://www.5axxw.com/questions/content/4enbno
		db=db.Clauses(clause.Insert{Modifier: "ignore"})
	}*/
	if debug{
		db=db.Debug()
	}
	//db=db.Logger

	gormConfig.Db=db
	return  db
}