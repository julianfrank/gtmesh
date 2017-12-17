package gtmesh

import (
	"fmt"

	"github.com/julianfrank/console"
)

//SetLocalHost Setup Local Host Details. must perform this before any other operation
func SetLocalHost(tcp string, ws string) error {
	console.Log("host.go::SetLocalHost(tcp:%s,ws:%s)", tcp, ws)
	if len(tcp) == 0 {
		return fmt.Errorf("TCP Cannot be empty")
	}
	LocalHost = Host{TCPUrl: tcp, WSUrl: ws}
	return nil
}
