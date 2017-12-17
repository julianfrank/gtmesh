package gtmesh

import (
	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//BufferRequest Request for a Buffer/[]byte based service
func BufferRequest(serviceName string, payLoad []byte) ([]byte, error) {
	//Basic Sanity Check
	if len(serviceName) == 0 {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName cannot be empty", serviceName, string(payLoad))
	}
	//Find where the service is located
	if ServiceStore[serviceName] == nil {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName Not Registered yet! %#v", serviceName, string(payLoad), ServiceStore)
	}
	//Iterate through each host till service result is obtained
	for _, host := range ServiceStore[serviceName] {
		s, err := gotalk.Connect("tcp", host)
		if err != nil {
			console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with %s and returned error %s", serviceName, payLoad, host, err.Error())
		} else {
			return s.BufferRequest(serviceName, payLoad)
		}
	}
	return nil, console.Error("gtmesh.go::BufferRequest(serviceName:%s,payload:%s) unable to connect with any hosts [%#v]", serviceName, payLoad, ServiceStore[serviceName])
}
