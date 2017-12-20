package gtmesh

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

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
	console.Log("host.go::Node.SetLocalHost(tcp: %s,ws: %s)", tcp, ws)
	//Sanitary Check
	if len(tcp) == 0 {
		return console.Error("TCP Cannot be empty")
	}
	temptcp, err := url.ParseRequestURI(tcp)
	if err != nil {
		return console.Error("SetLocalHost(tcp: %s,ws: %s) Error: tcp has Bad Format\t%s", tcp, ws, err.Error())
	}
	if temptcp.Scheme != "tcp" || temptcp.Host == "" || temptcp.Port() == "" {
		return console.Error("SetLocalHost(tcp: %s,ws: %s) Error: tcp has Bad Format", tcp, ws)
	}
	//Check ws only if not null...WebSocket is optional
	if ws != "" {
		tempws, err := url.ParseRequestURI(ws)
		if err != nil {
			return console.Error("SetLocalHost(tcp: %s,ws: %s) Error: ws has Bad Format\t%s", tcp, ws, err.Error())
		}
		if tempws.Scheme != "ws" || tempws.Host == "" || tempws.Port() == "" {
			return console.Error("SetLocalHost(tcp: %s,ws: %s) Error: ws has Bad Format", tcp, ws)
		}
		if temptcp.Port() == tempws.Port() {
			return console.Error("SetLocalHost(tcp: %s,ws: %s) Error: Port cannot be the same", tcp, ws)
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
	node.AddLocalService("sys.echo", echoHandler)
	node.AddLocalService("sys.addr", addrHandler)
	node.AddLocalService("sys.syncmap", syncMapHandler)

	//Attach Handlers to TCPServer's Handlers
	tcpServer.Handlers = node.ServiceHandlers

	//Start Accepting Connections
	node.tcpServer = tcpServer

	go node.tcpServer.Accept()

	return nil
}

//echoHandler Default Handler in All Servers to perform echo
func echoHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::echoHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))
	return payload, nil
}

//addrHandler Default Handler in All Servers to sendback address
func addrHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::addrHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))
	return []byte(s.Addr()), nil
}

// syncMap Structure used to hold the Synchronization Frames
type syncMap struct {
	SourceHostName string     `json:"source_host_name"`
	LastUpdate     time.Time  `json:"last_update"`
	Map            ServiceMap `json:"map"`
}

//AddPeer Add a New Peer to this Node
func (node *Node) AddPeer(peerURLString string) error {
	console.Log("host.go::Node.addPeer(peerURLString: %s)", peerURLString)

	// Build frame to send to Peer
	frame := syncMap{
		SourceHostName: node.Name,
		LastUpdate:     node.lastServiceUpdateTime,
		Map:            node.ServiceStore,
	}
	syncFrame, err := json.Marshal(frame)
	if err != nil {
		return console.Error("Node.AddPeer(peerURLString: %s)\tsyncFrame,err:=json.Marshal(frame)\tError: %s", peerURLString, err.Error())
	}
	console.Log("syncFrame:%s", string(syncFrame))

	// Connect to Peer
	s, err := gotalk.Connect("tcp", peerURLString)
	if err != nil {
		return console.Error("Node.AddPeer(peerURLString: %s)\tError: %s", peerURLString, err.Error())
	}
	console.Log("\ns:%+v\nframe:%+v", s, frame)

	//Attach Standard Utilities to Connection
	s.UserData = node.Name
	//[TBD]

	// Invoke SyncMap with local map as seed
	res, err := s.BufferRequest("sys.syncmap", syncFrame)
	if err != nil {
		return console.Error("s.BufferRequest(`sys.syncmap`, syncFrame:%+v)\tError: %s", syncFrame, err.Error())
	}
	console.Log("res:%s", string(res))

	return nil
}

//syncMapHandler Default Handler in All Servers to sendback address
func syncMapHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::syncMapHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))

	//Retreive the syncMap from the payload
	var remoteMap syncMap
	err := json.Unmarshal(payload, &remoteMap)
	if err != nil {
		console.Log("json.Unmarshal(payload: %s ...\tError:%s", string(payload), err.Error())
	}
	console.Log("remoteMap:\t%+v", remoteMap)

	//Sync Up with local ServiceMaps
	localSS := localNode.ServiceStore
	remoteSS := remoteMap.Map
	console.Log("\n\nlocalSS:%+v\nremoteSS:%+v\n", localSS, remoteSS)
	localST := localNode.lastServiceUpdateTime
	remoteST := remoteMap.LastUpdate
	console.Log("\n\nlocalST:%+v\nremoteST:%+v\n", localST, remoteST)

	switch {
	case localST.Equal(remoteST):
		console.Log("LocalST == remoteST")
		return nil, nil

	case localST.After(remoteST):
		console.Log("LocalST > remoteST")

	case localST.Before(remoteST):
		console.Log("LocalST < remoteST")
	}

	//Prepare List of Host to Propagate Sync. Exclude Sender.Also Do not perform if sync Date of sender is older

	//Initiate Sync with identified Hosts as a separate GoRoutine

	//Respond with Updated Map if new else just send nil

	return []byte("syncMap"), nil
}

/* Future Stuff - Dont Bother Right Now





















































 */

//StartWSServer Start the tcp server
//[WARNING : NOT READY]
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
