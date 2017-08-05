package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func chracter_add(s string) int {
	num_10 := 10 * strings.Count(s, "<")
	num := strings.Count(s, "/")
	return num_10 + num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(scanner.Text(), "+")
	var total_num int
	for _, str := range s {
		total_num = total_num + chracter_add(str)
	}
	fmt.Print(total_num)
}
