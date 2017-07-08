package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s [2]string
	var i = 0
	for scanner.Scan() {
		s[i] = scanner.Text()
		i++
		if i == len(s) {
			break
		}
	}
	var rn, sn int
	for _, rune := range s[1] {
		switch rune {
		case 'R':
			rn++
		case 'S':
			sn++
		}

	}
	fmt.Printf("%v %v", rn, sn)
}
