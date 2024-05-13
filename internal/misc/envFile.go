package misc

import (
	"errors"
	"os"
	"strings"
)

type envMAP map[string]interface{}

var envMapInstance envMAP = Initialise()

func Load(filename string) {
	fileBodyByte, err := os.ReadFile(filename)
	ErrorHandling(err)

	fileBody := strings.Split(string(fileBodyByte), "\n")

	for _, i := range fileBody {
		keyValue := strings.Split(i, "=")
		if len(keyValue) == 2 {
			envMapInstance[keyValue[0]] = keyValue[1]
		}
	}
}

func ( mapper envMAP ) FindENV(key string) (interface{}, error) {
	if result := mapper[key]; result != nil {
		return result, nil
	}
	return nil, errors.New("key not found")
}

func Initialise() envMAP{
	return make(envMAP)
}