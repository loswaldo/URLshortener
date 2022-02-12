package main

import (
	"github.com/go-yaml/yaml"
	"github.com/loswaldo/URLshortener/internal/app/serverapi"
	"io/ioutil"
	"log"
)

func main() {
	// read config
	// if postgre{
	//	create postgre table
	//} else{
	//	create inmemory table
	//}
	/*need create interface for two db*/

	//need struct for server; contains http.Server
	config := loadConfig()
	s := serverapi.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}

func loadConfig() *serverapi.Config {
	config := serverapi.NewConfig()
	yamlFile, err := ioutil.ReadFile("./configs/postgre.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatal(err)
	}
	return config

}
