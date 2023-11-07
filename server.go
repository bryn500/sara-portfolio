package main

import (
	"net/http"
	"saraweb/startup"
)

func Start(config startup.AppConfig) {
	http.Handle("/", http.FileServer(http.Dir(config.OutputDir)))
	http.ListenAndServe(":"+config.Port, nil)
}
