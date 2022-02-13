package inmemory

import (
	"testing"
)

func TestInMemoryDB_AddNewURL(t *testing.T) {
	db := NewInMemoryDB()
	longURL := "https://test/test"
	shortUrl := "http://localhost:8080/sh?tocken=123"
	err := db.AddNewURL(longURL, shortUrl)
	if err != nil {
		t.Error()
	}
	err = db.AddNewURL(longURL, shortUrl)
	if err == nil {
		t.Error()
	}
}

func TestInMemoryDB_GetLongURL(t *testing.T) {
	db := NewInMemoryDB()
	longURL := "https://test/test"
	shortUrl := "http://localhost:8080/sh?tocken=123"
	err := db.AddNewURL(longURL, shortUrl)
	if err != nil {
		t.Error()
	}
	expLongURL, err := db.GetLongURL(shortUrl)
	if longURL != expLongURL {
		t.Error()
	}
	expLongURL, err = db.GetLongURL("buba")
	if expLongURL != "" {
		t.Error()
	}
}

func TestInMemoryDB_GetShortURL(t *testing.T) {
	db := NewInMemoryDB()
	longURL := "https://test/test"
	shortUrl := "http://localhost:8080/sh?tocken=123"
	err := db.AddNewURL(longURL, shortUrl)
	if err != nil {
		t.Error()
	}
	expShortURL, err := db.GetShortURL(longURL)
	if shortUrl != expShortURL {
		t.Error()
	}
	expShortURL, err = db.GetShortURL("buba")
	if expShortURL != "" {
		t.Error()
	}
}
