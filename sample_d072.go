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
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	z, _ := strconv.Atoi(s[2])
	syou := x / y
	if (x % y) != 0 {
		syou++
	}
	total_cost := syou * z

	fmt.Print(total_cost)
}
