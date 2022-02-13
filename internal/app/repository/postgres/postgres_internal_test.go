package postgres

import "testing"

func createTestDB(t *testing.T) (*PostgresDB, func()) {
	t.Helper()
	conf := CreateConfig()
	conf.DBName = "test"
	db, err := NewPostgresDB(conf)
	if err != nil {
		t.Fatal(err)
	}
	return db, func() {
		if _, err := db.db.Exec("DROP TABLE IF EXISTS %s", "test"); err != nil {
			t.Fatal(err)
		}
	}
}

func TestPostgresDB_AddNewURL(t *testing.T) {
	db, fc := createTestDB(t)
	defer fc()
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

func TestPostgresDB_GetLongURL(t *testing.T) {
	db, fc := createTestDB(t)
	defer fc()
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

func TestPostgresDB_GetShortURL(t *testing.T) {
	db, fc := createTestDB(t)
	defer fc()
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
