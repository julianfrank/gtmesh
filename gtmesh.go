package gtmesh

import (
	"time"

	"github.com/julianfrank/console"
)

//Node Base Object holding the Mesh for each Logical Server
type Node struct {
	Name string
}

//GetNode Get a Fresh Instance of the GTMesh Node
func GetNode(nodeName string) Node {
	console.Log("gtmesh.go::GetNode(nodeName:%s)", nodeName)
	if nodeName == "" {
		return Node{Name: time.Now().Format("06JanMon150405")}
	}
	return Node{Name: nodeName}
}
