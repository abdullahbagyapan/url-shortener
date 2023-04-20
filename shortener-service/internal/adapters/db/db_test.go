package db

import (
	"log"
	"testing"
)

var longURL = "www.google.com"
var shortURL = "abc"

func TestSaveURL(t *testing.T) {

	dbAdapter := NewAdapter()

	err := dbAdapter.SaveURL(longURL, shortURL)

	if err != nil {
		log.Fatalf("FAILED TestSaveURL(%s,%s) , error : %v", longURL, shortURL, err)
	}

	log.Printf("PASSED TestSaveURL(%s,%s)", longURL, shortURL)

}

func TestGetURL(t *testing.T) {

	dbAdapter := NewAdapter()

	longURLResult, err := dbAdapter.GetURL(shortURL)

	if err != nil {
		log.Fatalf("FAILED TestGetURL(%s) , error : %v", shortURL, err)
	}

	log.Printf("PASSED TestGetURL(%s) , LongURL : %v", shortURL, longURLResult)

}
