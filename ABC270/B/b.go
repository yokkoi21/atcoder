package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [1]string
	var split_input []string
	var X int
	var Y int
	var Z int
	var ans int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	X, _ = strconv.Atoi(split_input[0])
	Y, _ = strconv.Atoi(split_input[1])
	Z, _ = strconv.Atoi(split_input[2])

	if X >= 0 {
		if Y < 0 {
			ans += int(math.Abs(float64(X)))
		} else if Y > X {
			ans += int(math.Abs(float64(X)))
		} else {
			if Y > Z {
				if Z > 0 {
					ans += int(math.Abs(float64(X)))
				} else {
					ans += 2*int(math.Abs(float64(Z))) + int(math.Abs(float64(X)))
				}
			} else {
				ans = -1
			}
		}
	} else {
		if Y > 0 {
			ans += int(math.Abs(float64(X)))
		} else if Y < X {
			ans += int(math.Abs(float64(X)))
		} else {
			if Y < Z {
				if Z < 0 {
					ans += int(math.Abs(float64(X)))
				} else {
					ans += 2*int(math.Abs(float64(Z))) + int(math.Abs(float64(X)))
				}
			} else {
				ans = -1
			}
		}
	}
	fmt.Println(ans)
}
