package oAuth

import (
	"fmt"
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/oAuth/handlers"
	"github.com/gorilla/mux"
)

var gitHub = handlers.GitHub{}

func Default(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var err error = nil

	switch vars["service"] {
	case "github":
		err = gitHub.Handle(w, r)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func Callback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var err error = nil

	switch vars["service"] {
	case "github":
		err = gitHub.Callback(w, r)
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/oauth/callback/{service}", Callback)
	router.HandleFunc("/oauth/{service}", Default)
}
