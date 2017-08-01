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
	var min float64
	var max float64
	var number int

	for i := 0; scanner.Scan(); i++ {
		if first_flag == true {
			number, _ = strconv.Atoi(scanner.Text())
			first_flag = false
		} else {
			s := strings.Split(scanner.Text(), " ")
			len, _ := strconv.ParseFloat(s[1], 32)
			if s[0] == "le" && (max > len || max == 0) {
				max = len
			} else if s[0] == "ge" && min < len {
				min = len
			}
		}
		if i == number {
			break
		}
	}
	fmt.Printf("%.1f %.1f\n", min, max)
}
