package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input string

	sc.Scan()
	input = sc.Text()

	switch input {
	case "Monday":
		fmt.Println("5")
	case "Tuesday":
		fmt.Println("4")
	case "Wednesday":
		fmt.Println("3")
	case "Thursday":
		fmt.Println("2")
	case "Friday":
		fmt.Println("1")
	}
}
