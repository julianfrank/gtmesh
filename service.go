package gtmesh

import (
	"time"

	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//ServiceHostMap Structure for Each Host Map in each Service
type ServiceHostMap struct {
	TimeStamp time.Time
	Map       map[string]time.Time //string = host
}

//ServiceMap map of services to hosts
type ServiceMap struct {
	TimeStamp time.Time
	Map       map[string]ServiceHostMap //string = service
}

//LocalServiceMap map of Local Services
type LocalServiceMap map[string]gotalk.BufferReqHandler

//AddLocalService add a Local Service to the map
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
func (node *Node) addService(service string, tcp string) error {
	console.Log("service.go::addService(service:%#v,tcp:%s)", service, tcp)
	if len(service) == 0 {
		return console.Error("service.go::addService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if len(tcp) == 0 {
		return console.Error("service.go::addService(tcp:%#v\tError:tcp Cannot be empty", tcp)
	}
	if node.ServiceStore.Map == nil { //Store Does not exist so Create
		node.ServiceStore.Map = make(map[string]ServiceHostMap)
		node.ServiceStore.Map[service] = ServiceHostMap{
			TimeStamp: time.Now().UTC(),
			Map:       make(map[string]time.Time),
		}
	}
	node.ServiceStore.TimeStamp = time.Now().UTC()
	node.ServiceStore.Map[service].TimeStamp = time.Now().UTC()
	node.ServiceStore.Map[service].Map[tcp] = time.Now().UTC()
	console.Log("ZZZZZZZZZZZZZZZZZZZZZZZZZz       %#v", node.ServiceStore.Map[service].TimeStamp)
	//node.ServiceStore.Map[service].TimeStamp = time.Now().UTC()
	//node.ServiceStore.Map[service].Map[tcp] = time.Now().UTC()
	return nil
}

//BufferRequest Request for a Buffer/[]byte based service
func (node *Node) BufferRequest(serviceName string, payLoad []byte) ([]byte, error) {
	console.Log("service.go::BufferRequest(serviceName: %s,payload:%s)", serviceName, string(payLoad))
	//Basic Sanity Check
	if len(serviceName) == 0 {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName cannot be empty", serviceName, string(payLoad))
	}
	//Find where the service is located
	/*if node.ServiceStore.Map[serviceName] == nil {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName Not Registered yet! %#v", serviceName, string(payLoad), node.ServiceStore)
	}*/

	//Iterate through each host till service result is obtained
	/*for _, host := range node.ServiceStore.Map[serviceName] {
		s, err := gotalk.Connect("tcp", host.URL)
		if err != nil {
			console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with %s and returned error %s", serviceName, payLoad, host.URL, err.Error())
		} else {
			return s.BufferRequest(serviceName, payLoad)
		}
	}*/

	return nil, console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with any hosts [%#v]", serviceName, payLoad, node.ServiceStore.Map[serviceName])
}
