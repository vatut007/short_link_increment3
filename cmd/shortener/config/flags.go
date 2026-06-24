package config

import "flag"

var flagRunPort int
var flagRunAddr string

func parseFlags() {
	flag.IntVar(&flagRunPort, "a", 8080, "port to run server")
	flag.StringVar(&flagRunAddr, "-b", "http://localhost:8000", "address to run server")
	flag.Parse()
}
