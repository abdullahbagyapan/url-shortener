package db

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"url-shortener/auth-service/internal/ports"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter() *Adapter {

	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/users")

	if err != nil {
		log.Fatalf("Error connecting to db %v", err)
	}

	// CONNECTION TEST
	err = db.Ping()

	if err != nil {
		log.Fatalf("Error testing to db %v", err)
	}

	// CREATE TABLE
	query := "CREATE TABLE IF NOT EXISTS users(id VARCHAR(255) PRIMARY KEY,name VARCHAR(255) NOT NULL,username VARCHAR(255) NOT NULL , email  VARCHAR(50) NOT NULL,password VARCHAR(255) NOT NULL,date timestamp NOT NULL)"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err = db.ExecContext(ctx, query)

	if err != nil {
		log.Fatalf("Error %s when creating table", err)
	}

	log.Printf("Succesfully connected to the database")
	return &Adapter{db: db}
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()

	if err != nil {
		log.Fatalf("Error closing to db connection %v", err)
	}
}

func (da Adapter) SaveUser(user *ports.User) error {

	query := "INSERT INTO users(id, name,username,email,password, date) VALUES (?,?,?,?,?,?)"

	_, err := da.db.Exec(query, user.Id, user.Name, user.Username, user.Email, user.Password, time.Now())

	return err

}

func (da Adapter) FindByUsername(username string) (*ports.User, error) {

	query := "SELECT id,password FROM users WHERE username = ?"

	rows, err := da.db.Query(query, username)

	if err != nil {
		return nil, err
	}

	var user ports.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Password)

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
