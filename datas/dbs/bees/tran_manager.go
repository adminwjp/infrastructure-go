package bees

import (
	"github.com/beego/beego/v2/client/orm"
)

type TranManager struct {
   Tx orm.TxOrmer
   Db orm.Ormer
   has bool
   trans map[string]*orm.Ormer
}
func (tr *TranManager) New()*TranManager{
	return &TranManager{Db: tr.Db}
}
func (tr *TranManager) GetDb()orm.QueryExecutor{
	if tr.Tx!=nil{
		return tr.Tx
	}
	if tr.Db==nil{
		tr.Db=orm.NewOrm()
	}
	return tr.Db
}

/*func (tr *TranManager) GetDb()orm.TxOrmer{
	if tr.Tx!=nil{
		return tr.Tx
	}
	tx,_:=tr.Db.Begin()
	return tx
}*/
func (tr *TranManager)Begin()  {
	if tr.has{
		return
	}
	tr.GetDb()
	tr.Tx,_=tr.Db.Begin()
	tr.has=true
}

func (tr *TranManager)Commit()  {
	if tr.Tx!=nil{
		tr.Tx.Commit()
		tr.Tx=nil
		tr.has=false
	}
}

func (tr *TranManager)Rollback()  {
	if tr.Tx!=nil{
		tr.Tx.Rollback()
		tr.Tx=nil
		tr.has=false
	}
}