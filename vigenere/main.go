package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var letters = []rune(" 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+,-—./[]^_:;<=>?@{|}~абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
var n = len(letters)
var key []rune

func encrypt(input []rune) []rune {
	size := len(input)
	encrypted := make([]rune, size)
	var input_index, key_index int

	for len(key) < size {
		key = append(key, key...)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < n; j++ {
			if letters[j] == input[i] {
				input_index = j
			}
			if letters[j] == key[i] {
				key_index = j
			}
		}
		encrypted[i] = letters[(input_index+key_index)%n]
	}
	return encrypted
}

func decrypt(input []rune) []rune {
	size := len(input)
	decrypted := make([]rune, size)
	var input_index, key_index int

	for len(key) < size {
		key = append(key, key...)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < n; j++ {
			if letters[j] == input[i] {
				input_index = j
			}
			if letters[j] == key[i] {
				key_index = j
			}
		}
		decrypted[i] = letters[(input_index+n-key_index)%n]
	}
	return decrypted
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter key\n")
	fmt.Print(">")
	text, _ := reader.ReadString('\n')
	key = []rune(trim(text))
	help()

	for {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		text = trim(text)

		if strings.Compare("encrypt", text) == 0 {
			fmt.Print("enter filename: \n")
			fmt.Print(">")
			filename, _ := reader.ReadString('\n')
			filename = trim(filename)

			file, err := os.ReadFile(filename)
			check(err)

			runes := []rune(string(file))
			encrypted := encrypt(runes)

			bytes := []byte(string(encrypted))
			err = os.WriteFile("e_"+filename, bytes, 0644)
			check(err)

			fmt.Println("file content encrypted")
		} else if strings.Compare("decrypt", text) == 0 {
			fmt.Print("enter filename: \n")
			fmt.Print(">")
			filename, _ := reader.ReadString('\n')
			filename = trim(filename)

			file, err := os.ReadFile(filename)
			check(err)

			runes := []rune(string(file))
			decrypted := decrypt(runes)

			bytes := []byte(string(decrypted))
			err = os.WriteFile("d_"+filename, bytes, 0644)
			check(err)

			fmt.Println("file content decrypted")
		} else if strings.Compare("exit", text) == 0 || strings.Compare("quit", text) == 0 || strings.Compare("q", text) == 0 {
			os.Exit(0)
		} else {
			help()
		}
	}
}
