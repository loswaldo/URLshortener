package serverapi

import (
	"github.com/gorilla/mux"
	"github.com/loswaldo/URLshortener/internal/app/repository"
	"github.com/loswaldo/URLshortener/internal/app/repository/inmemory"
	"github.com/matoous/go-nanoid"
	"text/template"
	//"github.com/loswaldo/URLshortener/internal/app/store"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ServerAPI struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *repository.Rep
	//store *store.Store
}

func New(config *Config) *ServerAPI {
	return &ServerAPI{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  repository.NewRep(inmemory.NewInMemoryDB()),
	}
}

func (s *ServerAPI) Start() error {
	s.logger.SetLevel(logrus.DebugLevel)

	s.configureRouter()

	//if err := s.configureStore(); err != nil{
	//	return err
	//}
	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ServerAPI) configureRouter() {
	s.router.HandleFunc("/URLShortener", s.helloHandler())
}

func (s *ServerAPI) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("./getShort.html")
		if err != nil {
			panic(err)
		}
		err = html.Execute(w, nil)
		if err != nil {
			panic(err)
		}
		longUrl := r.FormValue("longUrl")
		var shortURL string
		if longUrl != "" {
			id, err := gonanoid.Nanoid()
			if err != nil {
				panic(err)
			}
			shortURL = "http://localhost:8080/URLShortener" + id
			err = s.store.AddNewURL(longUrl, shortURL)
			if err != nil {
				/*todo: error*/
			}
		}

		io.WriteString(w, shortURL)
		//io.WriteString(w, longUrl)
		//io.WriteString(w, "Hello")
	}
}
