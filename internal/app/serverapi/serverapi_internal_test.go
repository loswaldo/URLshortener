package serverapi

import (
	"fmt"
	"github.com/loswaldo/URLshortener/internal/app/repository"
	"github.com/loswaldo/URLshortener/internal/app/repository/inmemory"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerAPI_URLShortenerHandler(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	s.SetStorage(repository.NewRep(inmemory.NewInMemoryDB()))
	longURL := "https://gobyexample.com/environment-variables"
	htmlFile := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>get short</title>\n</head>\n<body>\n<form method=\"POST\">\n    <label>long url</label><br>\n    <div><input type=\"text\" name=\"longUrl\"></div>\n    <div><input type=\"submit\"></div>\n\n</form>\n\n<form method=\"GET\">\n    <label>short url</label><br>\n    <div><input type=\"text\" name=\"shortUrl\"></div>\n    <div><input type=\"submit\"></div>\n\n</form>\n</body>\n</html>"
	shortUrl := "http://localhost:8080/sh?tocken=123"
	err := s.store.AddNewURL(longURL, shortUrl)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/URLShortener?longUrl=%s", longURL), nil)
	s.URLShortenerHandler().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), htmlFile+shortUrl)
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/URLShortener?shortUrl=%s", shortUrl), nil)
	s.URLShortenerHandler().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), htmlFile+shortUrl+htmlFile+longURL)

}
