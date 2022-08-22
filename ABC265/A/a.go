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
	var input [103]string
	var split_input []string
	var X int
	var Y int
	var N int
	var cash int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	X, _ = strconv.Atoi(split_input[0])
	Y, _ = strconv.Atoi(split_input[1])
	N, _ = strconv.Atoi(split_input[2])

	if (X * 3) > Y {
		for N >= 3 {
			cash = cash + Y
			N = N - 3
		}
		for N >= 1 {
			cash = cash + X
			N = N - 1
		}
	} else {
		for N >= 1 {
			cash = cash + X
			N = N - 1
		}
	}

	fmt.Println(cash)

}
