package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	first_flag := true
	member := 0
	var result_array [][]string

	for scanner.Scan() {
		if first_flag == true {
			member, _ = strconv.Atoi(scanner.Text())
			result_array = make([][]string, member)
			for i := range result_array {
				result_array[i] = make([]string, member)
			}
			first_flag = false
		} else {
			s := strings.Split(scanner.Text(), " ")
			m, _ := strconv.Atoi(s[0])
			n, _ := strconv.Atoi(s[1])
			result_array[m-1][n-1] = "W"
			result_array[n-1][m-1] = "L"
		}
	}
	for i := 0; i < member; i++ {
		for j := 0; j < member; j++ {
			if i == j {
				fmt.Print("-")
			} else {
				fmt.Print(result_array[i][j])
			}
			if j != member-1 {
				fmt.Print(" ")
			}
		}
		if i != member-1 {
			fmt.Println()
		}
	}
}
