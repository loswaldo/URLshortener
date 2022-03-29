package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/loswaldo/URLshortener/internal/app/repository/inmemory"

	"github.com/go-yaml/yaml"
	"github.com/loswaldo/URLshortener/internal/app/repository"
	"github.com/loswaldo/URLshortener/internal/app/repository/postgres"
	"github.com/loswaldo/URLshortener/internal/app/serverapi"
)

func main() {
	config := loadConfig()
	s := serverapi.New(config)
	var storage repository.Store
	if value, ok := os.LookupEnv("STORAGE"); ok == true && value == "POSTGRES" {
		var err error
		storage, err = postgres.NewPostgresDB(postgres.CreateConfig())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		storage = inmemory.NewInMemoryDB()
	}
	s.SetStorage(repository.NewRep(storage))
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
