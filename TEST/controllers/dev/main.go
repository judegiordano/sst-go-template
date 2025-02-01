package dev

import (
	"net/http"

	"github.com/judegiordano/sst_template/TEST/responses"
)

type Ping struct {
	Ok bool `json:"ok"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	responses.Created(w, Ping{Ok: true})
}

func Router() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("GET /ping", ping)
	return r
}
