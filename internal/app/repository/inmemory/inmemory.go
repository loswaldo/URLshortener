package inmemory

import "errors"

type InMemoryDB struct {
	store map[string]string // key - longURL value - shortURL
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{store: make(map[string]string)}
}

func (db *InMemoryDB) GetLongURL(shortURL string) (string, error) {
	for key, value := range db.store {
		if value == shortURL {
			return key, nil
		}
	}
	return "", nil
}

func (db *InMemoryDB) GetShortURL(longURL string) (string, error) {
	if value, ok := db.store[longURL]; ok {
		return value, nil
	}
	return "", nil
}

func (db *InMemoryDB) AddNewURL(longURL, shortURL string) error {
	if value, _ := db.GetLongURL(shortURL); value != "" {
		return errors.New("I have this URL in my db")
	}
	if value, _ := db.GetShortURL(longURL); value != "" {
		return errors.New("I have this URL in my db")
	}

	db.store[longURL] = shortURL
	return nil
}
