package core

import (
	"golang.org/x/crypto/bcrypt"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return new(Adapter)
}

func (A Adapter) Hash(password []byte) ([]byte, error) {

	hashPwd, err := bcrypt.GenerateFromPassword(password, 15)

	if err != nil {
		return nil, err
	}

	return hashPwd, nil

}
