package hashs

type Hash interface {
	GetHash(bytes []byte)(int64,error)

}
