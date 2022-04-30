package signOut

import (
	"fmt"
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/gorilla/mux"
)

func Default(w http.ResponseWriter, r *http.Request) {
	session, err := global.SessionStore.Get(r, global.SessionKey)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		query := r.URL.Query()

		ajax := query.Get("ajax")

		helpers.SignOut(session)
		err = session.Save(r, w)

		if err != nil {
			if ajax == "1" {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")

				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// TODO: not ajax
			}
		} else {
			if ajax == "1" {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")

				http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
			} else {
				// TODO: not ajax
			}
		}
	}
}

func AddRoutes(router *mux.Router) {
	router.
		HandleFunc("/sign-out", Default).
		Methods(http.MethodGet)
}
