package internal

import (
	"net/http"

	"github.com/judegiordano/sst_template/MUX/controllers/admin"
	"github.com/judegiordano/sst_template/MUX/controllers/dev"
	"github.com/judegiordano/sst_template/MUX/middleware"
)

func Server() http.Handler {
	api := http.NewServeMux()
	// middleware
	stack := middleware.Stack(
		middleware.ErrorHandler,
		middleware.TransformJson,
		middleware.Logger,
	)
	// routes
	api.Handle("/dev/", http.StripPrefix("/dev", dev.Router()))
	api.Handle("/admin/", http.StripPrefix("/admin", admin.Router()))
	return stack(api)
}
