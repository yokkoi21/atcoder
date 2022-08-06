package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var arr [5]string
	var tmp string
	var line []string
	var num []int
	i := 0
	for sc.Scan() {
		arr[i] = sc.Text()
		i++
		if i == 1 {
			break
		}
	}

	tmp = arr[0]
	line = strings.Split(tmp, " ")
	//fmt.Println(line)
	for i = 0; i < 5; i++ {
		tmp_value, _ := strconv.Atoi(line[i])
		num = append(num, tmp_value)
	}
	sort.Sort(sort.IntSlice(num))
	if num[0] == num[2] {
		if num[3] == num[4] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		if num[0] == num[1] && num[2] == num[4] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
