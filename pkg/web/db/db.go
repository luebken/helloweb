package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func (database *Database) Open() (err error) {
	dbPath := os.Getenv("DB_PATH")
	fmt.Printf("Using DB_PATH: '%v'\n", dbPath)
	database.DB, err = sql.Open("sqlite3", dbPath+"/sqlite3.db")
	return err
}

func (database *Database) Close() {
	database.DB.Close()
}

func (db *Database) GetUser() User {
	user := User{}

	query := `SELECT email, name from USER`
	row := db.DB.QueryRow(query)
	err := row.Scan(&user.Email, &user.Name)
	if err != nil {
		fmt.Printf("Err %v\n", err)
	}

	return user
}
