package signUp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azizka85/azizka-go-my-routes/global"
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
		global.Settings.PageRoot+"sign-up",
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
			"page '/sign-up' couldn't find",
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
		global.Settings.PageRoot+"sign-up?ajax=1&init=1",
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
			"page '/sign-up' couldn't find",
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
		global.Settings.PageRoot+"sign-up?ajax=1",
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
			"page '/sign-up' couldn't find",
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
		global.Settings.PageRoot+"sign-up?ajax=1",
		nil,
	)

	r.ParseForm()

	r.PostForm.Set("email", "test@mail.ru")
	r.PostForm.Set("password", "lock")
	r.PostForm.Set("fullName", "Test")

	w := httptest.NewRecorder()

	AddRoutes(router)

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
