package gtmesh

import (
	"net/url"

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
	temptcp, err := url.ParseRequestURI(tcp)
	if err != nil {
		return console.Error("SetLocalHost(tcp:%s,ws:%s) Error: tcp has Bad Format\t%s", tcp, ws, err.Error())
	}
	if temptcp.Scheme != "tcp" || temptcp.Host == "" || temptcp.Port() == "" {
		return console.Error("SetLocalHost(tcp:%s,ws:%s) Error: tcp has Bad Format", tcp, ws)
	}
	if ws != "" {
		tempws, err := url.ParseRequestURI(ws)
		if err != nil {
			return console.Error("SetLocalHost(tcp:%s,ws:%s) Error: ws has Bad Format\t%s", tcp, ws, err.Error())
		}
		if tempws.Scheme != "ws" || tempws.Host == "" || tempws.Port() == "" {
			return console.Error("SetLocalHost(tcp:%s,ws:%s) Error: ws has Bad Format", tcp, ws)
		}
		if temptcp.Port() == tempws.Port() {
			return console.Error("SetLocalHost(tcp:%s,ws:%s) Error: Port cannot be the same", tcp, ws)
		}
	}

	LocalHost = Host{TCPUrl: tcp, WSUrl: ws}
	return nil
}

//StartServers Start the tCP and WebSocket Servers
func StartServers() error {
	console.Log("host.go::StartServers()")
	err := startTCPServer()
	if err != nil {
		return err
	}
	err = startWSServer()
	if err != nil {
		return err
	}
	return nil
}

//startTCPServer Start the tcp server
func startTCPServer() error {
	console.Log("host.go::startTCPServer(tcpURL:%s)")
	return nil
}

//startTCPServer Start the tcp server
func startWSServer() error {
	return nil
}
