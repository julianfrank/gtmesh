package console

import (
	"strings"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	type args struct {
		pattern string
		vars    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"stringtest", args{"%s", []interface{}{"stringtest"}}, "stringtest"},
		{"stringstest", args{"%s %s", []interface{}{"string1", "string2"}}, "string1 string2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Log(tt.args.pattern, tt.args.vars...); got != tt.want {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		pattern string
		vars    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"stringtest", args{"%s", []interface{}{"stringtest"}}, true},
		{"stringstest", args{"%s %s", []interface{}{"string1", "string2"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(tt.args.pattern, tt.args.vars...); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimed(t *testing.T) {
	type args struct {
		startTime time.Time
		pattern   string
		vars      []interface{}
	}

	var st = time.Now()

	tests := []struct {
		name string
		args args
		want string
	}{
		{"str1", args{st, "%s", []interface{}{"str1"}}, "str1"},
		{"str2", args{st, "%s-%s", []interface{}{"str1", "str2"}}, "str1-str2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Timed(tt.args.startTime, tt.args.pattern, tt.args.vars...); strings.Split(got, "\t")[0] != tt.want {
				t.Errorf("Timed() = %v, want %v", got, tt.want)
			}
		})
	}
}
