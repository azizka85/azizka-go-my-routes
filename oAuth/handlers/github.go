package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/azizka85/azizka-go-my-routes/data"
)

type GitHub struct{}

const GitHubAuthorizeUrl = "https://github.com/login/oauth/authorize"
const GitHubAccessTokenUrl = "https://github.com/login/oauth/access_token"
const GitHubUserUrl = "https://api.github.com/user"

func (gitHub *GitHub) Handle(w http.ResponseWriter, r *http.Request) error {
	u, err := url.Parse(GitHubAuthorizeUrl)

	if err != nil {
		return err
	}

	q := url.Values{}

	q.Set("client_id", os.Getenv("GITHUB_CLIENT_ID"))

	rq := r.URL.Query()

	state := data.OAuthServiceRequestState{}

	state.Lang = rq.Get("lang")
	state.Ajax = rq.Get("ajax")

	stateStr, err := json.Marshal(&state)

	if err != nil {
		return err
	}

	q.Set("state", string(stateStr))

	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusFound)

	return nil
}

func (gitHub *GitHub) Callback(w http.ResponseWriter, r *http.Request) error {
	rq := r.URL.Query()

	/* var state data.OAuthServiceRequestState

	err := json.Unmarshal([]byte(rq.Get("state")), &state)

	if err != nil {
		return err
	}

	lang := state.Lang

	if lang == "" {
		lang = settings.GlobalSettings.DefaultLanguage
	}

	ajax := false

	if state.Ajax == "1" {
		ajax = true
	}

	language, ok := settings.GlobalSettings.Languages[lang]

	translator := &i18n.Translator{}

	if ok {
		translator = language.Translator
	} */

	params := map[string]string{
		"client_id":     os.Getenv("GITHUB_CLIENT_ID"),
		"client_secret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"code":          rq.Get("code"),
	}

	paramsStr, err := json.Marshal(params)

	if err != nil {
		return err
	}

	u, err := url.Parse(GitHubAccessTokenUrl)

	if err != nil {
		return err
	}

	client := http.Client{}

	req, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewBuffer(paramsStr),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprint(len(paramsStr)))
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	var accessRes map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&accessRes)

	if err != nil {
		return err
	}

	u, err = url.Parse(GitHubUserUrl)

	if err != nil {
		return err
	}

	client = http.Client{}

	req, err = http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Set("Authorization", "token "+fmt.Sprint(accessRes["access_token"]))

	resp, err = client.Do(req)

	if err != nil {
		return err
	}

	var userRes map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&userRes)

	if err != nil {
		return err
	}

	res := map[string]map[string]interface{}{
		"access": accessRes,
		"user":   userRes,
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	data, _ := json.Marshal(res)

	fmt.Fprint(w, string(data))

	return nil
}
