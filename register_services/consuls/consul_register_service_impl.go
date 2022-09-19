package consuls

import (
	"errors"
	"fmt"
	"github.com/adminwjp/infrastructure-go/register_services"
	"github.com/hashicorp/consul/api"
	"strconv"
	"unsafe"
)
type ConsulRegisterService struct{
	consulServiceRegistry
}
func (registerService *ConsulRegisterService)Connect(host string, port int, token string)  error{
	client, err := newConsul(host,port,token)
	registerService.client=client
	return err
}
//https://www.cnblogs.com/zyndev/p/13811589.html
type consulServiceRegistry struct {
	serviceInstances     map[string]map[string]register_services.ServiceInfo
	client               *api.Client
	localServiceInstance register_services.ServiceInfo
}

func (c *consulServiceRegistry) Register(serviceInstance register_services.ServiceInfo) bool {
	// 创建注册到consul的服务到
	registration := new(api.AgentServiceRegistration)
	registration.ID = serviceInstance.Id
	registration.Name = serviceInstance.ServiceName
	registration.Port = serviceInstance.Port
	var tags []string=serviceInstance.Tags
	registration.Tags = tags

	registration.Address = serviceInstance.Ip

	// 增加consul健康检查回调函数
	check := new(api.AgentServiceCheck)

	schema := "http"
	if serviceInstance.IsHttps {
		schema = "https"
	}
	//check.HTTP = fmt.Sprintf("%s://%s:%d/actuator/health", schema, registration.Address, registration.Port)
	//check.HTTP = fmt.Sprintf("%s://%s:%d/actuator/health", schema, registration.Address, registration.Port)
	check.HTTP = fmt.Sprintf("%s://%s:%d/test", schema, registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "20s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err := c.client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if c.serviceInstances == nil {
		c.serviceInstances = map[string]map[string]register_services.ServiceInfo{}
	}

	services := c.serviceInstances[serviceInstance.Id]

	if services == nil {
		services = map[string]register_services.ServiceInfo{}
	}

	services[serviceInstance.Id] = serviceInstance

	c.serviceInstances[serviceInstance.Id] = services

	c.localServiceInstance = serviceInstance

	return true
}

// deregister a service
func (c *consulServiceRegistry) Deregister() {
	if c.serviceInstances == nil {
		return
	}

	services := c.serviceInstances[c.localServiceInstance.Id]

	if services == nil {
		return
	}

	delete(services, c.localServiceInstance.Id)

	if len(services) == 0 {
		delete(c.serviceInstances, c.localServiceInstance.Id)
	}

	_ = c.client.Agent().ServiceDeregister(c.localServiceInstance.Id)

	//c.localServiceInstance = nil
}

// new a consulServiceRegistry instance
// token is optional
func NewConsulServiceRegistry(host string, port int, token string) (*consulServiceRegistry, error) {
	client, err := newConsul(host,port,token)
	return &consulServiceRegistry{client: client}, err
}

func newConsul(host string, port int, token string) (*api.Client,error){
	if len(host) < 3 {
		return nil, errors.New("check host")
	}

	if port <= 0 || port > 65535 {
		return nil, errors.New("check port, port should between 1 and 65535")
	}

	config := api.DefaultConfig()
	config.Address = host + ":" + strconv.Itoa(port)
	config.Token = token
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return client,nil
}
type DiscoveryClient interface {

	/**
	 * Gets all ServiceInstances associated with a particular serviceId.
	 * @param serviceId The serviceId to query.
	 * @return A List of ServiceInstance.
	 */
	GetInstances(serviceId string) ([]register_services.ServiceInfo, error)

	/**
	 * @return All known service IDs.
	 */
	GetServices() ([]string, error)
}

type consulDiscoveryClient struct {
	serviceInstances     map[string]map[string]register_services.ServiceInfo
	client               api.Client
	localServiceInstance register_services.ServiceInfo
}

func (c *consulDiscoveryClient) GetInstances(serviceId string) ([]register_services.ServiceInfo, error) {
	catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	if len(catalogService) > 0 {
		result := make([]register_services.ServiceInfo, len(catalogService))
		for index, sever := range catalogService {
			s := register_services.ServiceInfo{
				Id: sever.ServiceID,
				ServiceName:  sever.ServiceName,
				Host:       sever.Address,
				Port:       sever.ServicePort,
				//Metadata:   sever.ServiceMeta,
			}
			result[index] = s
		}
		return result, nil
	}
	return nil, nil
}

func (c *consulDiscoveryClient) GetServices() ([]string, error) {
	services, _, _ := c.client.Catalog().Services(nil)
	result := make([]string, unsafe.Sizeof(services))
	index := 0
	for serviceName, _ := range services {
		result[index] = serviceName
		index++
	}
	return result, nil
}

// new a consulServiceRegistry instance
// token is optional
func NewConsulDiscoveryClient(host string, port int, token string) (*consulDiscoveryClient, error) {
	client, err := newConsul(host,port,token)
	return &consulDiscoveryClient{client: *client}, err
}
