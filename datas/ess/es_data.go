package ess

import (
	"context"
	"encoding/json"
	"github.com/adminwjp/infrastructure-go/datas"
	"github.com/adminwjp/infrastructure-go/utils"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)


type EsData struct {
	ElasticUtil
	TranManager *datas.TranManager
	Config *datas.DataConfig
	//gorm gorm jinzhus bee orm mong es
	dataWay string
}
func (data1 *EsData)GetTranction()datas.ITranManager{
	return data1.TranManager
}
func (data1 *EsData) GetConfig()*datas.DataConfig{
	return data1.Config
}
func (data1 *EsData)	Insert(obj interface{})(int,error){
	var entity =obj.(datas.IDataEntity)
	id:=reflect.ValueOf(obj).FieldByName(entity.GetIdProName())
	respnse, err := data1.Client.Index().Index(entity.GetTable()).Id(id.String()).
		BodyJson(obj).Do(context.Background())
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s insert err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else if respnse.Status >0{
			log.Printf("%s %s insert suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s insert fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return respnse.Status ,err
}
func (data1 *EsData)	Update(obj interface{})(int,error){
	var entity =obj.(datas.IDataEntity)
	//if reflect.DeepEqual(entity,reflect.New(reflect.TypeOf(obj))){
	//	return 0, nil
	//}
	id:=reflect.ValueOf(obj).FieldByName(entity.GetIdProName())
	respnse, err := data1.Client.Update().Index(entity.GetTable()).Id(id.String()).
		Doc(obj).Do(context.Background())
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s update err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else if respnse.Status >0{
			log.Printf("%s %s update suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s update fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return respnse.Status ,err
}
func (data1 *EsData)	Delete(obj interface{})(int,error){
	var entity =obj.(datas.IDataEntity)
	id:=reflect.ValueOf(obj).FieldByName(entity.GetIdProName())
	return data1.DeleteById(id.String(),entity)
}
func (data1 *EsData)	DeleteById(id interface{},entity datas.IDataEntity)(int,error){
	respnse, err := data1.Client.Delete().Index(entity.GetTable()).Id(id.(string)).
		Do(context.Background())
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s delete by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
			return 0, err
		}else if respnse.Status >0{
			log.Printf("%s %s delete by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s delete by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return respnse.Status ,err
}

func (data1 *EsData)	DeleteByIds(ids []interface{},entity datas.IDataEntity)(int,error) {
	query := elastic.NewBoolQuery()
	for i := 0; i < len(ids); i++ {
		query.Must(elastic.NewTermQuery(entity.GetIdColName(), ids[i]))
	}
	respnse, err := data1.Client.DeleteByQuery(entity.GetTable()).Query(query).
		Do(context.Background())
	if data1.Config.EnableLog {
		if err != nil {
			log.Printf("%s %s delete by id err,err:%s", entity.GetDescription(), data1.dataWay, err.Error())
			return 0, err
		} else if respnse.Deleted > 0 {
			log.Printf("%s %s delete by id suc", entity.GetDescription(), data1.dataWay)
		} else {
			log.Printf("%s %s delete by id fail", entity.GetDescription(), data1.dataWay)
		}
	}
	return int(respnse.Deleted), err
}
func (data1 *EsData)	Get(id interface{},entity datas.IDataEntity)(interface{},error){

	_, err := data1.Client.Search("admin").Query(nil).
		Source(nil).Do(context.Background())
	respnse, err := data1.Client.Get().Index(entity.GetTable()).Id(id.(string)).
		Do(context.Background())
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s get by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}
	}
	j,err:=respnse.Source.MarshalJSON()
	var obj=reflect.New(reflect.TypeOf(entity))
	utils.SeriablizeUtil.JsonDesriablizeObject(&obj,j)
	if (&obj)!=nil{
		log.Printf("%s %s get by id suc",entity.GetDescription(),data1.dataWay)
	}else{
		log.Printf("%s %s get by id fail",entity.GetDescription(),data1.dataWay)
	}
	return obj,nil
}
func (data1 *EsData)	Count(id interface{},entity datas.IDataEntity)(int,error){
	respnse, err := data1.Client.Count().Index(entity.GetTable()).Q("{\"match_all\":{\"term\"{\"id\":\""+id.(string)+"\"}}}").
		Do(context.Background())
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s count by id err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if respnse>0{
			log.Printf("%s %s count by id suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s count by id fail",entity.GetDescription(),data1.dataWay)
		}
	}
	return int(respnse),nil
}

func (data1 *EsData)	List(all interface{},entity datas.IDataEntity)(interface{},error){
	respnse1, err := data1.Client.Search(entity.GetTable()).Do(context.Background())
	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s List  err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}
	}
	if respnse1.TookInMillis<1{
		return nil, err
	}
	bs:=make([][]byte,respnse1.Hits.TotalHits.Value)
	//t:=reflect.TypeOf(all)
	reflectValue:=reflect.ValueOf(all)
	switch reflectValue.Kind() {
		case reflect.Slice, reflect.Array:{
			value:=reflect.MakeSlice(reflectValue.Type(),0,int(respnse1.Hits.TotalHits.Value))
			for i,v := range respnse1.Hits.Hits {
				bs[i],_=v.Source.MarshalJSON()
				elem:= reflectValue.Index(int(i))
				json.Unmarshal(bs[i],&elem)
			}
			return value.Interface(), err

		}
		default:
			return nil, err
	}
}
func (data1 *EsData)	ListByPage(all interface{},entity datas.IDataEntity,page int,size int)(interface{},int64,error){
	respnse1, err := data1.Client.Search(entity.GetTable()).From((page-1)*size+1).
		Size(size).Do(context.Background())
	if err!=nil{
		return nil, 0, err
	}
	if respnse1.TookInMillis<1{
		return nil, 0, err
	}
	respnse, err := data1.Client.Count().Index(entity.GetTable()).
		Do(context.Background())

	if data1.Config.EnableLog{
		if err!=nil{
			log.Printf("%s %s ListByPage  err,err:%s",entity.GetDescription(),data1.dataWay,err.Error())
		}else if respnse>0{
			log.Printf("%s %s ListByPage  suc",entity.GetDescription(),data1.dataWay)
		}else{
			log.Printf("%s %s ListByPage  fail",entity.GetDescription(),data1.dataWay)
		}
	}
	bs:=make([][]byte,respnse1.Hits.TotalHits.Value)
	//t:=reflect.TypeOf(all)
	reflectValue:=reflect.ValueOf(all)
	switch reflectValue.Kind() {
	case reflect.Slice, reflect.Array:{
		value:=reflect.MakeSlice(reflectValue.Type(),0,int(respnse1.Hits.TotalHits.Value))
		for i,v := range respnse1.Hits.Hits {
			bs[i],_=v.Source.MarshalJSON()
			elem:= reflectValue.Index(int(i))
			json.Unmarshal(bs[i],&elem)
		}
		return value.Interface(),respnse, err

	}
	default:
		return nil, respnse, err
	}
}