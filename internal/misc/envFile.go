package misc

import (
	"os"
	"strings"
)

var envMap map[interface{}]interface{}

func Load(filename string) {
	fileBodyByte, err := os.ReadFile(filename)
	ErrorHandling(err)

	fileBody := strings.Split(string(fileBodyByte), "\n")
}