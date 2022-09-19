package jinzhus

import "github.com/jinzhu/gorm"

type TranManager struct {
   Tx *gorm.DB
   Db *gorm.DB
   has bool
   trans map[string]*gorm.DB
}
func (tr *TranManager) New()*TranManager{
	return &TranManager{Db: tr.Db}
}
func (tr *TranManager) SetDb(db *gorm.DB){
	if tr.Tx!=nil{
		tr.Tx=db
	}
}
func (tr *TranManager) GetDb()*gorm.DB{
	if tr.Tx!=nil{
		return tr.Tx
	}
	return tr.Db
}
func (tr *TranManager)Begin()  {
	if tr.has{
		return
	}
	tr.Tx=tr.Db.Begin()
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