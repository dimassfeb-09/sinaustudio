package app

import (
	"database/sql"
	"log"
)

func ConnectionDatabases() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sinaustudio?autocommit=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
