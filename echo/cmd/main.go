package main

import (
	"log"

	server "github.com/totoyk/echo/internal"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
