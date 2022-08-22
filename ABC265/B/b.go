package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [100010]string
	var N int
	var M int
	var T int
	var A [100010]int
	var X int
	var Y [100010]int
	var i int

	sc.Split(bufio.ScanWords)
	sc.Scan()
	input[0] = sc.Text()
	sc.Scan()
	input[1] = sc.Text()
	sc.Scan()
	input[2] = sc.Text()
	// fmt.Println(input[0])
	// fmt.Println(input[1])
	// fmt.Println(input[2])
	N, _ = strconv.Atoi(input[0])
	M, _ = strconv.Atoi(input[1])
	T, _ = strconv.Atoi(input[2])

	for i = 1; i < N; i++ {
		sc.Scan()
		input[i+2] = sc.Text()
		// fmt.Println(input[i+2])
		A[i], _ = strconv.Atoi(input[i+2])
		//fmt.Println(A)
	}

	for i = 0; i < M; i++ {
		sc.Scan()
		input[0] = sc.Text()
		sc.Scan()
		input[1] = sc.Text()
		// fmt.Println(input[i])
		// fmt.Println(split_input)
		X, _ = strconv.Atoi(input[0])
		Y[X], _ = strconv.Atoi(input[1])
	}

	for i = 1; i < N; i++ {
		if T <= A[i] {
			fmt.Println("No")
			return
		}
		T -= A[i]
		T += Y[i+1]
	}
	fmt.Println("Yes")
}
