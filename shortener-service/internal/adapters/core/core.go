package core

import (
	"math/rand"
)

type Adapter struct {
}

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewAdapter() *Adapter {
	return new(Adapter)
}

func (A Adapter) RandomURL() (string, error) {

	var generatedURL = make([]rune, 10)
	var i = 0

	for i = 0; i < 10; i++ {

		randInt := rand.Intn(len(runes))
		generatedURL[i] = runes[randInt]
	}

	return string(generatedURL), nil

}
