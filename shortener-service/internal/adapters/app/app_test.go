package app

import (
	"log"
	"testing"
	redis "url-shortener/shortener-service/internal/adapters/cache"
	"url-shortener/shortener-service/internal/adapters/core"
	"url-shortener/shortener-service/internal/adapters/db"
)

var longURL = "www.google.com"
var shortURL string

func TestCreateURL(t *testing.T) {

	dbAdapter := db.NewAdapter()
	redisAdapter := redis.NewAdapter()

	coreAdapter := core.NewAdapter()

	appAdapter := NewAdapter(coreAdapter, dbAdapter, redisAdapter)

	shortURLResult, err := appAdapter.CreateURL(longURL)

	if err != nil {
		log.Fatalf("Failed TestCreateURL(%s) , error =%v", longURL, err)
	}
	log.Printf("Passed TestCreateURL(%s) shortURL : %v ", longURL, shortURLResult)

	shortURL = shortURLResult

}

func TestGetURL(t *testing.T) {

	dbAdapter := db.NewAdapter()
	redisAdapter := redis.NewAdapter()

	coreAdapter := core.NewAdapter()

	appAdapter := NewAdapter(coreAdapter, dbAdapter, redisAdapter)

	longURLResult, err := appAdapter.GetURL(shortURL)

	if err != nil {
		log.Fatalf("Failed TestGetURL(%s) , error =%v", shortURL, err)
	}
	log.Printf("Passed TestGetURL(%s) longURL : %v ", shortURL, longURLResult)

}
