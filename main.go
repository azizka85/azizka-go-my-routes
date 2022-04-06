package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/settings"
	"github.com/gorilla/mux"
)

func main() {
	port := 3000

	router := mux.NewRouter()

	router.HandleFunc("/hello/{name:(?:Aziz|Ulugbek|Umid)}", func(w http.ResponseWriter, r *http.Request) {
		/* vars := mux.Vars(r)

		fmt.Fprintf(w, "Hello, %v!", vars["name"]) */

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		data, _ := json.Marshal(settings.GlobalSettings)

		fmt.Fprint(w, string(data))
	})

	http.Handle("/", router)

	address := fmt.Sprintf(":%v", port)

	fmt.Printf("Listening %v\n", address)

	log.Fatal(http.ListenAndServe(address, nil))
}
