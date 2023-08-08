package models

import (
	"database/sql"
	"log"
)

var (
	db *sql.DB
	UserRepo = NewUserRepository(db)
	PostRepo = NewPostRepository(db)
)

func init() {
	d, err := sql.Open("sqlite3", "./data/sql/forum.db")
	if err != nil {
		log.Fatal("❌ Couldn't open the database")
	}
	db = d
	if err = db.Ping(); err != nil {
		log.Fatal("❌ Connection to the database is dead")
	}
	defer db.Close()
}
