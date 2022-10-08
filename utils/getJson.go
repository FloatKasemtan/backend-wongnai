package utils

import (
	"encoding/json"
	"net/http"
)

func GetJson[T any](url string, target T) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
