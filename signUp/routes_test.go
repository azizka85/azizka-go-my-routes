package signUp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azizka85/azizka-go-my-routes/settings"
)

func TestDefaultWithDefaultLanguage(t *testing.T) {
	r := httptest.NewRequest(
		http.MethodGet,
		settings.GlobalSettings.PageRoot+"sign-up",
		nil,
	)
	w := httptest.NewRecorder()

	Default(w, r)

	res := w.Result()

	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)

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
	r := httptest.NewRequest(
		http.MethodGet,
		settings.GlobalSettings.PageRoot+"sign-up?ajax=1&init=1",
		nil,
	)
	w := httptest.NewRecorder()

	Default(w, r)

	res := w.Result()

	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)

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
		settings.GlobalSettings.PageRoot+"sign-up?ajax=1",
		nil,
	)
	w = httptest.NewRecorder()

	Default(w, r)

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
