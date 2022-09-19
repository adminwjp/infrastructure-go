package datas

var TranManagerEmpty=&TranManager{}
type TranManager struct {

}

func (tr *TranManager)Begin()  {
}

func (tr *TranManager)Commit()  {

}

func (tr *TranManager)Rollback()  {

}

type ITranManager interface {
	Begin()
	Commit()
	Rollback()
}

