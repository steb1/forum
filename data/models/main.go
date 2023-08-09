package models

import (
	"database/sql"
	"log"
	"os"
)

var (
	db       *sql.DB
	UserRepo *UserRepository
	PostRepo *PostRepository
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

	query, err := os.ReadFile("./data/sql/init.sql")
	if err != nil {
		log.Fatal("couldn't read setup.sql")
	}
	if _, err = db.Exec(string(query)); err != nil {
		log.Fatal("database setup wasn't successful")
	}

	UserRepo = NewUserRepository(db)
	PostRepo = NewPostRepository(db)
	log.Println("✅ Database init with success")
}
