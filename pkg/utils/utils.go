package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func CreateJson(w http.ResponseWriter, statusCode int, msg string, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(map[string]any{"message": msg, "data": data})
}
