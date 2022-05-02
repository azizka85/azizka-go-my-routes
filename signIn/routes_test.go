package signIn

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/azizka85/azizka-go-my-routes/global"
	"github.com/azizka85/azizka-go-my-routes/helpers"
	"github.com/azizka85/azizka-go-my-routes/mocks"
)

func TestDefaultWithDefaultLanguage(t *testing.T) {
	router, db, err := mocks.PrepareForTesting()

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	defer db.Close()

	r := httptest.NewRequest(
		http.MethodGet,
		global.Settings.PageRoot+"sign-in",
		nil,
	)
	w := httptest.NewRecorder()

	AddRoutes(router)

	router.ServeHTTP(w, r)

	res := w.Result()

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if res.StatusCode == http.StatusNotFound {
		t.Errorf(
			"page '/sign-in' couldn't find",
		)
	}

	if res.Header.Get("Content-Type") != "text/html;charset=UTF-8" {
		t.Errorf(
			"expected 'Content-Type' to be 'text/html;charset=UTF-8' got %v",
			w.Header().Get("Content-Type"),
		)
	}
}

func TestDefaultAjax(t *testing.T) {
	router, db, err := mocks.PrepareForTesting()

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	defer db.Close()

	r := httptest.NewRequest(
		http.MethodGet,
		global.Settings.PageRoot+"sign-in?ajax=1&init=1",
		nil,
	)
	w := httptest.NewRecorder()

	AddRoutes(router)

	router.ServeHTTP(w, r)

	res := w.Result()

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if res.StatusCode == http.StatusNotFound {
		t.Errorf(
			"page '/sign-in' couldn't find",
		)
	}

	if res.Header.Get("Content-Type") != "text/html;charset=UTF-8" {
		t.Errorf(
			"expected 'Content-Type' to be 'text/html;charset=UTF-8' got %v",
			w.Header().Get("Content-Type"),
		)
	}

	r = httptest.NewRequest(
		http.MethodGet,
		global.Settings.PageRoot+"sign-in?ajax=1",
		nil,
	)
	w = httptest.NewRecorder()

	AddRoutes(router)

	router.ServeHTTP(w, r)

	res = w.Result()

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if res.StatusCode == http.StatusNotFound {
		t.Errorf(
			"page '/sign-in' couldn't find",
		)
	}

	if res.Header.Get("Content-Type") != "application/json;charset=UTF-8" {
		t.Errorf(
			"expected 'Content-Type' to be 'application/json;charset=UTF-8' got %v",
			w.Header().Get("Content-Type"),
		)
	}
}

func TestPost(t *testing.T) {
	router, db, err := mocks.PrepareForTesting()

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	defer db.Close()

	r := httptest.NewRequest(
		http.MethodPost,
		global.Settings.PageRoot+"sign-in?ajax=1",
		nil,
	)

	user := &data.User{
		Email:    "test@mail.ru",
		Password: "lock",
	}

	r.ParseForm()

	r.PostForm.Set("email", user.Email)
	r.PostForm.Set("password", user.Password)

	w := httptest.NewRecorder()

	AddRoutes(router)

	_, err = helpers.CreateUser(user, global.Db)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	router.ServeHTTP(w, r)

	res := w.Result()

	defer res.Body.Close()

	cookies := res.Cookies()

	var sessionCookie *http.Cookie

	for _, cookie := range cookies {
		if cookie.Name == global.SessionKey {
			sessionCookie = cookie

			break
		}
	}

	if sessionCookie == nil {
		t.Errorf("expected 'session' to be set in cookie")
	}

	r.AddCookie(sessionCookie)

	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if res.StatusCode == http.StatusBadRequest {
		t.Errorf(
			"Error: %v",
			string(content),
		)
	}

	session, err := global.SessionStore.Get(r, global.SessionKey)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if _, ok := session.Values["userId"]; !ok {
		t.Errorf("expected session value 'userId' to be set")
	}

	if _, ok := session.Values["service"]; ok {
		t.Errorf("expected session value 'service' to be not set")
	}

	if _, ok := session.Values["oauth"]; ok {
		t.Errorf("expected session value 'oauth' to be not set")
	}
}
