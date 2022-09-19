package gorms

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

	"gorm.io/driver/sqlite"
	//_ "gorm.io/driver/oracle"
	"log"
	//gorm.io/gorm
	//_ "github.com/mattn/go-oci8"
	//_ "github.com/CengSin/oracle"

	//"github.com/cengsin/oracle"
)

type GormConfig struct {
	Db *gorm.DB
}

func (gormConfig *GormConfig)Register()  {

}
func (gormConfig *GormConfig) UpdateDb(dialect datas.DbFalg,connectionString string,debug bool) *gorm.DB{

	var dia gorm.Dialector
	switch dialect {
		case datas.DbSqlite:
			dia=sqlite.Open(connectionString)
			break
		case datas.DbMysql:
			dia=mysql.Open(connectionString)
			break
		case datas.DbTidb:

			break
		case datas.DbOracle:
			// dia=oracle.Open(connectionString)
			break
		case datas.DbSqlserver:
			dia=sqlserver.Open(connectionString)
			break
		case datas.DbPostgre:
			dia=postgres.Open(connectionString)
			break
		default:
			break

	}
	db, err := gorm.Open(dia, &gorm.Config{})
	if err!= nil{
		log.Printf(" gorm jinzhus dialect => %s connectionString => %s, connection database fail",
			dialect,connectionString)
		panic(err)
	}
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