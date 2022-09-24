package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input []string
	var N int
	var tmp_s string
	var tmp_i int
	var a []int
	var i int
	var max int = 0
	var min int = 1000000000
	var cnt int

	sc.Split(bufio.ScanWords)
	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])

	for i = 0; i < N; i++ {
		sc.Scan()
		tmp_s = sc.Text()
		tmp_i, _ = strconv.Atoi(tmp_s)
		a = append(a, tmp_i)
	}

	for {
		cnt++
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
		if max%min == 0 {
			a = append(a[:i], a[i+1:]...)
		} else {
			a[i] = max % min
		}
		if len(a) == 1 {
			break
		}
	}
	fmt.Println(cnt)
}
