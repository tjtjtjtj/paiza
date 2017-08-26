package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Print_b_w(s string) {
	s_int, _ := strconv.Atoi(s)
	if s_int >= 128 {
		fmt.Print("1")
	} else {
		fmt.Print("0")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	h_w := strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(h_w[0])
	w, _ := strconv.Atoi(h_w[1])

	for i := 0; i < h; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		for k, v := range s {
			Print_b_w(v)
			if k != w-1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}
