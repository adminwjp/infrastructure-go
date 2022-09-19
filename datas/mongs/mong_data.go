package mongs

import (
	"github.com/adminwjp/infrastructure-go/datas"
	"log"
)


type MongData struct {
	MongHelper
	TranManager *datas.TranManager
	Config *datas.DataConfig
	//gorm gorm jinzhus bee orm mong es
	dataWay string
}
func (data1 *MongData)GetTranction()datas.ITranManager{
	return data1.TranManager
}
func (data1 *MongData) GetConfig()*datas.DataConfig{
	return data1.Config
}
func (data1 *MongData)	Insert(obj interface{})(int,error){
	var entity =obj.(datas.IDataEntity)
	err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Insert(obj)
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s insert err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else {
			log.Printf("%s %s insert suc",entity.GetDescription(),data1.dataWay)
		}
	}
	return 1,err
}
func (data1 *MongData)	Update(obj interface{})(int,error){
	var entity =obj.(datas.IDataEntity)
	err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Update(map[string]interface{}{entity.GetIdColName():entity.GetId()},obj)
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s update err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else{
			log.Printf("%s %s update suc",entity.GetDescription(),data1.dataWay)
		}
	}
	return 1,err
}
func (data1 *MongData)	Delete(obj interface{})(int,error){
	var entity =obj.(datas.IDataEntity)
	//err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).RemoveId(obj.GetId())
	err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Remove(map[string]interface{}{entity.GetIdColName():entity.GetId()})
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s delete err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else {
			log.Printf("%s %s delete suc",entity.GetDescription(),data1.dataWay)
		}
	}
	return 1,err
}
func (data1 *MongData)	DeleteById(id interface{},entity datas.IDataEntity)(int,error){
	//err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).RemoveId(id)
	err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Remove(map[string]interface{}{entity.GetIdColName():id})
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else {
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}
	}
	return 1,err
}
func (data1 *MongData)	DeleteByIds(ids []interface{},entity datas.IDataEntity)(int,error){
	err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Remove(map[string]interface{}{"$in":map[string]interface{}{entity.GetIdColName():ids}})
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else {
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}
	}
	return 1,err
}
func (data1 *MongData)	Get(id interface{},entity datas.IDataEntity)(interface{},error){
	//err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).FindId(obj.GetId())
	err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Find(map[string]interface{}{entity.GetIdColName():id}).One(&entity)
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
func (data1 *MongData)	Count(id interface{},entity datas.IDataEntity)(int,error){
	//err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).FindId(obj.GetId())
	r,err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Find(map[string]interface{}{entity.GetIdColName():id}).Count()
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s count by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s count by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s count by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return r, err
}
func (data1 *MongData)	List(all interface{},entity datas.IDataEntity)(interface{},error){
	//err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).FindId(obj.GetId())
	query:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Find(map[string]interface{}{})
	err:=query.All(&all)
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s List err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if (&all)!=nil{
			log.Printf("%s %s List suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s List fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &all,err
}
func (data1 *MongData)	ListByPage(all interface{},entity datas.IDataEntity,page int ,size int)(interface{},int64,error){
	//err:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).FindId(obj.GetId())
	query:=data1.Session.DB(entity.GetDb()).C(entity.GetTable()).Find(map[string]interface{}{})
	err:=query.Skip((page-1)*size).Limit(size).All(&all)
	r,err:=query.Count()
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s ListByPage err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if r>0{
			log.Printf("%s %s ListByPage suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s ListByPage fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return &all,int64(*&r), err
}