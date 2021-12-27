package main

import (
	"encoding/hex"
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
		keyByte, err := os.ReadFile(args[2])
		check(err)
		input, err := os.ReadFile(args[3])
		check(err)
		_, err = os.Stat(args[4])
		check(err)

		var idea ideaCipher
		var result = make([]byte, len(input))
		k, _ := hex.DecodeString(string(keyByte))
		err = cipherInit(idea, k)
		check(err)

		if strings.Compare(args[1], "encrypt") == 0 {
			cryptBlock(idea.encryptKey[:], input, result[:])
			check(err)
			fmt.Println("File content encrypted")
		} else if strings.Compare(args[1], "decrypt") == 0 {
			cryptBlock(idea.decryptKey[:], input, result[:])
			check(err)
			fmt.Println("File content decrypted")
		} else {
			help()
		}
		err = os.WriteFile(args[4], result[:], 0644)
		check(err)
		//{"30303030303030303030303030303030", "3030303030303030", "4EE30E9A0DF346B7"}
	}
}
