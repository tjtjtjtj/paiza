package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	foods int = iota
	books
	clothes
	others
)

type Customer struct {
	prices []int
	points int
}

func (c *Customer) Add_price(s string, p string) {
	s_int, _ := strconv.Atoi(s)
	p_int, _ := strconv.Atoi(p)
	c.prices[s_int] = c.prices[s_int] + p_int
}

func (c *Customer) Points_cal() {
	for k, v := range c.prices {
		p := v / 100
		switch k {
		case foods:
			c.points = c.points + (p * 5)
		case books:
			c.points = c.points + (p * 3)
		case clothes:
			c.points = c.points + (p * 2)
		case others:
			c.points = c.points + (p * 1)
		}
	}
}

func main() {
	var person Customer
	person.prices = make([]int, 4)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		person.Add_price(s[0], s[1])
	}
	person.Points_cal()
	fmt.Println(person.points)
}
