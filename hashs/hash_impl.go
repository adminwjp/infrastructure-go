package hashs

type HashImpl struct {

}
func (hash *HashImpl) GetHash(bytes []byte)(int64,error){
	return  hash.GetDefaultHash(bytes)
}
func (*HashImpl) GetDefaultHash(bytes []byte)(int64,error)  {
	if bytes==nil{
		return 0,nil
	}
	var hash int64=0
	for _, v := range bytes {
		hash=hash+int64(v)
	}
	return  hash,nil
}

