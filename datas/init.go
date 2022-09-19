package datas

type  DataFlag int
const(
	DataNone DataFlag=iota
	DataDb
	DataMong
	DataEs
)

type DbFalg int

const  (
	DbNone DbFalg=iota
	DbSqlite
	DbMysql
	DbSqlserver
	DbPostgre
	DbOracle
	DbTidb
)

type DbOrmFlag int

const  (
	DbOrmNone DbOrmFlag=iota
	DbOrmBee
	DbOrmGormio
	DbOrmJinzhuGorm
)
