package models

import (
	"database/sql"
	"forum/lib"
	"log"
	"os"
)

var (
	db           *sql.DB
	UserRepo     *UserRepository
	PostRepo     *PostRepository
	CommentRepo  *CommentRepository
	CategoryRepo *CategoryRepository
)

func init() {
	lib.LoadEnv(".env")
	d, err := sql.Open("sqlite3", os.Getenv("DATABASE"))
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
	CommentRepo = NewCommentRepository(db)
	CategoryRepo = NewCategoryRepository(db)

	log.Println("✅ Database init with success")
}
