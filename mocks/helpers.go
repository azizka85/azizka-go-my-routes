package mocks

import (
	"database/sql"
	"os"

	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/wader/gormstore/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func PrepareForTesting() (*mux.Router, *sql.DB, error) {
	godotenv.Load()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		return nil, nil, err
	} else {
		global.Db = db
		helpers.MigrateSchemas(db)
	}

	global.SessionStore = gormstore.New(global.Db, []byte(os.Getenv("SESSION_SECRET")))

	router := mux.NewRouter()

	subRouter := router.
		PathPrefix(global.Settings.PageRoot).
		Subrouter()

	dbInstance, err := db.DB()

	return subRouter, dbInstance, err
}
