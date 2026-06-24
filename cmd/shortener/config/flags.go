package config

import "flag"

func ParseFlags() {
	flag.StringVar(&ConfigApp.ServerAddress, "a", "localhost:8080", "HTTP server address (e.g. localhost:8888)")
	flag.StringVar(&ConfigApp.BaseURL, "b", "http://localhost:8080", "Base URL for short links (e.g. http://localhost:8000)")
	flag.Parse()
}
