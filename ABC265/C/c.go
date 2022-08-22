package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [501]string
	var split_input []string
	var H int
	var W int
	var i int
	var j int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	H, _ = strconv.Atoi(split_input[0])
	W, _ = strconv.Atoi(split_input[1])

	matrix := make([][]string, H)
	for i = 0; i < H; i++ {
		matrix[i] = make([]string, W)
	}
	for i = 0; i < H; i++ {
		sc.Scan()
		split_input = strings.Split(sc.Text(), "")
		for j = 0; j < W; j++ {
			matrix[i][j] = split_input[j]
		}
	}
	i = 0
	j = 0
	now := time.Now()
	for {
		switch matrix[i][j] {
		case "U":
			if i != 0 {
				i--
			} else {
				fmt.Println(strconv.Itoa(i+1) + " " + strconv.Itoa(j+1))
				return
			}
		case "D":
			if i != H-1 {
				i++
			} else {
				fmt.Println(strconv.Itoa(i+1) + " " + strconv.Itoa(j+1))
				return
			}
		case "L":
			if j != 0 {
				j--
			} else {
				fmt.Println(strconv.Itoa(i+1) + " " + strconv.Itoa(j+1))
				return
			}
		case "R":
			if j != W-1 {
				j++
			} else {
				fmt.Println(strconv.Itoa(i+1) + " " + strconv.Itoa(j+1))
				return
			}
		}
		if time.Since(now).Milliseconds() > 1900 {
			fmt.Println("-1")
			return
		}
	}

}
