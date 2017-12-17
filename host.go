package gtmesh

import (
	"net/http"
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
		LocalHost = Host{TCPUrl: temptcp.Host, WSUrl: tempws.Host}
	} else {
		LocalHost = Host{TCPUrl: temptcp.Host, WSUrl: ""}
	}

	return nil
}

//StartServers Start the tCP and WebSocket Servers
func StartServers() error {
	console.Log("host.go::StartServers()")

	err := startTCPServer()
	if err != nil {
		return console.Error("StartServers() Error:%s", err.Error())
	}

	if LocalHost.WSUrl != "" {
		err = startWSServer()
		if err != nil {
			return console.Error("StartServers() Error:%s", err.Error())
		}
	}

	return nil
}

//startTCPServer Start the tcp server
func startTCPServer() error {
	console.Log("startTCPServer() for LocalHost.TCPURL:%s", LocalHost.TCPUrl)

	tcpServer, err := gotalk.Listen("tcp", LocalHost.TCPUrl)
	if err != nil {
		return console.Error("startTCPServer() Error:%s", err.Error())
	}
	tcpServer.Handlers = ServiceHandlers

	go tcpServer.Accept()
	return nil
}

//startTCPServer Start the tcp server
func startWSServer() error {
	console.Log("host.go::startWSServer() for LocalHost.WSURL:%s", LocalHost.WSUrl)

	if LocalHost.WSUrl != "" {
		wsServer = gotalk.WebSocketHandler()
		wsServer.Handlers = ServiceHandlers
		//wsServer.OnAccept = onAccept
		http.Handle("/gotalk/", wsServer)
		//http.Handle("/", http.FileServer(http.Dir(".")))

		//[TODO]This is NOT the right way to do this...Need to rework!
		go func() {
			err := http.ListenAndServe(LocalHost.WSUrl, nil)
			if err != nil {
				console.Error("startWSServer() with LocalHost.WSUrl=%s has Error:%s", LocalHost.WSUrl, err.Error())
			}
		}()

	}
	return nil
}
