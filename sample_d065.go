package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	for _, rune := range s {
		switch rune {
		case '2':
			fmt.Println("ok")
		case '4':
			fmt.Println("error")
		default:
			fmt.Println("unknown")
		}
		break
	}
}
