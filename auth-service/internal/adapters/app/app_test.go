package app

import (
	"log"
	"testing"
	"url-shortener/auth-service/internal/adapters/core"
	"url-shortener/auth-service/internal/adapters/db"
)

var name = "test"
var username = "username"
var password = "testpassword"
var email = "test@test.com"

func TestRegister(t *testing.T) {

	dbAdapter := db.NewAdapter()

	coreAdapter := core.NewAdapter()

	appAdapter := NewAdapter(coreAdapter, dbAdapter)

	userID, err := appAdapter.Register(name, username, password, email)

	if err != nil {
		log.Fatalf("FAILED TestRegister(%s,%s,%s,%s) , ERROR : %v", name, username, password, email, err)
	}

	log.Printf("PASSED TestRegister(%s,%s,%s,%s) , USERID : %v", name, username, password, email, userID)

}

func TestLogin(t *testing.T) {
	dbAdapter := db.NewAdapter()

	coreAdapter := core.NewAdapter()

	appAdapter := NewAdapter(coreAdapter, dbAdapter)

	userID, err := appAdapter.Login(username, password)

	if err != nil {
		log.Fatalf("FAILED TestLogin(%s,%s) , ERROR : %v", username, password, err)
	}

	log.Printf("PASSED TestLogin(%s,%s) , USERID : %v", username, password, userID)

}
