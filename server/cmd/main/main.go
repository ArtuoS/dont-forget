package main

import (
	"database/sql"

	"github.com/ArtuoS/dont-forget/internal/database"
	"github.com/ArtuoS/dont-forget/internal/database/api"
	"github.com/ArtuoS/dont-forget/internal/database/repository"
	_ "github.com/mattn/go-sqlite3"
)

var (
	itemRepo *repository.ItemRepository
)

func main() {
	db, err := sql.Open("sqlite3", "../../items.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.NewDatabase(db).Create()
	if err != nil {
		panic(err)
	}

	itemRepo = repository.NewItemRepository(db)
	api.NewRouter(itemRepo).StartRouting()
}
