package app

import (
	"log"
	"testing"
	"url-shortener/token-service/internal/adapters/core"
)

var userID = "2662f7c5-c619-4a57-84f2-54d28819405d"

var token string

func TestGenerateToken(t *testing.T) {

	coreAdapter := core.NewCoreAdapter()

	appAdapter := NewAppAdapter(coreAdapter)

	generatedToken, err := appAdapter.GenerateToken(userID)

	if err != nil {
		log.Fatalf("FAILED  TestGenerateToken(%s) , ERROR : %v", userID, err)
	}
	log.Printf("PASSED TestGenerateToken(%s) , TOKEN = %v", userID, generatedToken)
	token = generatedToken
}

func TestValidateToken(t *testing.T) {

	coreAdapter := core.NewCoreAdapter()

	appAdapter := NewAppAdapter(coreAdapter)

	userIDResult, isValid, err := appAdapter.ValidateToken(token)

	if err != nil {
		log.Fatalf("FAILED  TestValidateToken(%s) , ERROR : %v", token, err)
	}

	if !isValid {
		log.Fatalf("FAILED  TestValidateToken(%s) , TOKEN IS NOT VALID !", token)
	}

	log.Printf("PASSED TestGenerateToken(%s) , USERID = %v", token, userIDResult)

}
