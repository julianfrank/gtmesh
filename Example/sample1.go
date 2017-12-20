//Main Package
package main

import (
	"github.com/julianfrank/console"
	"github.com/julianfrank/gtmesh"
)

func main() {
	console.LogMode = false
	console.Log("Starting Server A")
	serverA := gtmesh.GetNode("ServerA")
	serverA.SetLocalHost("tcp://localhost:7070", "")
	serverA.StartTCPServer()
	//serverA.BufferRequest("sys.echo", []byte("Testing Echo from Server A"))
	//console.Log("serverA %+v", serverA)

	console.Log("Starting Server B")
	serverB := gtmesh.GetNode("ServerB")
	serverB.SetLocalHost("tcp://localhost:7071", "")
	serverB.StartTCPServer()
	//serverB.BufferRequest("sys.addr", []byte("Testing Addr from Server B"))
	//console.Log("serverB %+v", serverB)

	console.LogMode = true

	console.Log("Going to serverA.AddPeer(serverB: %s )", serverB.LocalHost.TCPUrl)
	serverA.AddPeer(serverB.LocalHost.TCPUrl)
	console.Log("\nserverA %+v\nserverB %+v", serverA, serverB)
}
