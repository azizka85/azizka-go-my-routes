package helpers

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func Clone(v url.Values) url.Values {
	clone := url.Values{}

	for itemKey, itemValue := range v {
		clone[itemKey] = itemValue
	}

	return clone
}

func ToggleQuery(v url.Values, key string) string {
	clone := Clone(v)

	if clone.Has(key) {
		clone.Del(key)
	} else {
		clone.Set(key, "1")
	}

	return clone.Encode()
}

func ChangeRoutePath(request *http.Request, query url.Values, pairs ...string) string {
	route := mux.CurrentRoute(request)
	vars := mux.Vars(request)
	merge := FlatToMap(pairs)
	flat := MapToFlat(vars, merge)

	url, err := route.URL(flat...)

	queryString := query.Encode()

	if queryString != "" {
		queryString = "?" + queryString
	}

	if err != nil {
		return queryString
	} else {
		return url.String() + queryString
	}
}
