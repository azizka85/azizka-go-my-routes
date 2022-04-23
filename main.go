package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/azizka85/azizka-go-my-routes/home"
	"github.com/azizka85/azizka-go-my-routes/oAuth"
	"github.com/azizka85/azizka-go-my-routes/settings"
	"github.com/azizka85/azizka-go-my-routes/signIn"
	"github.com/azizka85/azizka-go-my-routes/signUp"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 3000
	}

	router := mux.NewRouter()

	subRouter := router.
		PathPrefix(settings.GlobalSettings.PageRoot).
		Subrouter()

	oAuth.AddRoutes(subRouter)

	signIn.AddRoutes(subRouter)
	signUp.AddRoutes(subRouter)

	home.AddRoutes(subRouter)

	http.Handle("/", router)

	address := fmt.Sprintf(":%v", port)

	fmt.Printf("Listening %v\n", address)

	log.Fatal(http.ListenAndServe(address, nil))
}
