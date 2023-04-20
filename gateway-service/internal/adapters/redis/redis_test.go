package redis

import (
	"log"
	"testing"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE0ODgxODYsImp0aSI6IjhlZjQ4YTEzLWY3NzEtNDU2Yi1hNWJmLTMxNmM1YTY4YWM1MiJ9.lZfI1ORh6gvDbDj8zRsa4yjzJDO6KY01bJeU6_YTbTM"

func TestSetToken(t *testing.T) {

	redisA := NewAdapter()

	err := redisA.SetToken(token)

	if err != nil {
		log.Printf("Error testing SetToken(token) %v", err)
	}

}

func TestIsHasToken(t *testing.T) {

	redisA := NewAdapter()

	result := redisA.IsHasToken(token)

	log.Printf("Result = %v", result)

}
