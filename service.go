package gtmesh

import (
	"time"

	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//ServiceMap map of services to hosts
type ServiceMap map[string][]string

//LocalServiceMap map of Local Services
type LocalServiceMap map[string]gotalk.BufferReqHandler

var (
	//LocalServiceStore Hosts the map of all local Services
	LocalServiceStore LocalServiceMap
	//ServiceStore Hosts the mapping of all services in the mesh mapped to their hosts
	ServiceStore ServiceMap
	//ServiceHandlers Bank of Handlers used by GoTalk
	ServiceHandlers *gotalk.Handlers
	//lastServiceUpdateTime
	lastServiceUpdateTime time.Time
)

//AddLocalService add a Local Service to the map
func AddLocalService(service string, handler gotalk.BufferReqHandler) error {
	console.Log("service.go::AddLocalService(service:%#v,handler:%#v)", service, handler)
	if len(service) == 0 {
		return console.Error("service.go::AddLocalService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if handler == nil {
		return console.Error("service.go::AddLocalService(handler:%#v\tError:handler Cannot be nil", handler)
	}
	if LocalServiceStore == nil {
		LocalServiceStore = LocalServiceMap{}
		ServiceHandlers = gotalk.NewHandlers()
	}
	LocalServiceStore[service] = handler
	ServiceHandlers.HandleBufferRequest(service, handler)

	return AddService(service, LocalHost.TCPUrl)
}

//AddService add a service to Service Map
func AddService(service string, tcp string) error {
	console.Log("service.go::AddService(service:%#v,tcp:%s)", service, tcp)
	if len(service) == 0 {
		return console.Error("service.go::AddService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if len(tcp) == 0 {
		return console.Error("service.go::AddService(tcp:%#v\tError:tcp Cannot be empty", tcp)
	}
	if ServiceStore == nil {
		ServiceStore = ServiceMap{}
		ServiceStore[service] = []string{}
	} else {
		if ServiceStore[service] == nil {
			ServiceStore[service] = []string{}
		}
	}
	ServiceStore[service] = append(ServiceStore[service], tcp)

	return nil
}
