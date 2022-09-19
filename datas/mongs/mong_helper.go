package mongs

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

//https://studygolang.com/articles/14055
type MongHelper struct {
	Session *mgo.Session

}
func (data *MongHelper) Conn1(addres []string ,db string, username string,password string)  *mgo.Session {
	var dailInfo *mgo.DialInfo
	/*if util.FileUtil.CheckFileIsExists("config/cfg.ini"){
		cfg:= util.ConfigUtil
		cf:= goconfig.LoadFile("config/cfg.ini")
		addr=cfg.GetStringValue(cf,"Mong","MongIp","127.0.0.1")
		username=cfg.GetStringValue(cf,"Mong","MongUser","")
		username=cfg.GetStringValue(cf,"Mong","MongPwd","")
	}*/
	dailInfo=&mgo.DialInfo{
		Addrs:addres,
		Direct:false,
		Timeout:time.Second*1,
		Database:db,
		Source:"",
		Username:username,
		Password:password,
		PoolLimit:1024,
	}
	session,err:=mgo.DialWithInfo(dailInfo)
	if err!=nil{
		log.Println("mong conn err ,errror:"+err.Error())
		return nil
	}
	data.Session=session
	//defer  session.Clone()
	session.SetMode(mgo.Monotonic,true)
	log.Println("mong conn suc ")
	return  session
}
func (data *MongHelper) Conn()  *mgo.Session {
	addr:="127.0.0.1"
	username:=""
	password:=""
	return data.Conn1([]string{addr},"test",username,password)

}

func (data *MongHelper)  Insert(db string,col string,doc ...interface{}) (bool,error) {
	c:=data.Session.DB(db).C(col)
	err:=c.Insert(doc...)
	return err==nil,err
}

func (data *MongHelper)  Update(db string,col string,where interface{},update interface{}) (bool,error){
	c:=data.Session.DB(db).C(col)
	err:= c.Update(where,update)
	return err==nil,err
}


func (data *MongHelper)  RemoveId(db string,col string,id interface{}) (bool,error) {
	c:=data.Session.DB(db).C(col)
	err:= c.RemoveId(id)
	return err==nil,err
}

func (data *MongHelper)  Remove(db string,col string,where interface{}) (bool,error) {
	c:=data.Session.DB(db).C(col)
	err:= c.Remove(where)
	return err==nil,err
}

func (data *MongHelper)  FindId(db string,col string,id interface{},result interface{}) (interface{},error) {
	c:=data.Session.DB(db).C(col)
	query:= c.FindId(id)
	err:=query.One(&result)
	return result,err
}

func (data *MongHelper)  CountId(db string,col string,id interface{},result interface{}) (int,error) {
	c:=data.Session.DB(db).C(col)
	query:= c.FindId(id)
	return query.Count()
}

