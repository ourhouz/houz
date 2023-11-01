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
