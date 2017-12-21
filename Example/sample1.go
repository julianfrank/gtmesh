//Main Package
package main

import (
	"github.com/julianfrank/console"
	"github.com/julianfrank/gtmesh"
	"github.com/rsms/gotalk"
)

func main() {
	console.LogMode = true
	console.Log("Starting Server A")
	serverA := gtmesh.GetNode("ServerA")
	serverA.SetLocalHost("tcp://localhost:7070", "")
	serverA.AddLocalService("h1s", h1svcHandler)
	serverA.StartTCPServer()
	//serverA.BufferRequest("sys.echo", []byte("Testing Echo from Server A"))
	//console.Log("serverA %+v", serverA)
/*
	console.Log("Starting Server B")
	serverB := gtmesh.GetNode("ServerB")
	serverB.SetLocalHost("tcp://localhost:7071", "")
	serverB.AddLocalService("h2s", h2svcHandler)
	serverB.StartTCPServer()
	//serverB.BufferRequest("sys.addr", []byte("Testing Addr from Server B"))
	//console.Log("serverB %+v", serverB)

	console.LogMode = true

	console.Log("Going to serverA.AddPeer(serverB: %s )", serverB.LocalHost.TCPUrl)
	serverA.AddPeer(serverB.LocalHost.TCPUrl)
	console.Log("\nserverA %+v\nserverB %+v", serverA, serverB)*/
}

func h1svcHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::h1svcHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))
	return []byte("h1svcHandler"), nil
}

func h2svcHandler(s *gotalk.Sock, op string, payload []byte) ([]byte, error) {
	console.Log("host.go::h2svcHandler(s.Addr(): %s,op: %s,payload: %s)", s.Addr(), op, string(payload))
	return []byte("h2svcHandler"), nil
}
