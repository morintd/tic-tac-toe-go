package main

import (
	"flag"

	"tic-tac-toe/internal/core"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	server := core.NewServer()
	server.Start(*addr)
}
