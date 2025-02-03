package admin

import (
	"encoding/json"
	"net/http"

	"github.com/judegiordano/sst_template/MUX/middleware"
)

type Ping struct {
	Admin bool `json:"admin"`
}

func me(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Ping{Admin: true})
}

func Router() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("GET /me", me)
	// middleware
	stack := middleware.Stack(middleware.Authenticate)
	return stack(r)
}
