package config

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"
	"github.com/rijalfm/go-todo/internal/domain/service"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "1234567890",
		Addr: "127.0.0.1:5432",
		Database: "db_todo",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	service.CreateTodoTable(db)
	service.InitiateDB(db)
	return db
}