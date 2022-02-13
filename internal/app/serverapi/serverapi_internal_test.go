package serverapi

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerAPI_URLShortenerHandler(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	s.URLShortenerHandler().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
