package main

import (
	"kirby/api"
)

func main() {
	server := api.NewServer()
	server.Start()
}
