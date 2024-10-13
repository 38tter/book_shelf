package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	path := os.Getenv("INSTANCE_UNIX_SOCKET")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("pgx", fmt.Sprintf("user=%s password=%s database=%s host=%s",
		user, password, dbName, path))
	if err != nil {
		fmt.Errorf("failed to connect database")

		return
	}

	upsert := fmt.Sprintf("INSERT INTO books(name, author) VALUES ('%s','%s');", "foo", "bar")
	_, err = db.Exec(upsert)
	if err != nil {
		log.Fatalf("SQL ERROR: %s", err)
	}

	return
}
