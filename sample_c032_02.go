package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	point_base_price int = 100
)

type Item struct {
	no        string
	name      string
	point_add int
}

type Customer struct {
	prices []int
	points int
}

func (c *Customer) Add_price(s string, p string) {
	s_int, _ := strconv.Atoi(s)
	p_int, _ := strconv.Atoi(p)
	c.prices[s_int] = c.prices[s_int] + p_int
}

func (c *Customer) Points_cal(items []Item) {
	for k, v := range c.prices {
		p := v / point_base_price
		c.points = c.points + (items[k].point_add * p)
	}
}

func Init_item(items []Item) {
	//func Init_item() []Item {
	fmt.Printf("%p\n", items)
	items = []Item{
		{no: "0", name: "foods", point_add: 5},
		{no: "1", name: "books", point_add: 3},
		{no: "2", name: "clothes", point_add: 2},
		{no: "3", name: "others", point_add: 1},
	}
	fmt.Printf("%p\n", items)
	//return items
}

func main() {
	var items []Item
	Init_item(items)
	//items = Init_item()
	//fmt.Print(len(items))
	//fmt.Print(items[0].name)

	var person Customer
	person.prices = make([]int, len(items))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		person.Add_price(s[0], s[1])
	}
	person.Points_cal(items)
	fmt.Println(person.points)
}
