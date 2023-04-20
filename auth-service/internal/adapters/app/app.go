package app

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"url-shortener/auth-service/internal/ports"
)

type Adapter struct {
	core ports.CorePort
	db   ports.DbPort
}

func NewAdapter(core ports.CorePort, db ports.DbPort) *Adapter {
	return &Adapter{db: db, core: core}
}

func (A Adapter) Login(username, password string) (string, error) {

	user, err := A.db.FindByUsername(username)

	if err != nil {
		log.Printf("Error querying database %v", err)
		return "", err
	}

	bytePwd := []byte(password)

	err = bcrypt.CompareHashAndPassword(user.Password, bytePwd)

	if err != nil {
		return "", err
	}
	return user.Id, nil

}

func (A Adapter) Register(name, username, password, email string) (string, error) {

	bytePwd := []byte(password)

	hashPwd, err := A.core.Hash(bytePwd)

	if err != nil {
		log.Printf("Error hashing password %v", err)
		return "", err
	}

	var user = ports.User{Id: uuid.New().String(), Password: hashPwd, Username: username, Name: name, Email: email}

	err = A.db.SaveUser(&user)

	if err != nil {
		log.Printf("Error saving user %v", err)
		return "", err
	}

	return user.Id, nil

}
