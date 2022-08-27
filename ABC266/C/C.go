// this code is not complete

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type V struct {
	x int
	y int
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [4]string
	var split_input []string
	var A V
	var B V
	var C V
	var D V
	var tmp float64

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	A[0], _ = strconv.ParseFloat(split_input[0], 64)
	A[1], _ = strconv.ParseFloat(split_input[1], 64)
	sc.Scan()
	input[1] = sc.Text()
	split_input = strings.Split(input[1], " ")
	B[0], _ = strconv.ParseFloat(split_input[0], 64)
	B[1], _ = strconv.ParseFloat(split_input[1], 64)
	sc.Scan()
	input[2] = sc.Text()
	split_input = strings.Split(input[2], " ")
	C[0], _ = strconv.ParseFloat(split_input[0], 64)
	C[1], _ = strconv.ParseFloat(split_input[1], 64)
	sc.Scan()
	input[3] = sc.Text()
	split_input = strings.Split(input[3], " ")
	D[0], _ = strconv.ParseFloat(split_input[0], 64)
	D[1], _ = strconv.ParseFloat(split_input[1], 64)

	tmp = A[0]*C[1] - A[1]*C[0]
	if tmp < 0 {
		fmt.Println("No")
		return
	}
	tmp = B[0]*D[1] - B[1]*D[0]
	if tmp < 0 {
		fmt.Println("No")
		return
	}
	tmp = C[0]*A[1] - C[1]*A[0]
	if tmp < 0 {
		fmt.Println("No")
		return
	}
	tmp = D[0]*B[1] - D[1]*B[0]
	if tmp < 0 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")

}
