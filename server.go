package main

import (
	"net/http"
)

func StartServer(config AppConfig) {
	http.Handle("/", http.FileServer(http.Dir(config.OutputDir)))
	http.ListenAndServe(":"+config.Port, nil)
}
