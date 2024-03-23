package main

import (
	"fmt"
	"log"
	"os"

	_ "gabriel/search-in-feed/matchers"
	"gabriel/search-in-feed/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <argument>")
		return
	}
	argument := os.Args[1]

	search.Run(argument)
}
