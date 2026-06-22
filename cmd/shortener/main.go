package main

import (
	"net/http"
	"shortener/handlers"
)

func main() {
	handlers.RegisterHandlers()
	http.ListenAndServe(":8080", handlers.Main_router)
}
