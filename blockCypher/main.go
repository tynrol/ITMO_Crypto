package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		help()
	} else if len(args) != 5 {
		fmt.Println("Invalid number of command line arguments")
		help()
	} else {
		key, err := os.ReadFile(args[2])
		input, err := os.ReadFile(args[3])
		_, err = os.Stat(args[4])
		var result []byte
		check(err)
		if strings.Compare(args[1], "encrypt") == 0 {
			result, err = encrypt(input, key)
			check(err)
		} else if strings.Compare(args[1], "decrypt") == 0 {
			result, err = decrypt(input, key)
			check(err)
		} else {
			help()
		}
		err = os.WriteFile(args[4], result, 0644)
		check(err)
	}
}
