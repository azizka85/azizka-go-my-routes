package signUp

import (
	"fmt"
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/azizka85/azizka-go-my-routes/templates"
	"github.com/azizka85/azizka-go-my-routes/templates/components"
	"github.com/azizka85/azizka-go-my-routes/templates/pages"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func Default(w http.ResponseWriter, r *http.Request) {
	lang,
		ajax,
		init,
		query,
		_,
		user,
		translator,
		status,
		err := helpers.ViewData(w, r)

	content := ""

	if status == http.StatusOK {
		content = templates.RenderPage(
			ajax == "1",
			lang,
			pages.SignUp(
				global.Settings.PageRoot,
				lang,
				components.AuthService(
					lang,
					global.Settings.PageRoot,
					translator,
				),
				translator,
			),
			[]string{},
			query,
			r,
			user,
			&global.Settings,
			translator,
		)
	}

	helpers.RenderDataOrRedirect(
		w,
		r,
		nil,
		content,
		ajax,
		init,
		"",
		status,
		err,
	)
}

func Post(w http.ResponseWriter, r *http.Request) {
	_,
		ajax,
		init,
		_,
		session,
		_,
		translator,
		status,
		err := helpers.ViewData(w, r)

	var user *data.User

	if status == http.StatusOK {
		user = &data.User{}

		decoder := schema.NewDecoder()

		r.ParseForm()

		err = decoder.Decode(user, r.PostForm)

		if err != nil {
			fmt.Println(err)

			status = http.StatusBadRequest
		} else {
			err = helpers.SignUp(
				user,
				translator,
				session,
				global.Db,
			)

			if err != nil {
				fmt.Println(err)

				status = http.StatusBadRequest
			} else {
				err = session.Save(r, w)

				if err != nil {
					status = http.StatusInternalServerError
				} else {
					status = http.StatusFound
				}
			}
		}
	}

	// TODO: change location
	helpers.RenderDataOrRedirect(
		w,
		r,
		user,
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
		HandleFunc("/sign-up", Default).
		Methods(http.MethodGet)

	router.
		HandleFunc("/{lang}/sign-up", Default).
		Methods(http.MethodGet)

	router.
		HandleFunc("/sign-up", Post).
		Methods(http.MethodPost)

	router.
		HandleFunc("/{lang}/sign-up", Post).
		Methods(http.MethodPost)
}
