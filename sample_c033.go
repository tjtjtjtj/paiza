package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Count struct {
	strike int
	ball   int
}

func (c *Count) judge(s string) {
	switch s {
	case "strike":
		c.strike++
		if c.strike == 3 {
			fmt.Println("out!")
		} else {
			fmt.Println("strike!")
		}
	case "ball":
		c.ball++
		if c.ball == 4 {
			fmt.Println("fourball!")
		} else {
			fmt.Println("ball!")
		}
	}
}

func main() {
	var count Count
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		count.judge(scanner.Text())
	}
}
