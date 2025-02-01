package main

import (
	"net/http"

	"github.com/judegiordano/sst_template/TEST/internal"

	"github.com/judegiordano/gogetem/pkg/logger"
)

func main() {
	handler := internal.Server()
	logger.Info("[API]", "running on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", handler)
}
