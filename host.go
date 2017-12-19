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

//SetLocalHost Setup Local Host Details. must perform this before any other operation
func (node *Node) SetLocalHost(tcp string, ws string) error {
	console.Log("host.go::Node.SetLocalHost(tcp:%s,ws:%s)", tcp, ws)
	//Sanitary Check
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
	//Check ws only if not null...WebSocket is optional
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
		node.LocalHost = Host{TCPUrl: temptcp.Host, WSUrl: tempws.Host}
	} else {
		node.LocalHost = Host{TCPUrl: temptcp.Host, WSUrl: ""}
	}

	return nil
}

//StartTCPServer Start the tcp server
func (node *Node) StartTCPServer() error {
	console.Log("StartTCPServer() for LocalHost.TCPURL:%s", node.LocalHost.TCPUrl)
	//Make TCPServer Listen on the TCPURL
	tcpServer, err := gotalk.Listen("tcp", node.LocalHost.TCPUrl)
	if err != nil {
		return console.Error("StartTCPServer() Error:%s", err.Error())
	}
	//Set echo as default service in all TCPServers
	node.AddLocalService("echo", echoHandler)
	node.AddLocalService("addr", addrHandler)
	//Attach Handlers to TCPServer's Handlers
	tcpServer.Handlers = node.ServiceHandlers
	//Start Accepting Connections
	node.tcpServer = tcpServer
	go node.tcpServer.Accept()
	return nil
}

//StartWSServer Start the tcp server
//[NOT READY]
func (node *Node) StartWSServer() error {
	console.Log("host.go::StartWSServer() for LocalHost.WSURL:%s", node.LocalHost.WSUrl)
	//Start only if ws url is not nil
	if node.LocalHost.WSUrl != "" {
		node.wsServer = gotalk.WebSocketHandler()
		node.wsServer.Handlers = node.ServiceHandlers
		//wsServer.OnAccept = onAccept
		http.Handle("/gotalk/", node.wsServer)
		//http.Handle("/", http.FileServer(http.Dir(".")))

		//[TODO]This is NOT the right way to do this...Need to rework!
		go func() {
			err := http.ListenAndServe(node.LocalHost.WSUrl, nil)
			if err != nil {
				console.Error("StartWSServer() with node.LocalHost.WSUrl=%s has Error:%s", node.LocalHost.WSUrl, err.Error())
			}
		}()

	}
	return nil
}

//echoHandler Default Handler in All Servers to perform echo
func echoHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	return payload, nil
}

//addrHandler Default Handler in All Servers to sendback address
func addrHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	return []byte(s.Addr()), nil
}
