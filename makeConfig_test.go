package main

import (
	"os"
	"testing"
)

func Test_mkConfigDir(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"execute 1"},
		{"execute 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mkConfigDir()
		})
	}
	// os.Remove(".goshirase")
}

func Test_mkConfigFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
		wantErr  bool
	}{
		{"test1", args{"config3"}, ".goshirase/config3", false},
		{"test2", args{"config4"}, ".goshirase/config4", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mkConfigFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("mkConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if _, err := os.Stat(got); os.IsNotExist(err) {
				t.Errorf("mkConfigFile() error not exist : %v", got)
			}
			if got != tt.wantName {
				t.Errorf("mkConfigFile() = %v, wantName %v", got, tt.wantName)
			}
		})
		os.Remove(tt.wantName)
	}
}
