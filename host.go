package gtmesh

import (
	"fmt"
)

//SetLocalHost Setup Local Host Details. must perform this before any other operation
func SetLocalHost(tcp string, ws string) error {
	if len(tcp) == 0 {
		return fmt.Errorf("TCP Cannot be empty")
	}
	LocalHost = Host{TCPUrl: tcp, WSUrl: ws}
	return nil
}
