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
	var arr [5]string
	var line []string
	var n int
	var p []int
	var dp []int
	var test []int

	i := 0
	for sc.Scan() {
		arr[i] = sc.Text()
		i++
		if i == 2 {
			break
		}
	}

	n, _ = strconv.Atoi(arr[0])
	line = strings.Split(arr[1], " ")
	fmt.Println(line)
	p = append(p, 0)
	for i = 1; i < n; i++ {
		tmp_value, _ := strconv.Atoi(line[i-1])
		p = append(p, tmp_value)
		p[i]--
	}
	fmt.Println(p)

	dp = append(dp, 0)
	for i = 1; i < n; i++ {
		dp = append(dp, dp[p[i]]+1)
	}
	fmt.Println(dp)
	fmt.Println(dp[n-1])
}
