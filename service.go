package gtmesh

import (
	"time"

	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//ServiceData Meta Data of the Host Entry
// version : 22dec2017
type ServiceData struct {
	Created time.Time
	Deleted time.Time
	Source  string
}

//ServiceMap map of services to hosts
// version : 22dec2017
type ServiceMap map[string]map[string]ServiceData

//LocalServiceMap map of Local Services
// version : 22dec2017
type LocalServiceMap map[string]gotalk.BufferReqHandler

//AddLocalService add a Local Service to the map
// version : 22dec2017
func (node *Node) AddLocalService(service string, handler gotalk.BufferReqHandler) error {
	console.Log("service.go::AddLocalService(service:%#v,handler:%#v)", service, handler)
	if len(service) == 0 {
		return console.Error("service.go::AddLocalService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if handler == nil {
		return console.Error("service.go::AddLocalService(handler:%#v\tError:handler Cannot be nil", handler)
	}

	if node.ServiceHandlers == nil {
		node.ServiceHandlers = gotalk.NewHandlers()
	}
	node.ServiceHandlers.HandleBufferRequest(service, handler)

	if node.LocalServiceStore == nil {
		node.LocalServiceStore = LocalServiceMap{}
	}
	node.LocalServiceStore[service] = handler

	return node.addService(service, node.LocalHost.TCPUrl)
}

//addService add a service to Service Map
// version : 22dec2017
func (node *Node) addService(service string, tcp string) error {
	console.Log("service.go::addService(service:%#v,tcp:%s)", service, tcp)
	if len(service) == 0 {
		return console.Error("service.go::addService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if len(tcp) == 0 {
		return console.Error("service.go::addService(tcp:%#v\tError:tcp Cannot be empty", tcp)
	}
	if node.ServiceStore == nil { //Store Does not exist so Init
		node.ServiceStore = make(map[string]map[string]ServiceData)
		node.ServiceStore[service] = make(map[string]ServiceData)
	}
	if node.ServiceStore[service] == nil { //Service Entry Does not Exist
		node.ServiceStore[service] = make(map[string]ServiceData)
	}
	node.ServiceStore[service][tcp] = ServiceData{Created: time.Now().UTC(), Source: node.LocalHost.TCPUrl}
	return nil
}

//BufferRequest Request for a Buffer/[]byte based service
// version : 22dec2017
func (node *Node) BufferRequest(serviceName string, payLoad []byte) ([]byte, error) {
	console.Log("service.go::BufferRequest(serviceName: %s,payload:%s)", serviceName, string(payLoad))
	//Basic Sanity Check
	if len(serviceName) == 0 {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName cannot be empty", serviceName, string(payLoad))
	}
	//Find where the service is located
	if node.ServiceStore[serviceName] == nil {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName Not Registered yet! %#v", serviceName, string(payLoad), node.ServiceStore)
	}

	//Iterate through each host till service result is obtained
	for host := range node.ServiceStore[serviceName] {
		s, err := gotalk.Connect("tcp", host)
		if err != nil {
			console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with %s and returned error %s", serviceName, payLoad, host, err.Error())
		} else {
			return s.BufferRequest(serviceName, payLoad)
		}
	}

	return nil, console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with any hosts [%#v]", serviceName, payLoad, node.ServiceStore[serviceName])
}
