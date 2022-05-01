package signOut

import (
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/gorilla/mux"
)

func Default(w http.ResponseWriter, r *http.Request) {
	_,
		ajax,
		init,
		_,
		session,
		_,
		_,
		status,
		err := helpers.ViewData(w, r)

	if status == http.StatusOK {
		helpers.SignOut(session)
		err = session.Save(r, w)

		if err != nil {
			status = http.StatusInternalServerError
		} else {
			status = http.StatusFound
		}
	}

	// TODO: change location
	helpers.RenderDataOrRedirect(
		w,
		r,
		nil,
		"",
		ajax,
		init,
		"",
		status,
		err,
	)
}

func AddRoutes(router *mux.Router) {
	router.
		HandleFunc("/sign-out", Default).
		Methods(http.MethodGet)
}
