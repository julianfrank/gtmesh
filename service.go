package gtmesh

import (
	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//Service service struct
type Service struct {
	ServiceName string `json:"service_name"`
	Sys         bool   `json:"Sys"` //Set true if the service is a gtmesh service
}

//ServiceMap map of services to hosts
type ServiceMap map[Service][]string

//LocalServiceMap map of Local Services
type LocalServiceMap map[Service]gotalk.BufferReqHandler

var (
	//LocalServiceStore Hosts the map of all local Services
	LocalServiceStore LocalServiceMap
	//ServiceStore Hosts the mapping of all services in the mesh mapped to their hosts
	ServiceStore ServiceMap
)

//AddLocalService add a Local Service to the map
func AddLocalService(service Service, handler gotalk.BufferReqHandler) error {
	if len(service.ServiceName) == 0 {
		return console.Error("service.go::AddLocalService(service:%#v\tError:service.ServiceName Cannot be empty", service)
	}
	if handler == nil {
		return console.Error("service.go::AddLocalService(handler:%#v\tError:handler Cannot be nil", handler)
	}
	if LocalServiceStore == nil {
		LocalServiceStore = LocalServiceMap{}
	}
	LocalServiceStore[service] = handler

	return AddService(service, LocalHost.TCPUrl)
}

//AddService add a service to Service Map
func AddService(service Service, tcp string) error {
	if len(service.ServiceName) == 0 {
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
