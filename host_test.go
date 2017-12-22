package gtmesh

import (
	"reflect"
	"testing"
	"time"

	"github.com/rsms/gotalk"
)

var (
	testNode = GetNode("TestNode")
)

func TestNode_SetLocalHost(t *testing.T) {
	type args struct {
		tcp string
		ws  string
	}
	tests := []struct {
		name    string
		node    *Node
		args    args
		wantErr bool
	}{
		{"empty tcp,ws", testNode, args{"", ""}, true},
		{"bad TCP and empty WS", testNode, args{"badtcp://localhost:7070", ""}, true},
		{"valid TCP and empty WS", testNode, args{"tcp://localhost:7070", ""}, false},
		{"valid TCP and bad WS", testNode, args{"tcp://localhost:7070", "badws"}, true},
		{"valid TCP and WS but same port", testNode, args{"tcp://localhost:7070", "ws://localhost:7070"}, true},
		{"valid TCP and WS & different ports", testNode, args{"tcp://localhost:7070", "ws://localhost:7071"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.node.SetLocalHost(tt.args.tcp, tt.args.ws)

			if (err != nil) != tt.wantErr {
				t.Errorf("Node.SetLocalHost() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if tt.args.tcp == tt.args.ws {
					t.Errorf("SetLocalHost(tcp=%s,ws=%s)\tWanted:tcp!=ws\tGot:%#v,%#v", tt.args.tcp, tt.args.ws, tt.args.tcp, tt.args.ws)
				}
			}
		})
	}
}

func TestNode_StartTCPServer(t *testing.T) {
	tests := []struct {
		name    string
		node    *Node
		wantErr bool
	}{
		{"-", testNode, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.node.StartTCPServer(); (err != nil) != tt.wantErr {
				t.Errorf("Node.StartTCPServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNode_StartWSServer(t *testing.T) {
	tests := []struct {
		name    string
		node    *Node
		wantErr bool
	}{
		{"-", testNode, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.node.StartWSServer(); (err != nil) != tt.wantErr {
				t.Errorf("Node.StartWSServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_echoHandler(t *testing.T) {
	type args struct {
		s       *gotalk.Sock
		op      string
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := echoHandler(tt.args.s, tt.args.op, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("echoHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("echoHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addrHandler(t *testing.T) {
	type args struct {
		s       *gotalk.Sock
		op      string
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addrHandler(tt.args.s, tt.args.op, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("addrHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addrHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastState(t *testing.T) {
	type args struct {
		h1 ServiceData
		h2 ServiceData
	}

	t0 := time.Now()
	tb := t0.Add(1 * time.Second)
	ta := tb.Add(1 * time.Second)

	tests := []struct {
		name   string
		args   args
		wantSd ServiceData
	}{
		{"h1 created after h2, no deletes", args{h1: ServiceData{Created: ta}, h2: ServiceData{Created: tb}}, ServiceData{Created: ta}},
		{"h2 created after h1, no deletes", args{h1: ServiceData{Created: tb}, h2: ServiceData{Created: ta}}, ServiceData{Created: ta}},
		{"h1 deleted after h2", args{h1: ServiceData{Created: t0, Deleted: ta}, h2: ServiceData{Created: t0, Deleted: tb}}, ServiceData{Created: t0, Deleted: ta}},
		{"h2 deleted after h1", args{h1: ServiceData{Created: t0, Deleted: tb}, h2: ServiceData{Created: t0, Deleted: ta}}, ServiceData{Created: t0, Deleted: ta}},
		{"h1 created after h2, with past delete", args{h1: ServiceData{Created: ta, Deleted: t0}, h2: ServiceData{Created: tb, Deleted: t0}}, ServiceData{Created: ta, Deleted: t0}},
		{"h2 created after h1, with past delete", args{h1: ServiceData{Created: tb, Deleted: t0}, h2: ServiceData{Created: ta, Deleted: t0}}, ServiceData{Created: ta, Deleted: t0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSd := lastState(tt.args.h1, tt.args.h2); !reflect.DeepEqual(gotSd, tt.wantSd) {
				t.Errorf("lastState() = %v, want %v", gotSd, tt.wantSd)
			}
		})
	}
}
