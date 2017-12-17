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
type ServiceMap map[Service][]Host

//LocalServiceMap map of Local Services
type LocalServiceMap map[Service]gotalk.BufferReqHandler

var (
	//LocalServiceStore Hosts the map of all local Services
	LocalServiceStore LocalServiceMap
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
	return nil
}
