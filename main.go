package main

import (
	"log"

	"github.com/hacdan/origin-backend-take-home-assignment/api"
)

func main() {
	listenAddr := ":8080"
	server := api.NewServer(listenAddr)
	log.Fatal(server.Start())
}
