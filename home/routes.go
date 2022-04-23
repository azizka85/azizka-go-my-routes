package home

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/azizka85/azizka-go-my-routes/settings"
	"github.com/azizka85/azizka-go-my-routes/templates"
	"github.com/azizka85/azizka-go-my-routes/templates/pages"
	"github.com/gorilla/mux"
)

func Default(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lang, ok := vars["lang"]

	if !ok {
		lang = settings.GlobalSettings.DefaultLanguage
	}

	language, ok := settings.GlobalSettings.Languages[lang]

	query := r.URL.Query()

	ajax := query.Get("ajax")
	init := query.Get("init")

	if ok {
		if ajax == "1" && init != "1" {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")

			data, _ := json.Marshal(&settings.GlobalSettings)

			fmt.Fprint(w, string(data))
		} else {
			layoutNames := []string{"main-layout"}

			if ajax == "1" {
				layoutNames = helpers.StringToArray(r.URL.Query().Get("layouts"))
			}

			w.Header().Set("Content-Type", "text/html;charset=UTF-8")

			fmt.Fprint(
				w,
				templates.RenderPage(
					ajax == "1",
					lang,
					pages.Home("luna"),
					layoutNames,
					query,
					&settings.GlobalSettings,
					language.Translator,
				),
			)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/", Default)
	router.HandleFunc("/{lang}", Default)
}
