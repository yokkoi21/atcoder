package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [1]string
	var split_input []string
	var A int
	var B int
	var C int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	A, _ = strconv.Atoi(split_input[0])
	B, _ = strconv.Atoi(split_input[1])

	if A-4 >= 0 || B-4 >= 0 {
		C += 4
		if A-4 >= 0 {
			A -= 4
		}
		if B-4 >= 0 {
			B -= 4
		}
	}
	if A-2 >= 0 || B-2 >= 0 {
		C += 2
		if A-2 >= 0 {
			A -= 2
		}
		if B-2 >= 0 {
			B -= 2
		}
	}
	if A-1 >= 0 || B-1 >= 0 {
		C += 1
		if A-1 >= 0 {
			A -= 1
		}
		if B-1 >= 0 {
			B -= 1
		}
	}

	fmt.Println(C)
}
