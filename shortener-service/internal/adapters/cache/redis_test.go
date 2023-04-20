package redis

import (
	"log"
	"testing"
)

var shortURL = "YPR0hHOwG6"
var longURL = "www.google.com"

func TestSaveURL(t *testing.T) {

	redisA := NewAdapter()

	err := redisA.SaveURL(shortURL, longURL)

	if err != nil {
		log.Fatalf("Failed SaveURL(%s,%s) , error =%v", shortURL, longURL, err)
	}
	log.Printf("Passed SaveURL(%s,%s) ", shortURL, longURL)

}

func TestGetURL(t *testing.T) {

	redisA := NewAdapter()

	longURLResult := redisA.GetURL(shortURL)

	if longURLResult == "" {
		log.Fatalf("Failed GetURL(%s) , error = URL NOT FOUND", shortURL)
	}

	log.Printf("Passed GetURL(%s) , LongURL = %s", shortURL, longURLResult)

}
