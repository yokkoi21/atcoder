package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [1]string
	var split_input []string
	var half int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], "")

	half = (len(split_input) - 1) / 2

	fmt.Println(split_input[half])

}
