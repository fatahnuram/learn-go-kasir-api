package helpers

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, r *http.Request, status int, msg any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msg)
}
