package datas
type IDataEntity interface {
	GetIdProName() string
	GetId() interface{}
	GetIdColName() string
	GetDescription() string
	GetDb() string
	GetTable() string
}


