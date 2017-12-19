//Main Package
package main

import (
	"github.com/julianfrank/console"
	"github.com/julianfrank/gtmesh"
)

func main() {
	console.Log("Starting Server A")
	serverA := gtmesh.GetNode("ServerA")
	serverA.SetLocalHost("tcp://localhost:7070", "")
	serverA.StartTCPServer()
	console.Log("serverA %#v", serverA)

	console.Log("Starting Server B")
	serverB := gtmesh.GetNode("ServerB")
	serverB.SetLocalHost("tcp://localhost:7071", "")
	serverB.StartTCPServer()
	console.Log("serverB %#v", serverB)

	console.Log("Going to serverA.AddPeer(serverB: %s )", serverB.LocalHost.TCPUrl)
	serverA.AddPeer(serverB.LocalHost.TCPUrl)
	console.Log("serverA %#v", serverA)
}
