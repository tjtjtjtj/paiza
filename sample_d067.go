package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		switch m % 2 {
		case 0:
			fmt.Println("OFF")
		case 1:
			fmt.Println("ON")
		}
	}
}
