package internal

import (
	"net/http"

	"github.com/judegiordano/sst_template/TEST/controllers/dev"
	"github.com/judegiordano/sst_template/TEST/middleware"
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
	// api.Handle("/dev/", http.StripPrefix("/dev", dev.Router()))
	api.Handle("/test1", dev.Router())
	api.Handle("/test2/", dev.Router())
	api.Handle("/test3/", http.StripPrefix("/test3", dev.Router()))
	// api.Handle("/admin/", admin.Router())
	// api.Handle("/admin/", http.StripPrefix("/admin", admin.Router()))
	return stack(api)
}
