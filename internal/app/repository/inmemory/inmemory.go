package inmemory

import "errors"

type InMemoryDB struct {
	store map[string]string // key - longURL value shortURL
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{store: make(map[string]string)}
}

func (db *InMemoryDB) GetLongURL(shortURL string) string {
	for key, value := range db.store {
		if value == shortURL {
			return key
		}
	}
	return ""
}

func (db *InMemoryDB) GetShortURL(longURL string) string {
	value := db.store[longURL]
	return value
}

func (db *InMemoryDB) AddNewURL(longURL, shortURL string) error {
	if value := db.GetLongURL(longURL); value != "" {
		return errors.New("I have this URL in my db")
	}
	if value := db.GetShortURL(shortURL); value != "" {
		return errors.New("I have this URL in my db")
	}

	db.store[longURL] = shortURL
	return nil
}
