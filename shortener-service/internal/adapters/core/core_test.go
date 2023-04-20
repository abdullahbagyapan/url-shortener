package core

import (
	"log"
	"testing"
)

func TestRandomURL(t *testing.T) {

	core := NewAdapter()

	result, err := core.RandomURL()

	if err != nil {
		log.Printf("Test failed!! Error getting random url %v", err)
	}

	log.Printf("Test passed!! random URL = %s", result)

}
