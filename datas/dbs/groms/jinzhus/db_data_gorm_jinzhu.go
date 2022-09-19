package jinzhus

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/jinzhu/gorm"
	"log"
)

type DbDataGormJinzhu struct {
	Db	*gorm.DB
	TranManager *TranManager
	Config *datas.DataConfig
	//gorm gorm jinzhus bee orm mong es
	dataWay string
}
func (data1 *DbDataGormJinzhu)GetTranction()datas.ITranManager{
	return data1.TranManager
}
func (data1 *DbDataGormJinzhu) GetConfig()*datas.DataConfig{
	return data1.Config
}
func (data1 *DbDataGormJinzhu)	Insert(obj interface{})(int,error){
	db:=data1.TranManager.GetDb().Create(obj)
	if data1.Config.EnableLog{
		var entity =obj.(datas.IDataEntity)
		if db.Error!=nil{
			log.Printf("%s %s insert err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s insert suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s insert fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(db.RowsAffected), db.Error
}
func (data1 *DbDataGormJinzhu)	Update(obj interface{})(int,error){
	db:=data1.TranManager.GetDb().Save(obj)
	if data1.Config.EnableLog{
		var entity =obj.(datas.IDataEntity)
		if db.Error!=nil{
			log.Printf("%s %s update err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s update suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s update fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(db.RowsAffected), db.Error
}
func (data1 *DbDataGormJinzhu)	Delete(obj interface{})(int,error){
	db:=data1.TranManager.GetDb().Delete(obj)
	if data1.Config.EnableLog{
		var entity =obj.(datas.IDataEntity)
		if db.Error!=nil{
			log.Printf("%s %s delete err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s delete suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(db.RowsAffected), db.Error
}
func (data1 *DbDataGormJinzhu)	DeleteById(id interface{},entity datas.IDataEntity)(int,error){
	db:=data1.TranManager.GetDb().Where(entity.GetIdColName()+"=?",id).Delete(entity)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(db.RowsAffected), db.Error
}
func (data1 *DbDataGormJinzhu)	DeleteByIds(ids []interface{},entity datas.IDataEntity)(int,error){
	db:=data1.TranManager.GetDb().Where(entity.GetIdColName()+"in(?)",ids).Delete(entity)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(db.RowsAffected), db.Error
}
func (data1 *DbDataGormJinzhu)	Get(id interface{},entity datas.IDataEntity)(interface{},error){
	db:=data1.TranManager.GetDb().Where(entity.GetIdColName()+"=?",id).Scan(&entity)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s get by id err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s get by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s get by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &entity, db.Error
}
func (data1 *DbDataGormJinzhu)	Count(id interface{},entity datas.IDataEntity)(int,error){
	var count int64
	db:=data1.TranManager.GetDb().Model(entity).Where(entity.GetIdColName()+"=?",id).Count(&count)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s count by id err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s count by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s count by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(*&count), db.Error
}
func (data1 *DbDataGormJinzhu)	List(all interface{},entity datas.IDataEntity)(interface{},error){
	db:=data1.TranManager.GetDb().Model(entity).Find(&all)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s List  err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s List suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s List  fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &all, db.Error
}
func (data1 *DbDataGormJinzhu)	ListByPage(all interface{},entity datas.IDataEntity,page int,size int)(interface{},int64,error){
	var count int64
	db:=data1.TranManager.GetDb().Model(entity)
	db=db.Find(&all)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s ListByPage err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
			return &all,*&count, db.Error
		}else if db.RowsAffected>0{
			log.Printf("%s %s ListByPage suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s ListByPage fail",entity.GetDescription(),data1.dataWay)
		}
	}
	db=db.Count(&count)
	if data1.Config.EnableLog{
		if db.Error!=nil{
			log.Printf("%s %s ListByPage count err,err:%s",entity.GetDescription(),data1.dataWay,db.Error.Error())
		}else if db.RowsAffected>0{
			log.Printf("%s %s ListByPage count suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s ListByPage count fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &all,*&count, db.Error
}