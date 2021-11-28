package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// var letters = []rune("abcdefghijklmnopqrstuvwxyz")
// var letters = []rune(" 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+,-—./[]^_:;<=>?@{|}")
var letters = []rune(" 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+,-—./[]^_:;<=>?@{|}~абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")

var m = len(letters)
var a = 5
var b = 7

func encrypt(input []rune) []rune {
	size := len(input)
	encrypted := make([]rune, size)

	for i := 0; i < size; i++ {
		for j := 0; j < m; j++ {
			if letters[j] == input[i] {
				new_index := (a*j + b) % m
				encrypted[i] = letters[new_index]
				break
			}
		}
	}
	return encrypted
}

func decrypt(input []rune) []rune {
	size := len(input)
	decrypted := make([]rune, size)
	neg_a := solve(a, m)

	for i := 0; i < size; i++ {
		for j := 0; j < m; j++ {
			if letters[j] == input[i] {
				new_index := neg_a * (j + m - b) % m
				decrypted[i] = letters[new_index]
				break
			}
		}
	}
	return decrypted
}

func solve(a int, size int) int {
	for i := 0; ; i++ {
		value := float64((1 + size*i)) / float64(a)
		if math.Mod(value, 1.0) == 0 {
			return int(value)
		}
	}
}

func analysis(input []rune) {
	runeMap := make(map[rune]int)
	for i := 0; i < len(input); i++ {
		runeMap[input[i]] = runeMap[input[i]] + 1
	}
	for k, v := range runeMap {
		fmt.Printf("letter %c appears %d times\n", k, v)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	help()

	for {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		text = trim(text)

		if strings.Compare("analysis", text) == 0 {
			fmt.Print("enter filename: \n")
			fmt.Print(">")
			filename, _ := reader.ReadString('\n')
			filename = trim(filename)

			file, err := os.ReadFile(filename)
			check(err)

			runes := []rune(string(file))
			analysis(runes)
		} else if strings.Compare("encrypt", text) == 0 {
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
