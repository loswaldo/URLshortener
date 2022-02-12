package repository

type Store interface {
	GetLongURL(shortURL string) string
	GetShortURL(longURL string) string
	AddNewURL(longURL, shortURL string) error
}

type Rep struct {
	Store
}

func NewRep(db Store) *Rep {
	return &Rep{
		Store: db,
	}
}
