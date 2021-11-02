package main

import (
	service "github.com/rodkevich/gmp-tickets/internal/server"
	"log"
)

const Version = "v0.1.0"

// localhost:12300 as default
func main() {
	apiServer := service.NewServer(Version)
	apiServer.Initialize()

	log.Println(apiServer.GetConfig())

	if err := apiServer.Run(); err != nil {
		log.Fatalf("ApiServer:%v", err.Error())
	}
}
