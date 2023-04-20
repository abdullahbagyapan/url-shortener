package app

import (
	"log"
	"url-shortener/shortener-service/internal/ports"
)

type Adapter struct {
	core  ports.CorePort
	db    ports.DbPort
	redis ports.CachePort
}

func NewAdapter(core ports.CorePort, db ports.DbPort, redis ports.CachePort) *Adapter {

	return &Adapter{db: db, core: core, redis: redis}

}

func (A Adapter) CreateURL(longURL string) (string, error) {

	// core
	shortURL, err := A.core.RandomURL()

	if err != nil {
		log.Printf("Error creating URL %v", err)
	}

	// database
	err = A.db.SaveURL(longURL, shortURL)

	if err != nil {
		log.Printf("Error saving URL %v", err)
	}

	// cache
	go func() {
		err = A.redis.SaveURL(shortURL, longURL)

		if err != nil {
			log.Printf("Error putting cache data %v", err)
		}

	}()

	return shortURL, nil
}

func (A Adapter) GetURL(shortURL string) (string, error) {

	url := A.redis.GetURL(shortURL)

	url, err := A.db.GetURL(shortURL)

	if err != nil {
		log.Printf("Error getting URL from database %v", err)
		return "", err
	}

	return url, nil
}

func (A Adapter) DeleteURL(id string) error {

	err := A.db.DeleteURL(id)

	if err != nil {
		log.Printf("Error deleting url %v", err)
		return err
	}

	return nil
}
