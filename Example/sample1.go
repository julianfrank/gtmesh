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
	console.Log("%#v", serverA)
}
