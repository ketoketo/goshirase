package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{"case 1", args{"testconfig/config.json"}, Config{ConsumerKey: "param1", ConsumerSecret: "param2", AccessToken: "param3", AccessSecret: "param4"}, false},
		{"case 2", args{"testconfig/config2.json"}, Config{ConsumerKey: "AAA", ConsumerSecret: "BBB", AccessToken: "CCC", AccessSecret: "DDD"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stdReader(t *testing.T) {
	type args struct {
		stdin   io.Reader
		message string
		param   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case 1", args{stdin: bytes.NewBufferString("Enter param\n"), message: "Key1", param: "parame"}, "Enter param"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stdReader(tt.args.stdin, tt.args.message, tt.args.param); got != tt.want {
				t.Errorf("stdReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
