package main

import (
	"reflect"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
)

func Test_createTwitterClient(t *testing.T) {
	type args struct {
		configName string
	}
	tests := []struct {
		name string
		args args
		want *twitter.Client
	}{
		{name: "case 1", args: args{configName: ".goshirase/config"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTwitterClient(tt.args.configName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTwitterClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
