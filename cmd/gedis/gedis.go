package main

import (
	"log"

	"github.com/jacksonopp/gedis/internal/server"
)

var url = "localhost"
var port = ":8080"

func main() {
	gedis := server.NewGedisServer("8080")
	err := gedis.Start()
	if err != nil {
		log.Fatalln("failed to start gedis server", err)
	}

}
