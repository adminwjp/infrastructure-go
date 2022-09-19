package register_services

type Service interface {
	GetId()string
	GetIp()string
	GetPort()int
	GetIsHttps()bool
	GetServiceName()string
	GetHost()string
}
//https://www.cnblogs.com/zyndev/p/13811589.html
type ServiceInfo struct {
	Id string
	Ip string
	Port int
	IsHttps bool
	ServiceName string
	Host string
	Tags []string
}
func (service *ServiceInfo)GetId()string{
	return service.Id
}
func (service *ServiceInfo)GetIp()string{
	return service.Ip
}
func (service *ServiceInfo)GetPort()int{
	return service.Port
}
func (service *ServiceInfo)GetIsHttps()bool{
	return service.IsHttps
}
func (service *ServiceInfo)GetServiceName()string{
	return service.ServiceName
}
func (service *ServiceInfo)GetHost()string{
	return service.Host
}
type ServiceRegistry interface {
	Register(serviceInstance ServiceInfo) bool

	Deregister()
}

type RegisterService interface {
	Get ()*Service
	Set (service *Service)
	Register()(bool,error)
	UnRegister()(bool,error)
}
