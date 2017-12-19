package gtmesh

import (
	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//ServiceMap map of services to hosts
type ServiceMap map[string][]string

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

	return node.AddService(service, node.LocalHost.TCPUrl)
}

//AddService add a service to Service Map
func (node *Node) AddService(service string, tcp string) error {
	console.Log("service.go::AddService(service:%#v,tcp:%s)", service, tcp)
	if len(service) == 0 {
		return console.Error("service.go::AddService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if len(tcp) == 0 {
		return console.Error("service.go::AddService(tcp:%#v\tError:tcp Cannot be empty", tcp)
	}
	if node.ServiceStore == nil {
		node.ServiceStore = ServiceMap{}
		node.ServiceStore[service] = []string{}
	} else {
		if node.ServiceStore[service] == nil {
			node.ServiceStore[service] = []string{}
		}
	}

	node.ServiceStore[service] = append(node.ServiceStore[service], tcp)

	return nil
}

//BufferRequest Request for a Buffer/[]byte based service
func (node *Node) BufferRequest(serviceName string, payLoad []byte) ([]byte, error) {
	//Basic Sanity Check
	if len(serviceName) == 0 {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName cannot be empty", serviceName, string(payLoad))
	}
	//Find where the service is located
	if node.ServiceStore[serviceName] == nil {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName Not Registered yet! %#v", serviceName, string(payLoad), node.ServiceStore)
	}

	//Iterate through each host till service result is obtained
	for _, host := range node.ServiceStore[serviceName] {
		s, err := gotalk.Connect("tcp", host)
		if err != nil {
			console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with %s and returned error %s", serviceName, payLoad, host, err.Error())
		} else {
			return s.BufferRequest(serviceName, payLoad)
		}
	}

	return nil, console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with any hosts [%#v]", serviceName, payLoad, node.ServiceStore[serviceName])
}
