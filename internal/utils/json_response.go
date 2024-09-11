package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONAnyResponse(w http.ResponseWriter, statusCode int, resp any) error {
	w.Header().Add("Content-Tyep", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(resp)
}

func JSONErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	JSONAnyResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

func JSONMessageResponse(w http.ResponseWriter, statusCode int, message string) {
	JSONAnyResponse(w, statusCode, struct {
		Message string `json:"message"`
	}{
		message,
	})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}
