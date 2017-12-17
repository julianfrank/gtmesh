package gtmesh

import (
	"github.com/julianfrank/console"
)

//BufferRequest Request for a Buffer/[]byte based service
func BufferRequest(serviceName string, payLoad []byte) ([]byte, error) {
	//Basic Sanity Check
	if len(serviceName) == 0 {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName cannot be empty", serviceName, string(payLoad))
	}
	//Find where the service is located
	if ServiceStore[serviceName] == nil {
		return nil, console.Error("BufferRequest(serviceName:%s,payLoad:%s) Error: serviceName cannot be empty", serviceName, string(payLoad))
	}
	return nil, nil
}
