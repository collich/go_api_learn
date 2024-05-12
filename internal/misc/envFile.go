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

	for _, i := range fileBody {
		keyValue := strings.Split(i, "=")
		envMap[keyValue[0]] = keyValue[1]
	}
}