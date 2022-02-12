package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/loswaldo/URLshortener/internal/app/repository"
	"github.com/loswaldo/URLshortener/internal/app/repository/inmemory"
	"github.com/loswaldo/URLshortener/internal/app/repository/postgres"
	"github.com/loswaldo/URLshortener/internal/app/serverapi"
	"io/ioutil"
	"log"
	"os"
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
	if value, ok := os.LookupEnv("STORAGE"); ok == true {
		fmt.Println(value)
		config := loadConfig()
		s := serverapi.New(config)
		postdb, err := postgres.NewPostgresDB(postgres.CreateConfig())
		if err != nil {
			log.Fatal(err)
			//s.logger.Fatal(err) /*todo: logger func*/
		}
		s.SetStorage(repository.NewRep(postdb))
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("inmemory")
		config := loadConfig()
		s := serverapi.New(config)
		s.SetStorage(repository.NewRep(inmemory.NewInMemoryDB()))
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}

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
