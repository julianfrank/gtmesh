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
// version : 22dec2017
type Host struct {
	TCPUrl string `json:"tcp_url"`
	WSUrl  string `json:"ws_url,omitempty"`
}

//SetLocalHost Setup Local Host Details. must perform this before any other operation
// version : 22dec2017
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
// version : 22dec2017
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

// syncMap Structure used to hold the Synchronization Frames
// version : 22dec2017
type syncMap struct {
	SourceHostName string     `json:"source_host_name"`
	ServiceMap     ServiceMap `json:"service_map"`
}

//AddPeer Add a New Peer to this Node
// version : 22dec2017
func (node *Node) AddPeer(peerURLString string) error {
	console.Log("host.go::Node.addPeer(peerURLString: %s)", peerURLString)

	// Build frame to send to Peer
	frame := syncMap{
		SourceHostName: node.Name,
		ServiceMap:     node.ServiceStore,
	}
	syncFrame, err := json.Marshal(frame)
	if err != nil {
		return console.Error("Node.AddPeer(peerURLString: %s)\tsyncFrame,err:=json.Marshal(frame)\tError: %s", peerURLString, err.Error())
	}
	//console.Log("syncFrame:%s", string(syncFrame))

	// Connect to Peer
	s, err := gotalk.Connect("tcp", peerURLString)
	if err != nil {
		return console.Error("Node.AddPeer(peerURLString: %s)\tError: %s", peerURLString, err.Error())
	}

	//Attach Standard Utilities to Connection
	s.UserData = node.Name
	//[TBD]

	// Invoke SyncMap with local map as seed
	res, err := s.BufferRequest("sys.syncmap", syncFrame)
	if err != nil {
		return console.Error("s.BufferRequest(`sys.syncmap`, syncFrame:%+v)\nError: %s", syncFrame, err.Error())
	}
	if res != nil {
		console.Log("Sync Happened!\tres:%s", string(res))
		//Retreive the syncMap from the payload
		var remoteMap syncMap
		err := json.Unmarshal(res, &remoteMap)
		if err != nil {
			return console.Error("Node.AddPeer::json.Unmarshal(payload: %s ...\tError:%s", string(res), err.Error())
		}

		//Copy synced serviceMap to Local Service Map
		node.ServiceStore = remoteMap.ServiceMap

	} else {
		console.Log("%s Sync Returned, Result: Already in Sync\tres:%s", node.Name, string(res))
	}

	return nil
}

//syncMapHandler Default Handler in All Servers to sendback address
// version : 22dec2017
func syncMapHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::syncMapHandler(s.Addr(): %s,\top: %s,\tpayload: %s)", s.Addr(), op, string(payload))

	//Retreive the syncMap from the payload
	var remoteMap syncMap
	err := json.Unmarshal(payload, &remoteMap)
	if err != nil {
		return nil, console.Error("json.Unmarshal(payload: %s ...\tError:%s", string(payload), err.Error())
	}
	//console.Log("remoteMap:\t%+v", remoteMap)

	//Sync Up with local ServiceMaps
	localSS := localNode.ServiceStore
	remoteSS := remoteMap.ServiceMap
	//console.Log("\nlocalSS:%+v\nremoteSS:%+v", localSS, remoteSS)

	baseFrame := remoteMap
	baseFrame.SourceHostName = localNode.Name
	//Sync by parsing the remoteSS
	//console.Log("Performing Remote Map Scan")
	for svc, hmap := range remoteSS {
		//console.Log("svc:%s\thm:%+v", svc, hmap)
		if localSS[svc] == nil {
			//Local Map does not have this service. Skip
			//console.Log("local Sync NOT Needed for service:%s", svc)
		} else {
			for h, d := range hmap {
				console.Log("svc:%s\thost:%s\tdetails:%+v", svc, h, d)
				if _, ok := localSS[svc][h]; ok {
					//console.Log("local Sync Needed for %+v", localSS[svc][h])

					latestState := lastState(localSS[svc][h], d)
					latestState.Source = localNode.LocalHost.TCPUrl
					baseFrame.ServiceMap[svc][h] = latestState

				} else {
					//This host does not exist in local Map. Skip
					//console.Log("local Sync NOT Needed for host:%s", h)
				}
			}
		}
	}
	//console.Log("Performing Local Map Scan now...")
	for svc, hmap := range localSS {
		if remoteSS[svc] == nil {
			//Remote does not have this service...Copy
			baseFrame.ServiceMap[svc] = hmap
			//console.Log("service %s being copied to Frame", svc)
		} else {
			for h, d := range hmap {
				//console.Log("svc:%s\thost:%s\tdetails:%+v", svc, h, d)
				if _, ok := remoteSS[svc][h]; ok {
					//Local Map has Detail as the remote Host . Sync

					latestState := lastState(localSS[svc][h], d)
					latestState.Source = localNode.LocalHost.TCPUrl
					baseFrame.ServiceMap[svc][h] = latestState

					//console.Log(" Sync Needed for Host:%s", h)
				} else {
					//Remote does not have detail about the local host map...Copy
					baseFrame.ServiceMap[svc][h] = localSS[svc][h]
					//console.Log("Frame being Updated with Host:%s", h)
				}
			}
		}
	}
	//console.Log("\n\nbaseFrame%s", prettyJSON(baseFrame))

	//Prepare List of Host to Propagate Sync. Exclude Sender.Also Do not perform if sync Date of sender is older
	//[TBD]

	//Initiate Sync with identified Hosts as a separate GoRoutine
	//[TBD]

	// Copy baseFrame to local ServiceMap
	localNode.ServiceStore = baseFrame.ServiceMap

	//Respond with Updated Map if new else just send nil
	syncFrame, err := json.Marshal(baseFrame)
	if err != nil {
		return nil, console.Error("host.go::syncMapHandler(s.Addr(): %s,\top: %s,\tpayload: %s) syncFrame, err := json.Marshal(baseFrame) Error: %s", s.Addr(), op, string(payload), err.Error())
	}
	return syncFrame, nil
}

// lastState Returns the Latest Object of the provides ServiceData Objects
// Version : 22Dec2017
func lastState(h1, h2 ServiceData) (sd ServiceData) {
	h1t := h1.Created
	if h1.Created.Before(h1.Deleted) {
		h1t = h1.Deleted
	}
	h2t := h2.Created
	if h2.Created.Before(h2.Deleted) {
		h2t = h2.Deleted
	}
	if h1t.After(h2t) {
		return h1
	}
	return h2
}

// timeDiff Returns the Difference between t1 and t1 as a time.Duration and t1>t2 as a bool.
// if t1==t2 then nil
// version : 21Dec2017
func timeDiff(t1, t2 time.Time) time.Duration {
	switch {
	case t1.Equal(t2):
		return 0
	case t1.After(t2):
		return t1.Sub(t2)
	default:
		return t2.Sub(t1)
	}
}

//echoHandler Default Handler in All Servers to perform echo
// version : 22dec2017
func echoHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::echoHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))
	return payload, nil
}

//addrHandler Default Handler in All Servers to sendback address
// version : 22dec2017
func addrHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::addrHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))
	return []byte(s.Addr()), nil
}

//prettyJSON Return string Formated as a Pretty JSON for the givent interface{}
// version : 22dec2017
func prettyJSON(i interface{}) string {
	b, _ := json.MarshalIndent(i, "|", " ")
	return string(b)
}

/* Future Stuff - Dont Bother Right Now





















































 */

//StartWSServer Start the tcp server
//[WARNING : NOT READY]
// version : 22dec2017
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
