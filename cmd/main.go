package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":1400",
		db:   dbConfig{}}

	api := application{
		config: cfg}

	if err := api.run(api.mount()); err != nil {
		log.Printf("error starting server: %v", err)
		os.Exit(1)
	}
}
