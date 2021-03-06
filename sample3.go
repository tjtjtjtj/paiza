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
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(s[0])
		n, _ := strconv.Atoi(s[1])
		fmt.Print(m - n)
	}
}
