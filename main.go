package main

import (
	"go-clean-arch/configuration"
)

func main() {
	server := configuration.WireServer()
	server.Start()
}
