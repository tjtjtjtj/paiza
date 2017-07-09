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
	t := strings.Split(scanner.Text(), " ")
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	n := 1
	var tint int
	for i := 0; i < len(t); i++ {
		tint, _ = strconv.Atoi(t[i])
		if tint < s {
			n++
		}
	}
	fmt.Println(n)
}
