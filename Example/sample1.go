//Main Package
package main

import (
	"github.com/julianfrank/console"
	"github.com/julianfrank/gtmesh"
)

func main() {
	console.Log("Starting Server A")
	serverA := gtmesh.GetNode("ServerA")

	console.Log("%#v", serverA)
}
