package gtmesh

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Check()
		})
	}
}
