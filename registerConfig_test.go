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
		want    *Config
		wantErr bool
	}{
		{"case 1", args{"testconfig/config.json"}, &Config{ConsumerKey: "param1", ConsumerSecret: "param2", AccessToken: "param3", AccessSecret: "param4"}, false},
		{"case 2", args{"testconfig/config2.json"}, &Config{ConsumerKey: "AAA", ConsumerSecret: "BBB", AccessToken: "CCC", AccessSecret: "DDD"}, false},
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

var testConfig1 = &Config{ConsumerKey: "param1", ConsumerSecret: "param2", AccessToken: "param3", AccessSecret: "param4"}
var testConfig2 = &Config{ConsumerKey: "AAA", ConsumerSecret: "CCC", AccessToken: "VVV", AccessSecret: "BBB"}

func Test_makeConfigJson(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case 1", args{config: testConfig1}, "{\"TWITTER_CONSUMER_KEY\":\"param1\",\"TWITTER_CONSUMER_SECRET\":\"param2\",\"TWITTER_ACCESS_TOKEN\":\"param3\",\"TWITTER_ACCESS_SECRET\":\"param4\"}", false},
		{"case 2", args{config: testConfig2}, "{\"TWITTER_CONSUMER_KEY\":\"AAA\",\"TWITTER_CONSUMER_SECRET\":\"CCC\",\"TWITTER_ACCESS_TOKEN\":\"VVV\",\"TWITTER_ACCESS_SECRET\":\"BBB\"}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeConfigJson(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeConfigJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got[:]), tt.want) {
				t.Errorf("makeConfigJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeConfig(t *testing.T) {
	type args struct {
		name string
		data []byte
	}
	case1, _ := makeConfigJson(testConfig1)
	case2, _ := makeConfigJson(testConfig2)
	tests := []struct {
		name string
		args args
	}{
		{"case 1", args{"TEST_CONFIG1", case1}},
		{"case 2", args{"TEST_CONFIG2", case2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeConfig(tt.args.name, tt.args.data)
		})
	}
}
