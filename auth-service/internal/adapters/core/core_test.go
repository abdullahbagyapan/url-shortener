package core

import (
	"log"
	"testing"
)

var password = []byte{1, 3, 4, 'a', 'd', 4, 6, 'x'}

func TestHash(t *testing.T) {

	coreAdapter := NewAdapter()

	_, err := coreAdapter.Hash(password)

	if err != nil {
		log.Printf("FAILED TestHash(%s) , ERROR : %v", password, err)
	}

	log.Printf("PASSED TestHash(%s)")

}
