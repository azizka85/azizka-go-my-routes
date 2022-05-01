package signIn

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
