package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func main() {
	file, err := os.ReadFile("read.txt")
	check(err)

	runes := []rune(string(file))
	for i := 0; i < len(runes); i++ {
		fmt.Printf("Char %c Unicode: %U, Rune pos: %d\n", runes[i], runes[i], i)
	}

	bytes := []byte(string(runes))
	err = os.WriteFile("write.txt", bytes, 0644)
	check(err)
}
