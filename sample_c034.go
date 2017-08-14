package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	items []string
}

func (p *problem) calcurate() int {
	var answer_num int
	var i int
	var pattern string
	var item_number [2]int
	for k, v := range p.items {
		switch v {
		case "x":
			answer_num = k
		case "+":
			pattern = "+"
		case "-":
			pattern = "-"
		case "=":
		default:
			item_number[i], _ = strconv.Atoi(v)
			i++
		}
	}
	if pattern == "+" {
		if answer_num == 4 {
			return item_number[0] + item_number[1]
		} else {
			return item_number[1] - item_number[0]
		}
	} else {
		if answer_num == 0 {
			return item_number[0] + item_number[1]
		} else {
			return item_number[0] - item_number[1]
		}
	}
}

func main() {
	var problem1 problem
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	problem1.items = strings.Split(scanner.Text(), " ")

	fmt.Println(problem1.calcurate())
}
