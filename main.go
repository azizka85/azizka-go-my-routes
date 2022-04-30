package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/azizka85/azizka-go-my-routes/home"
	"github.com/azizka85/azizka-go-my-routes/oAuth"
	"github.com/azizka85/azizka-go-my-routes/signIn"
	"github.com/azizka85/azizka-go-my-routes/signOut"
	"github.com/azizka85/azizka-go-my-routes/signUp"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/wader/gormstore/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	} else {
		global.Db = db
		helpers.MigrateSchemas(db)
	}

	global.SessionStore = gormstore.New(global.Db, []byte(os.Getenv("SESSION_SECRET")))

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 3000
	}

	router := mux.NewRouter()

	subRouter := router.
		PathPrefix(global.Settings.PageRoot).
		Subrouter()

	oAuth.AddRoutes(subRouter)

	signIn.AddRoutes(subRouter)
	signUp.AddRoutes(subRouter)

	signOut.AddRoutes(subRouter)

	home.AddRoutes(subRouter)

	http.Handle("/", router)

	address := fmt.Sprintf(":%v", port)

	fmt.Printf("Listening %v\n", address)

	log.Fatal(http.ListenAndServe(address, nil))
}
