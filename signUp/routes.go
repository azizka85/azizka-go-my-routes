package signUp

import (
	"encoding/json"
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
	session, err := global.SessionStore.Get(r, global.SessionKey)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		vars := mux.Vars(r)

		lang, ok := vars["lang"]

		if !ok {
			lang = global.Settings.DefaultLanguage
		}

		language, ok := global.Settings.Languages[lang]

		query := r.URL.Query()

		ajax := query.Get("ajax")
		init := query.Get("init")

		if ok {
			if ajax == "1" && init != "1" {
				w.Header().Set("Content-Type", "application/json;charset=UTF-8")

				data, _ := json.Marshal(&global.Settings)

				fmt.Fprint(w, string(data))
			} else {
				user := &data.User{}

				count, err := helpers.GetUserInfoFromSession(
					user,
					session,
					global.Db,
				)

				if err != nil {
					fmt.Println(err)
				} else if count == 0 {
					user = nil
				}

				w.Header().Set("Content-Type", "text/html;charset=UTF-8")

				fmt.Fprint(
					w,
					templates.RenderPage(
						ajax == "1",
						lang,
						pages.SignUp(
							global.Settings.PageRoot,
							lang,
							components.AuthService(
								lang,
								global.Settings.PageRoot,
								language.Translator,
							),
							language.Translator,
						),
						[]string{},
						query,
						r,
						user,
						&global.Settings,
						language.Translator,
					),
				)
			}
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
}

func postData(user *data.User, w http.ResponseWriter, r *http.Request) (int, error) {
	session, err := global.SessionStore.Get(r, global.SessionKey)

	if err != nil {
		fmt.Println(err)

		return http.StatusInternalServerError, err
	} else {
		vars := mux.Vars(r)

		lang, ok := vars["lang"]

		if !ok {
			lang = global.Settings.DefaultLanguage
		}

		language, ok := global.Settings.Languages[lang]

		if ok {
			decoder := schema.NewDecoder()

			r.ParseForm()

			err := decoder.Decode(user, r.PostForm)

			if err != nil {
				fmt.Println(err)

				return http.StatusBadRequest, err
			} else {
				err = helpers.SignUp(
					user,
					language.Translator,
					session,
					global.Db,
				)

				if err != nil {
					fmt.Println(err)

					return http.StatusInternalServerError, err
				} else {
					err = session.Save(r, w)

					if err != nil {
						return http.StatusInternalServerError, err
					} else {
						return http.StatusOK, nil
					}
				}
			}
		} else {
			return http.StatusNotFound, nil
		}
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var user data.User

	query := r.URL.Query()

	ajax := query.Get("ajax")

	status, err := postData(&user, w, r)

	if status != http.StatusOK {
		var errorTxt string

		if err == nil {
			errorTxt = http.StatusText(status)
		} else {
			errorTxt = err.Error()
		}

		if ajax == "1" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")

			http.Error(w, errorTxt, status)
		} else {
			// TODO: not ajax

			w.Header().Set("Content-Type", "text/html;charset=UTF-8")
		}
	} else {
		if ajax == "1" {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")

			data, _ := json.Marshal(&user)

			fmt.Fprint(w, string(data))
		} else {
			// TODO: not ajax

			w.Header().Set("Content-Type", "text/html;charset=UTF-8")
		}
	}
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
