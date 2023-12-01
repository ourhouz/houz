package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

// parseBody parses the JSON body of an HTTP request into the given type
func parseBody[T any](r *http.Request) (T, error) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		var zero T
		return zero, err
	}

	var b T
	if err = json.Unmarshal(raw, &b); err != nil {
		var zero T
		return zero, err
	}

	return b, nil
}

// writeJson marshals data as JSON and sends the bytes as the response body
func writeJson(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
