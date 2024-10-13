package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Upsert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w)

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
}

func main() {
	http.HandleFunc("/", Upsert)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
