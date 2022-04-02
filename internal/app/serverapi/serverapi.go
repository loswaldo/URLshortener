package serverapi

import (
	"errors"
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

		filePath := "/internal/app/serverapi/getShort.html"
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
		if r.Method == http.MethodPost {
			if err, code := URLShortenerPost(w, r, s); err != nil {
				http.Error(w, http.StatusText(code), code)
			}
		}
		if r.Method == http.MethodGet {
			if err, code := URLShortenerGet(w, r, s); err != nil {
				http.Error(w, http.StatusText(code), code)
			}
		}
	}
}
func URLShortenerGet(w http.ResponseWriter, r *http.Request, s *ServerAPI) (error, int) {
	shortURL := r.FormValue("shortUrl")
	if shortURL != "" {
		longURL, err := s.store.GetLongURL(shortURL)
		if err != nil {
			return err, http.StatusInternalServerError
		} else if longURL == "" {
			return err, http.StatusBadRequest
		} else {
			io.WriteString(w, longURL)
			return nil, 0
		}
	}
	return errors.New("bad request error"), http.StatusBadRequest
}
func URLShortenerPost(w http.ResponseWriter, r *http.Request, s *ServerAPI) (error, int) {
	var gonan = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	longUrl := strings.Trim(r.FormValue("longUrl"), " ")
	if _, err := url.ParseRequestURI(longUrl); err != nil || longUrl == "" {
		return err, http.StatusBadRequest
	}
	var shortURL string
	id, err := gonanoid.Generate(gonan, 10)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	shortURL = "http://localhost:8080/sh?tocken=" + id
	err = s.store.AddNewURL(longUrl, shortURL)
	if err != nil {
		if value, err := s.store.GetShortURL(longUrl); err != nil {
			return err, http.StatusBadRequest
		} else {
			io.WriteString(w, value)
			return nil, 0
		}
	}
	io.WriteString(w, shortURL)
	return nil, 0

}
