package gtmesh

import (
	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//Host host struct
type Host struct {
	TCPUrl string `json:"tcp_url"`
	WSUrl  string `json:"ws_url,omitempty"`
}

var (
	//LocalHost The TCP/WS Name of the Local Host
	LocalHost Host
	tcpServer *gotalk.Server
	wsServer  *gotalk.WebSocketServer
)

//SetLocalHost Setup Local Host Details. must perform this before any other operation
func SetLocalHost(tcp string, ws string) error {
	console.Log("host.go::SetLocalHost(tcp:%s,ws:%s)", tcp, ws)
	if len(tcp) == 0 {
		return console.Error("TCP Cannot be empty")
	}
	LocalHost = Host{TCPUrl: tcp, WSUrl: ws}
	return nil
}

//StartServers Start the tCP and WebSocket Servers
func StartServers(host Host) error {
	console.Log("host.go::StartServers()")
	if host.TCPUrl == "" {
		return console.Error("StartServers() Error: host.TCPUrl is empty")
	}
	err := startTCPServer(host.TCPUrl)
	if err != nil {
		return err
	}
	err = startWSServer(host.WSUrl)
	if err != nil {
		return err
	}
	return nil
}

//startTCPServer Start the tcp server
func startTCPServer(tcpURL string) error {
	return nil
}

//startTCPServer Start the tcp server
func startWSServer(wsURL string) error {
	return nil
}
