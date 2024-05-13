package misc

import (
	"errors"
	"os"
	"strings"
)

type envMAP map[string]interface{}

var envMap envMAP

func Load(filename string) {
	fileBodyByte, err := os.ReadFile(filename)
	ErrorHandling(err)

	fileBody := strings.Split(string(fileBodyByte), "\n")

	for _, i := range fileBody {
		keyValue := strings.Split(i, "=")
		if len(keyValue) == 2 {
			envMap[keyValue[0]] = keyValue[1]
		}
	}
}

func ( mapper envMAP ) FindENV(key string) (interface{}, error) {
	if result := mapper[key]; result != nil {
		return result, nil
	}
	return nil, errors.New("key not found")
}