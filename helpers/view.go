package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func ViewData(w http.ResponseWriter, r *http.Request) (
	lang string,
	ajax string,
	init string,
	query url.Values,
	session *sessions.Session,
	user *data.User,
	translator *i18n.Translator,
	status int,
	err error,
) {
	session, err = global.SessionStore.Get(r, global.SessionKey)

	if err != nil {
		fmt.Println(err)

		status = http.StatusInternalServerError

		return
	} else {
		vars := mux.Vars(r)

		var ok bool

		lang, ok = vars["lang"]

		if !ok {
			lang = global.Settings.DefaultLanguage
		}

		language, ok := global.Settings.Languages[lang]

		query = r.URL.Query()

		ajax = query.Get("ajax")
		init = query.Get("init")

		if ok {
			translator = language.Translator

			if ajax != "1" || init == "1" {
				user = &data.User{}

				var count int64

				count, err = GetUserInfoFromSession(
					user,
					session,
					global.Db,
				)

				if err != nil || count == 0 {
					user = nil
				}
			}
		} else {
			status = http.StatusNotFound

			return
		}
	}

	status = http.StatusOK

	return
}

func RenderDataOrRedirect(
	w http.ResponseWriter,
	r *http.Request,
	data interface{},
	content string,
	ajax string,
	init string,
	location string,
	status int,
	err error,
) {
	if status != http.StatusOK && status != http.StatusFound {
		errorTxt := http.StatusText(status)

		if err != nil {
			errorTxt = err.Error()
		}

		if ajax == "1" {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")

			http.Error(w, errorTxt, status)
		} else {
			if status == http.StatusInternalServerError {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")

				http.Error(w, errorTxt, status)
			} else {
				w.Header().Set("Content-Type", "text/html;charset=UTF-8")
				// TODO: Render for user error
				http.Error(w, errorTxt, status)
			}
		}
	} else {
		if status != http.StatusFound {
			if ajax == "1" && init != "1" {
				w.Header().Set("Content-Type", "application/json;charset=UTF-8")

				data, _ := json.Marshal(data)

				fmt.Fprint(w, string(data))
			} else {
				w.Header().Set("Content-Type", "text/html;charset=UTF-8")

				fmt.Fprint(
					w,
					content,
				)
			}
		} else {
			if ajax == "1" {
				http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
			} else {
				http.Redirect(w, r, location, status)
			}
		}
	}
}
