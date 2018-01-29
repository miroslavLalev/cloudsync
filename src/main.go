package main

import (
	"github.com/miroslavLalev/cloudsync/src/server"
)

func main() {
	server.GetRestServer().Start()
}
