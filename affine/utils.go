package main

import (
	"fmt"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func trim(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	str = strings.Trim(str, " ")
	return str
}

func help() {
	fmt.Printf("[encrypt|decrypt|analysis]\n")
}
