package home

import (
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/azizka85/azizka-go-my-routes/templates"
	"github.com/azizka85/azizka-go-my-routes/templates/pages"
	"github.com/gorilla/mux"
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

	var layoutNames []string

	if ajax != "1" || init == "1" {
		layoutNames = []string{"main-layout"}

		if ajax == "1" {
			layoutNames = helpers.StringToArray(query.Get("layouts"))
		}
	}

	content := ""

	if status == http.StatusOK {
		content = templates.RenderPage(
			ajax == "1",
			lang,
			pages.Home(),
			layoutNames,
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

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/", Default)
	router.HandleFunc("/{lang}", Default)
}
