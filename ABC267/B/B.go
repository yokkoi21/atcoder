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
	var column [7]string
	var i int
	var j int
	var split_check bool

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], "")
	//fmt.Println(split_input)

	column[0] += split_input[6]
	column[1] += split_input[3]
	column[2] += split_input[1] + split_input[7]
	column[3] += split_input[0] + split_input[4]
	column[4] += split_input[2] + split_input[8]
	column[5] += split_input[5]
	column[6] += split_input[9]

	if split_input[0] == "0" {
		for i = 0; i < len(column); i++ {
			if strings.Contains(column[i], "1") {
				for j = 1; j < len(column)-i; j++ {
					if strings.Contains(column[i+j], "1") {
						if split_check {
							fmt.Println("Yes")
							return
						}
					} else {
						split_check = true
					}
				}
			}
			split_check = false
		}
	}
	fmt.Println("No")
}
