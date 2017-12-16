package gtmesh

import (
	"fmt"

	"github.com/rsms/gotalk"
)

//Check Cool
func Check() {
	fmt.Printf("%#v", gotalk.NoLimits)
}

//Add adds
func Add(x int, y int) int {
	return x + y
}
