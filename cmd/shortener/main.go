package main

import (
	"fmt"
	"log"
	"net/http"
	"shortener/cmd/shortener/config"
	"shortener/handlers"
)

func main() {
	handlers.RegisterHandlers()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.ConfigApp.Port), handlers.Main_router); err != nil {
		log.Fatal(err)
	}
}
