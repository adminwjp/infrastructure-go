package datas


type Data interface {
	 GetTranction()ITranManager
	 Insert(obj interface{})(int,error)

	 Update(obj interface{})(int,error)

	 Delete(obj interface{})(int,error)

	 DeleteById(id interface{},entity IDataEntity)(int,error)

	DeleteByIds(ids []interface{},entity IDataEntity)(int,error)

	 Get(id interface{},entity IDataEntity)(interface{},error)

	 Count(id interface{},entity IDataEntity)(int,error)
	 List(all interface{},entity IDataEntity)(interface{},error)

	 ListByPage(all interface{},entity IDataEntity,page int,size int)(interface{},int64,error)

	 GetConfig()*DataConfig
}
