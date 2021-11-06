package main

import (
	"fmt"
	"strings"
)

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func check(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
	}
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
