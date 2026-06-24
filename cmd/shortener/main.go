package main

import (
	"log"
	"net/http"
	"shortener/cmd/shortener/config"
	"shortener/handlers"
)

func main() {
	config.ParseFlags()
	handlers.RegisterHandlers()
	if err := http.ListenAndServe(config.ConfigApp.ServerAddress, handlers.Main_router); err != nil {
		log.Fatal(err)
	}
}
