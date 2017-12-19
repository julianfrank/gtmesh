package gtmesh

import (
	"reflect"
	"testing"
)

func TestGetNode(t *testing.T) {
	type args struct {
		nodeName string
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		{"empty name", args{nodeName: ""}, Node{}},
		{"with name", args{nodeName: "testName"}, Node{Name: "testName"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNode(tt.args.nodeName); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
