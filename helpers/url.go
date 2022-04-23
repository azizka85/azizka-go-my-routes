package helpers

import "net/url"

func ToggleQuery(v url.Values, key string) string {
	if v.Has(key) {
		v.Del(key)
	} else {
		v.Set(key, "1")
	}

	return v.Encode()
}
