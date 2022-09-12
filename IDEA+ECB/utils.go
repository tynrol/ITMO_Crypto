package main

import (
	"fmt"
	"log"
)

func help() {
	fmt.Println("Usage: go run . encrypt/decrypt keyFile inputFile outputFile")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
