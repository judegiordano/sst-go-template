package responses

import (
	"encoding/json"
	"net/http"
)

func Ok(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Created(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func InternalServerError(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data)
}
