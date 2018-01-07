package gtmesh

import (
	"reflect"
	"testing"
	"time"

	"github.com/rsms/gotalk"
)

var (
	testNode  = GetNode("TestNode")
	testNodeB = GetNode("TestNodeB")
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

func TestNode_AddPeer(t *testing.T) {

	serverA := GetNode("ServerA")
	serverA.SetLocalHost("tcp://localhost:7080", "")
	serverA.StartTCPServer()
	serverB := GetNode("ServerB")
	serverB.SetLocalHost("tcp://localhost:7081", "")
	serverB.StartTCPServer()

	type args struct {
		peerURLString string
	}
	tests := []struct {
		name    string
		node    *Node
		args    args
		wantErr bool
	}{
		{"empty peerURLString", serverA, args{""}, true},
		{"bad peerURLString", serverA, args{"tcp://localhost:7072"}, true},
		{"valid peerURLString", serverA, args{serverB.LocalHost.TCPUrl}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.node.AddPeer(tt.args.peerURLString)

			if (err != nil) != tt.wantErr {
				t.Errorf("Node.AddPeer() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestNode_connectSync(t *testing.T) {
	type fields struct {
		Name              string
		LocalHost         Host
		tcpServer         *gotalk.Server
		wsServer          *gotalk.WebSocketServer
		LocalServiceStore LocalServiceMap
		ServiceStore      ServiceMap
		ServiceHandlers   *gotalk.Handlers
		ConvergenceWindow time.Duration
	}
	type args struct {
		peerURLString string
		syncFrame     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Name:              tt.fields.Name,
				LocalHost:         tt.fields.LocalHost,
				tcpServer:         tt.fields.tcpServer,
				wsServer:          tt.fields.wsServer,
				LocalServiceStore: tt.fields.LocalServiceStore,
				ServiceStore:      tt.fields.ServiceStore,
				ServiceHandlers:   tt.fields.ServiceHandlers,
				ConvergenceWindow: tt.fields.ConvergenceWindow,
			}
			if err := node.connectSync(tt.args.peerURLString, tt.args.syncFrame); (err != nil) != tt.wantErr {
				t.Errorf("Node.connectSync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_syncMapHandler(t *testing.T) {
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
			got, err := syncMapHandler(tt.args.s, tt.args.op, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("syncMapHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("syncMapHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_syncMapService(t *testing.T) {
	type args struct {
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
			got, err := syncMapService(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("syncMapService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("syncMapService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_finalState(t *testing.T) {
	type args struct {
		sd ServiceData
	}
	tests := []struct {
		name    string
		args    args
		wantFss FinalServiceState
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFss := finalState(tt.args.sd); !reflect.DeepEqual(gotFss, tt.wantFss) {
				t.Errorf("finalState() = %v, want %v", gotFss, tt.wantFss)
			}
		})
	}
}

func Test_timeDiff(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeDiff(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("timeDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prettyJSON(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettyJSON(tt.args.i); got != tt.want {
				t.Errorf("prettyJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
