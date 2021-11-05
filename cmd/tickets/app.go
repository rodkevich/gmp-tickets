package main

import (
	"github.com/rodkevich/gmp-tickets/internal/server/http"
	"log"
)

const Version = "v0.1.0"

// localhost:12300 as default
func main() {
	apiServer := http.NewServer(Version)
	apiServer.Initialize()

	// log.Println(apiServer.GetConfig())

	if err := apiServer.Run(); err != nil {
		log.Fatalf("ApiServer:%v", err.Error())
	}
}
