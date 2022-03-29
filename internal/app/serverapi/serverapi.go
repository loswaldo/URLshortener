package serverapi

import (
	"github.com/gorilla/mux"
	"github.com/loswaldo/URLshortener/internal/app/repository"
	"github.com/matoous/go-nanoid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
)

type ServerAPI struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  repository.Store
}

func New(config *Config) *ServerAPI {
	return &ServerAPI{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  nil,
	}
}

func (s *ServerAPI) SetStorage(r repository.Store) {
	s.store = r
}

func (s *ServerAPI) Start() error {
	s.logger.SetLevel(logrus.DebugLevel)

	s.configureRouter()

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.Addr, s.router)
}

func (s *ServerAPI) configureRouter() {
	s.router.HandleFunc("/URLShortener", s.URLShortenerHandler())
}

func (s *ServerAPI) URLShortenerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gonan = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
		filePath := /*strings.TrimPrefix(*/ "/internal/app/serverapi/getShort.html" /*, "/internal/app/serverapi")*/
		filePath = os.Getenv("PWD") + filePath
		html, err := template.ParseFiles(filePath)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = html.Execute(w, nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if r.Method == http.MethodPost { //получение короткой ссылки
			longUrl := strings.Trim(r.FormValue("longUrl"), " ")
			if _, err := url.ParseRequestURI(longUrl); err != nil || longUrl == "" { // если нам пришел не url или пустая строка, то bad request
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			var shortURL string
			id, err := gonanoid.Generate(gonan, 10) //генерация токена
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			shortURL = "http://localhost:8080/sh?tocken=" + id //генерация короткой ссылки
			err = s.store.AddNewURL(longUrl, shortURL)         //добавление в базу данных
			if err != nil {                                    //если не получилось добавить в бд
				if value, err := s.store.GetShortURL(longUrl); err != nil {
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				} else {
					io.WriteString(w, value)
				}

			} else {
				io.WriteString(w, shortURL)
			}
		}
		if r.Method == http.MethodGet {
			shortURL := r.FormValue("shortUrl")
			var longURL string
			if shortURL != "" {
				longURL, err = s.store.GetLongURL(shortURL)
				if err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				} else if longURL == "" {
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				} else {
					io.WriteString(w, longURL)
				}
			} else {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}
	}
}
