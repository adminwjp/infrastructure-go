package rpcs

type IRpc interface {
	StartServer()error
	StartClient()error
}
