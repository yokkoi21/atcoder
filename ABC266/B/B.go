package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [1]string
	var N int
	var tmp int

	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])
	tmp = N % 998244353
	//fmt.Println(tmp)
	if tmp < 0 {
		tmp += 998244353
	}
	fmt.Println(tmp)
}
