package gtmesh

import (
	"github.com/rsms/gotalk"
)

//Host host struct
type Host struct {
	TCPUrl string `json:"tcp_url"`
	WSUrl  string `json:"ws_url,omitempty"`
}

//Service service struct
type Service struct {
	ServiceName string `json:"service_name"`
	Local       bool   `json:"local"`
}

//HostMap map of host to services
type HostMap map[Host][]Service

//ServiceMap map of services to hosts
type ServiceMap map[Service][]Host

//opMap Map of Operations
type opMap map[string]gotalk.BufferReqHandler

var (
	//OpMap Map of Buffer Manager Operations
	OpMap opMap
	//LocalHost The TCP/WS Name of the Local Host
	LocalHost Host
)

func init() {
	OpMap = make(map[string]gotalk.BufferReqHandler)
}
