package bees

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/beego/beego/v2/client/orm"
	"log"
)

type DbDataBeeOrm struct {
	Tx orm.TxOrmer
	Db orm.Ormer
	TranManager *TranManager
	Config *datas.DataConfig
	//gorm gorm jinzhus bee orm mong es
	dataWay string
}
func (data1 *DbDataBeeOrm)GetTranction()datas.ITranManager{
	return data1.TranManager
}
func (data1 *DbDataBeeOrm) GetConfig()*datas.DataConfig{
	return data1.Config
}
func (data1 *DbDataBeeOrm)	Insert(obj interface{})(int,error){
	r,err:=data1.TranManager.GetDb().Insert(obj)
	if data1.Config.EnableLog{
		var entity =obj.(datas.IDataEntity)
		if err!=nil{
			log.Printf("%s %s insert err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s insert suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s insert fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	Update(obj interface{})(int,error){
	r,err:=data1.TranManager.GetDb().Update(obj)
	if data1.Config.EnableLog{
		var entity =obj.(datas.IDataEntity)
		if err!=nil{
			log.Printf("%s %s update err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s update suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s update fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	Delete(obj interface{})(int,error){
	r,err:=data1.TranManager.GetDb().Delete(obj)
	if data1.Config.EnableLog{
		var entity =obj.(datas.IDataEntity)
		if err!=nil{
			log.Printf("%s %s delete err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s delete suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	DeleteById(id interface{},entity datas.IDataEntity)(int,error){
	r,err:=data1.TranManager.GetDb().QueryTable(entity).Filter(entity.GetIdColName(),id).Delete()
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete  by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	DeleteByIds(ids []interface{},entity datas.IDataEntity)(int,error){
	r,err:=data1.TranManager.GetDb().QueryTable(entity).Filter(entity.GetIdColName()+"__in",ids).Delete()
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete  by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	Get(id interface{},entity datas.IDataEntity)(interface{},error){
	err:=data1.TranManager.GetDb().QueryTable(entity).Filter(entity.GetIdColName(),id).One(&entity)
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s get by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if (&entity)!=nil{
			log.Printf("%s %s get by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s get by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &entity, err
}
func (data1 *DbDataBeeOrm)	Count(id interface{},entity datas.IDataEntity)(int,error){
	r,err:=data1.TranManager.GetDb().QueryTable(entity).Filter(entity.GetIdColName(),id).Count()
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s count by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s count by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s count by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	List(all interface{},entity datas.IDataEntity)(interface{},error){
	r,err:=data1.TranManager.GetDb().QueryTable(entity).All(&all)
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s List err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s List suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s List fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(r), err
}
func (data1 *DbDataBeeOrm)	ListPage(all interface{},entity datas.IDataEntity,page int,size int)(interface{},int64,error){
	query:=data1.TranManager.GetDb().QueryTable(entity)
	query.Limit((page-1)*size).Offset(size).All(&all)
	r,err:=query.Count()
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s ListPage err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s ListPage suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s ListPage fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &all,r, err
}