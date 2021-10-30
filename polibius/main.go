package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func randomize() {
	letterRunes = []rune(" 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+,-—./[]^_:;<=>?@{|}~абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	saveRunes := ""
	fmt.Println("current polibus square: ")
	for i := 0; i < matrix_size; i++ {
		for j := 0; j < matrix_size; j++ {
			if len(letterRunes) == 0 {
				polibius[i][j] = []rune("&")[0]
			} else {
				randValue := rand.Intn(len(letterRunes))
				polibius[i][j] = letterRunes[randValue]
				saveRunes = saveRunes + string(letterRunes[randValue])
				letterRunes = delChar(letterRunes, randValue)
			}
			fmt.Print(string(polibius[i][j]), " ")
		}
		fmt.Println()
	}

}

func encrypt(msg string) string {
	var encrypted string
	var size int
	r := bufio.NewReader(strings.NewReader(msg))
	letters := make([]rune, len(msg))

	//определяем количество символов в файле и записываем их как руны
	//может есть более деликатный способ, но этот зато рабочий
	for i := 0; ; i++ {
		c, _, err := r.ReadRune()
		if err != nil {
			size = i
			break
		}
		letters[i] = c
	}

	x := make([]int, size)
	y := make([]int, size)
	z := make([]int, 2*size)

	//ищем позицию буквы в квадрате полибия, соответсвующую букве сообщения
	//когда такая буква находится, то мы запоминаем ее позицию в соответвующих массивах
	for k := 0; k < size; k++ {
		for i := 0; i < matrix_size; i++ {
			for j := 0; j < matrix_size; j++ {
				if polibius[i][j] == letters[k] {
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

	for k := 0; k < size; k++ {
		// fmt.Printf("encrypted: " + string(polibius[y[k]][x[k]]) + " k: ")
		encrypted = encrypted + string(polibius[y[k]][x[k]])
	}
	return encrypted
}

func decrypt(msg string) string {
	var decrypted string
	var size int

	r := bufio.NewReader(strings.NewReader(msg))
	letters := make([]rune, len(msg))

	//определяем количество символов в файле и записываем их как руны
	//может есть более деликатный способ, но этот зато рабочий
	for i := 0; ; i++ {
		c, _, err := r.ReadRune()
		if err != nil {
			size = i
			break
		}
		letters[i] = c
	}

	x := make([]int, size)
	y := make([]int, size)
	z := make([]int, 2*size)

	//считываем полученные координаты построчно
	for k := 0; k < size; k++ {
		for i := 0; i < matrix_size; i++ {
			for j := 0; j < matrix_size; j++ {
				if polibius[i][j] == letters[k] {
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
		decrypted = decrypted + string(polibius[y[k]][x[k]])
	}
	return decrypted
}

func help() {
	fmt.Printf("[encrypt|decrypt|randomize]\n")
}

func trim(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	str = strings.Trim(str, " ")
	return str
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
			// filename, _ := reader.ReadString('\n')
			// filename = trim(filename)
			filename := "e_rus_msg"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			decrypted := decrypt(string(content))
			err = ioutil.WriteFile("d_"+filename, []byte(decrypted), 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("file content decrypted")
		} else if strings.Compare("encrypt", text) == 0 {
			fmt.Print("enter filename: \n")
			fmt.Print(">")
			// filename, _ := reader.ReadString('\n')
			// filename = trim(filename)
			filename := "rus_msg"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			encrypted := encrypt(string(content))
			err = ioutil.WriteFile("e_"+filename, []byte(encrypted), 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("file content encrypted")
		} else if strings.Compare("exit", text) == 0 || strings.Compare("quit", text) == 0 || strings.Compare("q", text) == 0 {
			os.Exit(0)
		} else {
			help()
		}
	}
}
