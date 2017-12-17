package gtmesh

import (
	"github.com/rsms/gotalk"
)

//opMap Map of Operations
type opMap map[string]gotalk.BufferReqHandler

var (
	//OpMap Map of Buffer Manager Operations
	OpMap opMap
	//LocalHost The TCP/WS Name of the Local Host
	LocalHost Host
)

func init() {
	OpMap = make(map[string]gotalk.BufferReqHandler)
}
