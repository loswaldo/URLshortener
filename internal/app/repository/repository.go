package repository

type Store interface {
	GetLongURL(shortURL string) (string, error)
	GetShortURL(longURL string) (string, error)
	AddNewURL(longURL, shortURL string) error
}
