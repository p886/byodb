package main

import (
	"fmt"

	"github.com/p886/byo-database/repl"
)

const logFile = "data.log"

func main() {
	fmt.Println("Welcome! Enter command prefixed with PUT to store, GET to retrieve.")

	repl.Loop(logFile)
}
