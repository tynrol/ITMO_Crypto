package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const matrix_size = 13

var polibius = [matrix_size][matrix_size]rune{}
var letterRunes = make([]rune, 150)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomize() {
	// letterRunes = []rune(" 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+,-—./[]^_:;<=>?@{|}")
	letterRunes = []rune(" 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+,-—./[]^_:;<=>?@{|}~абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	letterRunes = append(letterRunes, 0xa, 0xb, 0xc, 0xd)
	saveRunes := ""
	fmt.Println("current polibus square: ")
	for i := 0; i < matrix_size; i++ {
		for j := 0; j < matrix_size; j++ {
			if len(letterRunes) == 0 {
				polibius[i][j] = rune(12345 + matrix_size*i + j)
			} else {
				randValue := rand.Intn(len(letterRunes))
				polibius[i][j] = letterRunes[randValue]
				saveRunes = saveRunes + string(letterRunes[randValue])
				letterRunes = delChar(letterRunes, randValue)
			}
			if polibius[i][j] <= 0xd && polibius[i][j] >= 0xa {
				fmt.Print("  ")
			} else {
				fmt.Print(string(polibius[i][j]), " ")
			}
		}
		fmt.Println()
	}

}

func encrypt(runes []rune) []rune {
	var encrypted []rune

	size := len(runes)

	x := make([]int, size)
	y := make([]int, size)
	z := make([]int, 2*size)

	//ищем позицию буквы в квадрате полибия, соответсвующую букве сообщения
	//когда такая буква находится, то мы запоминаем ее позицию в соответвующих массивах
	for k := 0; k < size; k++ {
		for i := 0; i < matrix_size; i++ {
			for j := 0; j < matrix_size; j++ {
				if polibius[i][j] == runes[k] {
					x[k] = j
					y[k] = i
				}
			}
		}
	}

	//записываем позицию описанную в двух массивах подряд в одномерный массив(строку)
	for k := 0; k < size; k++ {
		z[k] = x[k]
	}
	for k := size; k < 2*size; k++ {
		z[k] = y[k-size]
	}

	//считываем полученные координаты построчно, получая новые координаты и соответственно новые буквы
	for k := 0; k < size; k++ {
		x[k] = z[k*2]
		y[k] = z[k*2+1]
	}

	//находим соответсвующие буквы по координатам в массивах x, y
	for k := 0; k < size; k++ {
		encrypted = append(encrypted, polibius[y[k]][x[k]])
	}
	return encrypted
}

func decrypt(runes []rune) []rune {
	var decrypted []rune

	size := len(runes)

	x := make([]int, size)
	y := make([]int, size)
	z := make([]int, 2*size)

	//считываем полученные координаты построчно
	for k := 0; k < size; k++ {
		for i := 0; i < matrix_size; i++ {
			for j := 0; j < matrix_size; j++ {
				if polibius[i][j] == runes[k] {
					z[2*k] = j
					z[2*k+1] = i
				}
			}
		}
	}

	//записываем их в соответсвующие одномерные массивы
	for k := 0; k < size; k++ {
		x[k] = z[k]
	}
	for k := size; k < 2*size; k++ {
		y[k-size] = z[k]
	}

	//находим соответсвующие буквы по координатам в массивах x, y
	for k := 0; k < size; k++ {
		decrypted = append(decrypted, polibius[y[k]][x[k]])
	}
	return decrypted
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	randomize()
	help()

	for {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		text = trim(text)

		if strings.Compare("randomize", text) == 0 {
			randomize()
		} else if strings.Compare("decrypt", text) == 0 {
			fmt.Print("enter filename: \n")
			fmt.Print(">")
			filename, _ := reader.ReadString('\n')
			// filename := "e_read.txt"
			filename = trim(filename)
			file, err := os.ReadFile(filename)
			check(err)

			runes := []rune(string(file))
			decrypted := decrypt(runes)

			err = os.WriteFile("d_"+filename, []byte(string(decrypted)), 0644)
			check(err)

			fmt.Println("file content decrypted")
		} else if strings.Compare("encrypt", text) == 0 {
			fmt.Print("enter filename: \n")
			fmt.Print(">")
			filename, _ := reader.ReadString('\n')
			// filename := "read.txt"
			filename = trim(filename)
			file, err := os.ReadFile(filename)
			check(err)

			runes := []rune(string(file))
			encrypted := encrypt(runes)

			err = os.WriteFile("e_"+filename, []byte(string(encrypted)), 0644)
			check(err)

			fmt.Println("file content encrypted")
		} else if strings.Compare("exit", text) == 0 || strings.Compare("quit", text) == 0 || strings.Compare("q", text) == 0 {
			os.Exit(0)
		} else {
			help()
		}
	}
}
