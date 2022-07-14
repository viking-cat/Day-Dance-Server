package json

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/viking-cat/day-dance-server/src/structs"
)

func GetJson(file string) (*structs.ServerConfig, error) {
	log.Println("Reading " + file)

	rawFile, fileError := ioutil.ReadFile(file)
	if fileError != nil {
		log.Println(fileError)
		return nil, fileError
	}

	var config structs.ServerConfig

	marshalError := json.Unmarshal([]byte(rawFile), &config)
	if marshalError != nil {
		log.Println(marshalError)
		return nil, marshalError
	}

	log.Println("All: ", config)
	log.Println("port: ", config.Port)
	log.Println("port type: ", reflect.TypeOf(config.Port))
	log.Println("name type: ", config.Name)

	return &config, nil
}
