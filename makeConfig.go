package main

import (
	"os"
)

const CONF_DIR string = ".goshirase"

func mkConfigDir() {
	if _, err := os.Stat(CONF_DIR); os.IsNotExist(err) {
		os.Mkdir(CONF_DIR, 0744)
	}
}

func mkConfigFile(fileName string) (string, error) {
	result := CONF_DIR + "/" + fileName
	_, err := os.Create(fileName)
	return result, err
}
