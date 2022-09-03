package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [200002]string
	var N int
	var M int
	var tmp_s string
	var sum [200001]int
	var sumi [200001]int
	var a [200001]int
	var i int
	var s int
	var t int
	var max int = -1000000000000000000

	sc.Split(bufio.ScanWords)
	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])
	sc.Scan()
	input[1] = sc.Text()
	M, _ = strconv.Atoi(input[1])
	//fmt.Println(split_input)
	//fmt.Println(split_input)
	for i = 0; i < N; i++ {
		sc.Scan()
		tmp_s = sc.Text()
		a[i], _ = strconv.Atoi(tmp_s)
	}
	for i = 1; i <= N; i++ {
		sum[i] = sum[i-1] + a[i-1]
	}
	//fmt.Println(seq)
	//fmt.Println(sum)
	for i = 0; i < M; i++ {
		s += a[i] * (i + 1)
		t += a[i]
	}

	sumi[0] = s
	for i = 1; i < N-M+1; i++ {
		sumi[i] = sumi[i-1] + M*a[M+i-1] - (sum[M+i-1] - sum[i-1])
	}
	for i = 0; i < N-M+1; i++ {
		if max < sumi[i] {
			max = sumi[i]
		}
	}
	fmt.Println(max)
}
