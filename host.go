package gtmesh

import (
	"github.com/julianfrank/console"
)

//Host host struct
type Host struct {
	TCPUrl string `json:"tcp_url"`
	WSUrl  string `json:"ws_url,omitempty"`
}

//SetLocalHost Setup Local Host Details. must perform this before any other operation
func SetLocalHost(tcp string, ws string) error {
	console.Log("host.go::SetLocalHost(tcp:%s,ws:%s)", tcp, ws)
	if len(tcp) == 0 {
		return console.Error("TCP Cannot be empty")
	}
	LocalHost = Host{TCPUrl: tcp, WSUrl: ws}
	return nil
}
