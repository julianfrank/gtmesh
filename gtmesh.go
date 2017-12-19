package gtmesh

import (
	"time"

	"github.com/julianfrank/console"
	"github.com/rsms/gotalk"
)

//Node Base Object holding the Mesh for each Logical Server
type Node struct {
	Name string
	//LocalHost The TCP/WS Name of the Local Host
	LocalHost Host
	//tcpServer This Points to the local TCP Server
	tcpServer *gotalk.Server
	//wsServer This points to the local WebSocket Server
	wsServer *gotalk.WebSocketServer

	//LocalServiceStore Hosts the map of all local Services
	LocalServiceStore LocalServiceMap
	//ServiceStore Hosts the mapping of all services in the mesh mapped to their hosts
	ServiceStore ServiceMap
	//ServiceHandlers Bank of Handlers used by GoTalk
	ServiceHandlers *gotalk.Handlers
	//lastServiceUpdateTime
	lastServiceUpdateTime time.Time
}

//GetNode Get a Fresh Instance of the GTMesh Node
func GetNode(nodeName string) Node {
	console.Log("gtmesh.go::GetNode(nodeName:%s)", nodeName)
	if nodeName == "" {
		return Node{Name: time.Now().Format("06JanMon150405")}
	}
	return Node{Name: nodeName}
}
