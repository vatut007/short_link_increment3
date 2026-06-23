package main

import (
	"log"
	"net/http"
	"shortener/handlers"
)

func main() {
	handlers.RegisterHandlers()
	if err := http.ListenAndServe(":8080", handlers.Main_router); err != nil {
		log.Fatal(err)
	}
}
