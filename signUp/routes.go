package signUp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/azizka85/azizka-go-my-routes/settings"
	"github.com/azizka85/azizka-go-my-routes/templates"
	"github.com/azizka85/azizka-go-my-routes/templates/components"
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
		if r.Method == http.MethodPost {

		} else {
			if ajax == "1" && init != "1" {
				w.Header().Set("Content-Type", "application/json;charset=UTF-8")

				data, _ := json.Marshal(&settings.GlobalSettings)

				fmt.Fprint(w, string(data))
			} else {
				w.Header().Set("Content-Type", "text/html;charset=UTF-8")

				fmt.Fprint(
					w,
					templates.RenderPage(
						ajax == "1",
						lang,
						pages.SignUp(
							settings.GlobalSettings.PageRoot,
							lang,
							components.AuthService(
								lang,
								settings.GlobalSettings.PageRoot,
								language.Translator,
							),
							language.Translator,
						),
						[]string{},
						query,
						&settings.GlobalSettings,
						language.Translator,
					),
				)
			}
		}
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func AddRoutes(router *mux.Router) {
	router.HandleFunc("/sign-up", Default)
	router.HandleFunc("/{lang}/sign-up", Default)
}
