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
	var input [100000]string
	var split_input []string
	var N int
	// var X int
	// var Y int
	// var U []int
	// var V []int
	// var ans []int
	// var pre int
	// var pre_len int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(split_input[0])
	// X, _ = strconv.Atoi(split_input[1])
	// Y, _ = strconv.Atoi(split_input[2])

	matrix := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		matrix[i] = make([]int, N+1)
	}

	for i := 0; i < N-1; i++ {
		sc.Scan()
		input[0] = sc.Text()
		split_input = strings.Split(input[0], " ")
		a, _ := strconv.Atoi(split_input[0])
		b, _ := strconv.Atoi(split_input[1])
		matrix[a] = append([]int{b}, matrix[a][0:]...)
		matrix[b] = append([]int{a}, matrix[b][0:]...)
	}
	fmt.Println(matrix)

}
