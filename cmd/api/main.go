package main

import (
	"log"

	"github.com/buildwithbeejay/Golang-Backend-Api/internals/env"
)

func main() {
	cfg := config{
		//WE need to install direnv to load the env file   
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	r := app.Mount()
	log.Fatal(app.run(r))
}